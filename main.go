package main

import (
	"fmt"
	"regexp"
)

func getRegexp(search string, searchType string) string {
	var regexString string
	var replaceString string

	switch searchType {
	case "snort":
		regexString = `\t`
		replaceString = `,`
	case "zeek":
		regexString = `\t`
		replaceString = `\s*`
	}
	re := regexp.MustCompile(regexString)
	returnRegex := "^" + re.ReplaceAllString(search, replaceString)
	fmt.Printf("%v: returning regex: %v\n", searchType, returnRegex)
	return returnRegex
}

func main() {
	// haystack := "2022-06-26 00:01:01.123456\t         FSDFwersdflkjaSDDFSFD asdfdsaf asdfasdf 1.2.3.4 23.4.5.5"
	haystackZeek := "2022-06-26 00:01:01.123456\tFSDFwersdflkjaSDDFSFD asdfdsaf asdfasdf 1.2.3.4 2.3.4.5"
	haystackSnort := "2022-06-26 00:01:01.123456,FSDFwersdflkjaSDDFSFD,asdfdsaf,asdfasdf,1.2.3.4,2.3.4.5"
	search := "2022-06-26 00:01:01.123456\tFSDFwersdflkjaSDDFSF"
	// needle := "2022-06-26 00:01:01.123456\tFSDFwersdflkjaSDDFSF"
	// needle := `2022-06-26 00:01:01.123456\s*FSDFwersdflkjaSDDFSF`
	needle := getRegexp(search, "zeek")
	matched, err := regexp.MatchString(needle, haystackZeek)
	if err != nil {
		fmt.Printf("Error: MatchString zeek:needle: %v, haystack: %v, error: %v\n", needle, haystackZeek, err.Error())
	} else {
		if matched {
			fmt.Printf("MATCH: zeek: search: %v, needle: %v, haystack: %v\n", search, needle, haystackZeek)
		}
	}

	needle = getRegexp(search, "snort")
	matched, err = regexp.MatchString(needle, haystackSnort)
	if err != nil {
		fmt.Printf("Error: MatchString snort: needle: %v, haystack: %v, error: %v\n", needle, haystackSnort, err.Error())
	} else {
		if matched {
			fmt.Printf("MATCH: snort: search: %v, needle: %v, haystack: %v\n", search, needle, haystackZeek)
		}
	}
}
