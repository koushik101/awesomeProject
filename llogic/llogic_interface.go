package llogic

import "io"
import "awesomeProject/utils"

type RiskLogic interface {
	CreateRisk(jsonRisk io.ReadCloser) (utils.Risk, error)
	ListRisks() ([]utils.Risk, error)
	GetRisk(id string) (utils.Risk, error)
}
