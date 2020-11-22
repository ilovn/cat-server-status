package main

import (
	"cat-server-status/api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/record", api.Record).Methods("POST")
	router.HandleFunc("/record/{token}", api.Record).Methods("POST")
	router.HandleFunc("/node", api.Node).Methods("POST")
	router.HandleFunc("/node/{token}", api.Node).Methods("POST")

	fmt.Println("Start server ...")
	err := http.ListenAndServe(":5321", defaultHandler(router))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func defaultHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("defaultHandler process ...", r.URL.Path)
		//if r.URL.Path == "/" {
		//	http.Redirect(w, r, "/static/", http.StatusFound)
		//	return
		//}
		if r.URL.Path == "/static/" {
			next.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
			//if len(r.Header.Get("Token")) <= 0 {
			//	w.WriteHeader(http.StatusForbidden)
			//} else {
			//	next.ServeHTTP(w, r)
			//}
		}

	})
}