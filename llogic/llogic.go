package llogic

import (
	"awesomeProject/model"
	"awesomeProject/utils"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"io"
)

type RiskLogicImpl struct {
	Model model.RiskModel
}

// Helper to validate risk state
func isValidState(State utils.RiskState) bool {
	return utils.ValidStates[State]
}

func (rl RiskLogicImpl) CreateRisk(jsonRisk io.ReadCloser) (utils.Risk, error) {
	var newRisk utils.Risk
	if err := json.NewDecoder(jsonRisk).Decode(&newRisk); err != nil {
		return newRisk, errors.New("Invalid JSON payload")
	}
	if !isValidState(newRisk.State) {
		return newRisk, errors.New("Invalid state value")
	}

	if newRisk.Title == "" || newRisk.Description == "" {
		return newRisk, errors.New("Title and description are required")
	}

	newRisk.ID = uuid.New().String()

	rl.Model.AddRisk(newRisk)
	return newRisk, nil
}

func (rl RiskLogicImpl) ListRisks() ([]utils.Risk, error) {
	risks, err := rl.Model.ListRisk()
	return risks, err
}

func (rl RiskLogicImpl) GetRisk(id string) (utils.Risk, error) {

	risk, err := rl.Model.GetRisk(id)
	return risk, err
}
