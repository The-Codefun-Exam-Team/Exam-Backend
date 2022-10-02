package general

import (
	"math"
)

var epsilon = 1e-6

func CalculateScore(diff int, score float64, org_score float64, min_diff int) float64 {
	var edit_percentage, score_percentage float64
	if math.Abs(score-org_score) < epsilon {
		if math.Abs(score-100) < epsilon {
			score_percentage = 1.0
		} else {
			score_percentage = 0.0
		}
	} else {
		score_percentage = (score - org_score) / (100 - org_score)
	}
	if diff > min_diff {
		edit_percentage = float64(min_diff) / float64(diff)
	} else {
		edit_percentage = 1.0
	}

	finalscore := score_percentage * float64(edit_percentage) * 100

	if finalscore < 0 {
		finalscore = 0
	}
	if finalscore > 100 {
		finalscore = 100
	}

	return finalscore
}
