package main

import (
	"encoding/json"
	"net/http"
  "os"
	"strconv"
	"math/rand"
)

func main() {

	successRate, successRateEnvSet := strconv.ParseFloat(os.Getenv("SUCCESS_RATE"), 64)

	if (successRateEnvSet != nil) {
		panic("SUCCESS_RATE has not been set")
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		if (rand.Float64() < successRate) {
			w.WriteHeader(http.StatusOK)

			body, _ := json.Marshal(struct {
				Value string `json:"value"`
			}{
				Value: "success",
			})
 
			w.Write(body)

		} else {
			w.WriteHeader(http.StatusInternalServerError)

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