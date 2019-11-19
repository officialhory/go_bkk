package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(CallBKKAPI()))
}

var url string = "https://futar.bkk.hu/api/query/v1/ws/otp/api/where/arrivals-and-departures-for-stop.json?version=3&appVersion=apiary-1.0&includeReferences=true&stopId=BKK_F01869&onlyDepartures=false&limit=60&minutesBefore=2&minutesAfter=30"

// CallBKKAPI leave me alone ... :C
func CallBKKAPI() []byte {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func main() {
	s := &server{}
	http.Handle("/go_bkk", s)
	log.Fatal(http.ListenAndServe(":6000", nil))
}
