package api

import (
	"awesomeProject/llogic"
	"encoding/json"
	"net/http"
)

type RiskApiImpl struct {
	Logic llogic.RiskLogic
}

func (ri *RiskApiImpl) ListRisksHandler(w http.ResponseWriter, r *http.Request) {
	risks, err := ri.Logic.ListRisks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(risks)
}

func (ri *RiskApiImpl) GetRiskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/risks/"):]
	risk, err := ri.Logic.GetRisk(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(risk)
}

func (ri *RiskApiImpl) CreateRiskHandler(w http.ResponseWriter, r *http.Request) {

	newRisk, err := ri.Logic.CreateRisk(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newRisk.ID)
}
