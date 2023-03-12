package debugproblem_test

import (
	"testing"
)

func TestGetSingleProblem(t *testing.T) {
	testcases := []struct{
		code string
		token string
		expected string
	}{
		{
			"D001", "good-token",
			"to-be-added",
		},
		{
			"D002", "good-token",
			"to-be-added",
		},
		{
			"non-existent-problem", "good-token",
			"to-be-added",
		},
		{
			"D001", "bad-token",
			"to-be-added",
		},
	}

	for _, test := range testcases {
		
	}
}