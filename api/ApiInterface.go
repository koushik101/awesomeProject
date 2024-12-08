package api

import "net/http"

type RiskAPI interface {
	ListRisksHandler(w http.ResponseWriter, r *http.Request)
	CreateRiskHandler(w http.ResponseWriter, r *http.Request)
	GetRiskHandler(w http.ResponseWriter, r *http.Request)
}
