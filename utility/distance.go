package utility

import (
	"regexp"
)

func min(numbers ...int) int {
	mn := numbers[0]
	for _, num := range numbers {
		if num < mn {
			mn = num
		}
	}
	return mn
}

func Format(code string) string {
	whitespace_regex := regexp.MustCompile(`[\s\r\n\t]+`)
	return whitespace_regex.ReplaceAllString(code, "")
}

func EditDistance(code1, code2 string) int {
	return levenshteinDistance(Format(code1), Format(code2))
}

func levenshteinDistance(code1, code2 string) int {
	A := len(code1)
	B := len(code2)

	dp0 := make([]int, B+1)
	dp1 := make([]int, B+1)

	for i := 0; i <= B; i++ {
		dp1[i] = i
	}

	for i := 1; i <= A; i++ {
		for j := 0; j <= B; j++ {
			dp0[j] = dp1[j]
		}
		dp1[0] = i

		for j := 1; j <= B; j++ {
			if code1[i-1] == code2[j-1] {
				dp1[j] = dp0[j-1]
			} else {
				dp1[j] = min(dp0[j], dp1[j-1], dp0[j-1]) + 1
			}
		}
	}

	return dp1[B]
}
