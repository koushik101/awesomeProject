package llogic

import (
	"awesomeProject/model"
	"io"
	"strings"
	"testing"
)

func TestRiskLogicImpl_CreateRisk(t *testing.T) {
	type fields struct {
		Model model.RiskModel
	}
	type args struct {
		jsonRisk io.ReadCloser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "test_positive",
			fields: fields{Model: &model.RiskModelImpl{}},
			args: args{
				jsonRisk: io.NopCloser(strings.NewReader(`{"RiskState":"closed", "title":"critical", "description":"serverFailure"}`)),
			},
			wantErr: false,
		},
		{
			name:   "test_negative",
			fields: fields{Model: &model.RiskModelImpl{}},
			args: args{
				jsonRisk: io.NopCloser(strings.NewReader(`{"RiskState":"random", "title":"critical", "description":"serverFailure"}`)),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rl := RiskLogicImpl{
				Model: tt.fields.Model,
			}
			_, err := rl.CreateRisk(tt.args.jsonRisk)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRisk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
