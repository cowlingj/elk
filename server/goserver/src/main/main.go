package main

import (
	"encoding/json"
	"net/http"
  "os"
	"strconv"
	"math/rand"
	"log"
)

func main() {

	successRate, successRateEnvSet := strconv.ParseFloat(os.Getenv("SUCCESS_RATE"), 64)

	if (successRateEnvSet != nil) {
		panic("SUCCESS_RATE has not been set")
	}

	logfile, logfileErr := os.Create("/var/log/goserver/goserver.log")

	if (logfileErr != nil) {
		panic(logfileErr)
	}

	logger := log.New(logfile, "goserver", log.LstdFlags)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		logger.Println("got request")

		if (rand.Float64() < successRate) {
			w.WriteHeader(http.StatusOK)

			logger.Println("success")

			body, _ := json.Marshal(struct {
				Value string `json:"value"`
			}{
				Value: "success",
			})
 
			w.Write(body)

		} else {
			w.WriteHeader(http.StatusInternalServerError)

			logger.Println("failure")

			body, _ := json.Marshal(struct {
				Value string `json:"value"`
			}{
				Value: "error",
			})

			w.Write(body)
		}
	})

	http.ListenAndServe(":80", nil)
}