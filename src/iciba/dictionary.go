package iciba

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/CardInfoLink/log"
	"github.com/gorilla/mux"

	"github.com/wonsikin/dictionary/src/goconf"
	"github.com/wonsikin/dictionary/src/schema"
)

const (
	queryURL = "http://dict-co.iciba.com/api/dictionary.php?w=%s&key=%s&type=json"
)

type icibaDictionary struct {
	WordName string   `json:"word_name"`
	Symbols  []symbol `json:"symbols"`
}

type symbol struct {
	PhEn  string `json:"ph_en"`
	PhAm  string `json:"ph_am"`
	Parts []part `json:"parts"`
}

type part struct {
	Part  string   `json:"part"`
	Means []string `json:"means"`
}

func (e *iciba) Query(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	req, err := http.NewRequest("GET", fmt.Sprintf(queryURL, word, goconf.Config.Iciba.Key), nil)
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
	log.Debugf("body is %s", body)

	dictionary := icibaDictionary{}
	err = json.Unmarshal(body, &dictionary)
	if err != nil {
		log.Errorf("error occured when unmarshalling json: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var content = fmt.Sprintf("%s\n", dictionary.WordName)
	for _, sb := range dictionary.Symbols {
		content += fmt.Sprintf("英[%s] 美[%s]\n", sb.PhEn, sb.PhAm)
		for i, pt := range sb.Parts {
			content += fmt.Sprintf("%d. %s %s\n", i+1, pt.Part, strings.Join(pt.Means, "，"))
		}
	}

	rs := schema.Vocabulary{
		Content: content,
	}

	data, err := json.Marshal(rs)
	if err != nil {
		log.Errorf("error occured when unmarshalling json: %s", err)
	}
	log.Debugf("response data is %s", string(data))

	w.Write(data)
}
