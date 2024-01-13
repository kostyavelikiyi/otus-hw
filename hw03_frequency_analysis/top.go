package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

var re = regexp.MustCompile(`\s`)

type WordCount struct {
	Word  string
	Count int
}

func Top10(str string) []string {
	return Top(Sort(Count(str)), 10)
}

func Count(str string) map[string]int {
	res := make(map[string]int)

	for _, v := range re.Split(str, -1) {
		if v != "" {
			res[v]++
		}
	}

	return res
}

func Sort(wordCount map[string]int) []string {
	keys := make([]WordCount, 0, len(wordCount))

	for k, v := range wordCount {
		keys = append(keys, WordCount{Word: k, Count: v})
	}

	sort.Slice(keys, func(i, j int) bool {
		if keys[i].Count == keys[j].Count {
			return keys[i].Word < keys[j].Word
		}
		return keys[i].Count > keys[j].Count
	})

	res := make([]string, 0, len(wordCount))

	for _, k := range keys {
		res = append(res, k.Word)
	}

	return res
}

func Top(wordCount []string, topNumper int) []string {
	resLen := 0

	if len(wordCount) < topNumper {
		resLen = len(wordCount)
	} else {
		resLen = topNumper
	}

	res := make([]string, 0, resLen)
	for i := 0; i < resLen; i++ {
		res = append(res, wordCount[i])
	}
	return res
}
