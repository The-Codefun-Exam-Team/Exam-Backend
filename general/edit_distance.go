package general

import (
	"regexp"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func EditDistance(rawcode1 string, rawcode2 string) int {
    var i, j int

	code1 := Format(rawcode1)
	code2 := Format(rawcode2)

	n := len(code1)
	m := len(code2)

	dp1 := make([]int, m+1)
	dp2 := make([]int, m+1)

	for i = 0; i <= m; i++ {
		dp1[i] = i
	}

	for i = 1; i <= n; i++ {
		dp2[0] = i
		for j = 1; j <= m; j++ {
			if code1[i-1] == code2[j-1] {
				dp2[j] = dp1[j-1]
			} else {
                dp2[j] = 1 + min(dp1[j], min(dp2[j-1], dp1[j-1]))
            }
        }

		for j = 0; j <= m; j++ {
			dp1[j] = dp2[j]
		}
	}

	return dp1[m]
}

func Format(code string) string {
	whitespace_pattern := regexp.MustCompile(`/[\s\r\n\t]+/g`)
	code = whitespace_pattern.ReplaceAllString(code, "")

	return code
}