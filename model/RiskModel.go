package model

import (
	"awesomeProject/utils"
	"errors"
	"sync"
)

type RiskStore struct {
	sync.RWMutex
	data map[string]utils.Risk
}
type RiskModelImpl struct {
	riskStore RiskStore
}

var lock = &sync.Mutex{}
var riskStore = RiskStore{data: make(map[string]utils.Risk)}

func (rs *RiskModelImpl) GetRiskStore() *RiskStore {
	lock.Lock()
	defer lock.Unlock()

	if riskStore.data == nil || len(riskStore.data) == 0 {

		riskStore = RiskStore{data: make(map[string]utils.Risk)}
	}
	return &riskStore
}

func (rs *RiskModelImpl) AddRisk(risk utils.Risk) {
	riskStore := rs.GetRiskStore()
	riskStore.Lock()
	riskStore.data[risk.ID] = risk
	riskStore.Unlock()
}

func (rs *RiskModelImpl) ListRisk() ([]utils.Risk, error) {

	riskStore.RLock()
	defer riskStore.RUnlock()
	riskStore := rs.GetRiskStore()

	risks := make([]utils.Risk, 0, len(riskStore.data))
	for _, risk := range riskStore.data {
		risks = append(risks, risk)
	}
	return risks, nil
}

func (rs *RiskModelImpl) GetRisk(id string) (utils.Risk, error) {
	var risk utils.Risk
	if id == "" {
		return risk, errors.New("Risk ID is required")
	}

	riskStore.RLock()
	defer riskStore.RUnlock()
	riskStore := rs.GetRiskStore()

	risk, exists := riskStore.data[id]
	if !exists {
		return risk, errors.New("Risk not found")
	}
	return risk, nil
}
