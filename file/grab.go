package file

import (
	"fmt"
	"strings"
)

func find(split []string, tag string) (int, error) {
	for i, v := range split {
		if strings.Contains(v, tag) {
			return i, nil
			break
		}
	}
	return -1, fmt.Errorf("Tag: %s not found", tag)
}

func combine(split []string, startLine, endLine int) string {
	s := ""
	sep := ""
	for i := startLine; i < endLine; i++ {
		s = s + sep + split[i]
		sep = "\n"
	}
	return s
}

func GrabTab(s string, tagBegin, tagEnd string) (string, string, error) {
	split := strings.Split(s, "\n")
	startLine := 0
	endLine := 0
	var err error

	startLine, err = find(split, tagBegin)
	if err != nil {
		return s, "", err
	}

	endLine, err = find(split, tagEnd)
	if err != nil {
		return s, "", err
	}

	if startLine >= endLine {
		return s, "", fmt.Errorf("Start > End")
	}

	sub := combine(split, 0, startLine)
	tag := combine(split, startLine, endLine+1)

	return sub, tag, nil

}
