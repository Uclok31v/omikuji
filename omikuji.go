package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var fortune = map[int]string{
	0: "大吉",
	1: "吉",
	2: "中吉",
	3: "小吉",
	4: "末吉",
	5: "凶",
	6: "大凶",
}

type Response struct {
	Omikuji string `json:"omikuji"`
}

type Time struct {
	time.Time
}

func main() {
	t := Time{time.Now()}
	http.HandleFunc("/", t.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (t *Time) handler(w http.ResponseWriter, r *http.Request) {
	var response Response

	//正月（1/1-1/3）は大吉
	if t.Month() == 1 && t.Day() >= 1 && t.Day() <= 3 {
		response = Response{Omikuji: fortune[0]}
	} else {
		rand.Seed(t.UnixNano())
		response = Response{Omikuji: fortune[rand.Intn(7)]}
	}

	res, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(res))

}
