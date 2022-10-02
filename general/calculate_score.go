package general

func CalculateScore(diff int, score float64, org_score float64, min_diff int) float64 {
	score_percentage := (score - org_score) / (100 - org_score)
	edit_percentage := min_diff / diff

	finalscore := score_percentage * float64(edit_percentage) * 100

	if finalscore < 0 {
		finalscore = 0
	}
	if finalscore > 100 {
		finalscore = 100
	}

	return finalscore
}
