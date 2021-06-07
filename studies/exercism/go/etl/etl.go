package etl

import (
	"strings"
)

func Transform(scoreMap map[int][]string) map[string]int {
	transformedMap := make(map[string]int)
	for k, v := range scoreMap {
		for _, letter := range v {
			transformedMap[strings.ToLower(letter)] = k
		}
	}
	return transformedMap
}
