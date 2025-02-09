package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func (t *TechPointAPI) GetContainersStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GetContainersStatus was called")

		containersStatus, err := t.containerService.GetContainersStatus(r.Context())
		if err != nil {
			log.Printf("containerService.GetContainersStatus: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(containersStatus)
		if err != nil {
			log.Printf("json.Marshal: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	}
}
