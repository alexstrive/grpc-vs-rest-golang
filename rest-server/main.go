package main

import (
	"encoding/json"
	"log"
	"net/http"
	pb "stats"
)

func main() {
	http.HandleFunc("/covid.json", func(rw http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(pb.CovidCases)
		if err != nil {
			log.Printf("Unable to marshal data: %v", err)
		}

		rw.Write(data)
	})

	log.Printf("Handlers have been r::egistered")

	http.ListenAndServe(":8080", nil)
	log.Printf("Server has started")
}
