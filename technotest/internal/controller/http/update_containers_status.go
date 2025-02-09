package http

import (
	"encoding/json"
	"log"
	"net/http"
	"technotest/internal/entity"
)

func (t *TechPointAPI) UpdateContainersStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("UpdateContainersStatus was called")

		var containersStatus []entity.ContainerStatus

		if err := json.NewDecoder(r.Body).Decode(&containersStatus); err != nil {
			log.Println("json.NewDecoder", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := t.containerService.UpdateContainersStatus(r.Context(), containersStatus); err != nil {
			log.Println("UpdateContainersStatus", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
