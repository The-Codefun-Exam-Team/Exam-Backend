package utility

import (
	"math"
)

func CalculateScore(diff, min_diff int, score, old_score float64) float64 {
	var score_percentage, edit_percentage float64
	// TODO: Check for edge cases

	score_percentage = math.Max(score - old_score, 0) / (100 - old_score)
	edit_percentage = math.Min(float64(min_diff) / float64(diff), 1)

	final_score := score_percentage * float64(edit_percentage) * 100
	final_score = math.Max(final_score, 0)
	final_score = math.Min(final_score, 100)

	return final_score
}
