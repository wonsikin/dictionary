package bing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/CardInfoLink/log"
	"github.com/gorilla/mux"

	"github.com/wonsikin/dictionary/src/schema"
)

const (
	queryURL = "https://xtk.azurewebsites.net/BingDictService.aspx?Word=%s"
)

type bingDictionary struct {
	Word          string        `json:"word"`
	Pronunciation pronunciation `json:"pronunciation"`
	Definitions   []definition  `json:"defs"`
	Samples       []sample      `json:"sams"`
}

type pronunciation struct {
	AmE    string `json:"AmE"`
	AmEmp3 string `json:"AmEmp3"`
	BrE    string `json:"BrE"`
	BrEmp3 string `json:"BrEmp3"`
}

type definition struct {
	Pos        string `json:"pos"`
	Definition string `json:"def"`
}

type sample struct {
	English string `json:"eng"`
	Chinese string `json:"chn"`
	MP3URL  string `json:"mp3Url"`
	MP4URL  string `json:"mp4Url"`
}

func (e *bing) Query(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	req, err := http.NewRequest("GET", fmt.Sprintf(queryURL, word), nil)
	if err != nil {
		log.Errorf("error occured when creating request: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("error occured when sending request: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("error occured when reading all body: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	dictionary := bingDictionary{}
	err = json.Unmarshal(body, &dictionary)
	if err != nil {
		log.Errorf("error occured when unmarshalling xml: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var content = fmt.Sprintf("%s\n", dictionary.Word)
	content += fmt.Sprintf("英[%s] 美[%s]\n", dictionary.Pronunciation.BrE, dictionary.Pronunciation.AmE)
	for i, df := range dictionary.Definitions {
		content += fmt.Sprintf("%d. %s %s\n", i+1, df.Pos, df.Definition)
	}
	rs := schema.Vocabulary{
		Content: content,
	}

	data, err := json.Marshal(rs)
	if err != nil {
		log.Errorf("error occured when unmarshalling json: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Debugf("data is %s", data)

	w.Write(data)
}
