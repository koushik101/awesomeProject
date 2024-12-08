package main

import (
	"awesomeProject/api"
	"awesomeProject/llogic"
	"awesomeProject/model"
	"log"
	"net/http"
)

func main() {
	var service api.RiskAPI

	service = &api.RiskApiImpl{
		Logic: llogic.RiskLogicImpl{
			Model: &model.RiskModelImpl{},
		},
	}

	http.HandleFunc("/v1/risks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			service.ListRisksHandler(w, r)
		case http.MethodPost:
			service.CreateRiskHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/v1/risks/", service.GetRiskHandler)

	log.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
