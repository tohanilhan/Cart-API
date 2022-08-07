package main

import (
	"testing"
)

func Test_initUserID(t *testing.T) {
	tests := []struct {
		name string
	}{
		// Add test cases.

		{
			name: "initUserID",
		},

		{
			name: "initUserID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initUserID()
		})
	}
}

func Test_initGivenAmount(t *testing.T) {
	tests := []struct {
		name string
	}{
		// Add test cases.

		{
			name: "initGivenAmount",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initGivenAmount()
		})
	}
}
