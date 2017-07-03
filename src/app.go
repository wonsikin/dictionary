package src

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/wonsikin/dictionary/src/bing"
	"github.com/wonsikin/dictionary/src/iciba"
	"github.com/wonsikin/dictionary/src/youdao"
)

// NewApp return a new app
func NewApp(address string) *http.Server {
	r := Router()

	return &http.Server{
		Handler: r,
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

// Router create a router and register path mapping
func Router() *mux.Router {
	r := mux.NewRouter()
	// API interface
	r.HandleFunc("/wd/translate/{word}", youdao.Entry.Query).Methods("GET")
	r.HandleFunc("/wd/dictionary/{word}", youdao.Entry.QueryDictionary).Methods("GET")
	r.HandleFunc("/wd/bing/{word}", bing.Entry.Query).Methods("GET")
	r.HandleFunc("/wd/iciba/{word}", iciba.Entry.Query).Methods("GET")

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Printf("******** %s\n", t)
		return nil
	})
	return r
}
