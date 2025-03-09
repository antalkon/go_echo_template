package utils_test

import (
	"backend/internal/models"
	"backend/internal/utils"
	"testing"
)

func TestValidator_Validate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		i       interface{}
		wantErr bool
	}{
		{
			name: "valid client",
			i: &models.User{
				PhoneNumber: "+1234567890",
				Password:    "password",
				Name:        "Frodo Beggins",
			},
			wantErr: false,
		},
		{
			name: "invalid client",
			i: &models.User{
				PhoneNumber: "1234567890",
				Password:    "password",
				Name:        "Sauron",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := utils.NewValidator()
			gotErr := v.Validate(tt.i)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Validate() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Validate() succeeded unexpectedly")
			}
		})
	}
}
