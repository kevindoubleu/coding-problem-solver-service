package twosum

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Nums []int `json:"nums"`
	Target int `json:"target"`
}

func (r Request) String() string {
	marshalled, _ := json.Marshal(r)
	return string(marshalled)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	req := r.Context().Value("data").(Request)
	log.Print(req)
	w.Write([]byte(req.String()))
}