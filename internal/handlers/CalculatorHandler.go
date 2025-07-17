package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func CalcAddHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println(r.Method, r.URL.Path)
	fmt.Println(r.GetBody())
	w.Write([]byte("Add endpoint hit"))
}
