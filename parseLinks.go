package main

import (
	"regexp"
	"strings"
)

// parseLinks function parses the passed string and leaves only the domain names.
// For example, from the string
// "http://www.very-good-site.org/user?gjdfd=45&ldkfio=tttt"
// leaves
// "very-good-site.org"
func parseLinks(s string) string {
	s = strings.ReplaceAll(s, "[", "")
	s = strings.ReplaceAll(s, "]", "")
	s = strings.ReplaceAll(s, "http://", "")
	s = strings.ReplaceAll(s, "https://", "")
	s = strings.ReplaceAll(s, "www.", "")
	sls := strings.Split(s, " ")
	domStr := parse(sls)
	return domStr
}

// parse function gets a slice of strings,
// passes through each element of the slice,
// finds bytes in each string according to a regular expression,
// converts them to strings, removes duplicates,
// adds the result to the returned string
func parse(sls []string) string {
	regx := regexp.MustCompile(`[a-z0-9]+-?[a-z0-9]+\.[a-z]+`)
	m := make(map[string]int)
	res := ""
	for _, x := range sls {
		xbyte := regx.FindAll([]byte(x), 1)
		xstr := string(xbyte[0])
		if _, ok := m[xstr]; !ok {
			m[xstr] = 1
			res = res + xstr + ", "
		}
	}
	return strings.Trim(res, ", ")
}
