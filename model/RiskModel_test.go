package model

import (
	"awesomeProject/utils"
	"testing"
)

func TestRiskModelImpl_ListRisk(t *testing.T) {
	type fields struct {
		riskStore RiskStore
	}
	tests := []struct {
		name string
		//	fields  fields
		want    int
		wantErr bool
	}{
		{"test1", 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &RiskModelImpl{
				//	riskStore: tt.fields.riskStore,
			}
			rs.AddRisk(utils.Risk{
				State: "closed", Title: "critical", Description: "serverFailure",
			})
			got, err := rs.ListRisk()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRisk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("ListRisk() got = %v, want %v", len(got), tt.want)
			}
		})
	}
}
