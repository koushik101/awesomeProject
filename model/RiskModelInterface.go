package model

import "awesomeProject/utils"

type RiskModel interface {
	GetRiskStore() *RiskStore
	AddRisk(risk utils.Risk)
	ListRisk() ([]utils.Risk, error)
	GetRisk(id string) (utils.Risk, error)
}
