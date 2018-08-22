package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IsInteger(str []string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9]+$", s)
		if false == b {
			return b
		}
	}
	return b
}


func compare(a []string, biggest []string) []string {
	// fmt.Println("b1:", biggest)
	//[v1 10 0]    [v1 10 1]
	biggestlen := len(biggest)
	alen := len(a)
	short := biggestlen
	atag := 1
	if biggestlen > alen {
		short = alen
		atag = 2
	}
	for i := 0; i < short; i++ {
		c1 := biggest[i]
		c2 := a[i]
		if i == 0 {
			c1 = strings.TrimLeft(c1, "v")
			c2 = strings.TrimLeft(c2, "v")
		}

		biggestToInt, _ := strconv.Atoi(c1)
		aToInt, _ := strconv.Atoi(c2)
		if aToInt > biggestToInt {
			biggest = a
			return biggest
		} else if aToInt < biggestToInt {
			return biggest
		}
		if i == short-1 {
			if atag == 1 {
				biggest = a
			}
			return biggest
		}
	}
	return biggest
}


//从众多tag中取到最新的tag号并返回
func GetNewestTag(a []string) string {
	biggest := []string{}

	var taglist [][]string
	var tags []string
	fmt.Println("---------")
	for _, v := range a {
		asplit := strings.Split(v, "/")
		tags = append(tags, asplit[len(asplit)-1])
	}
	for _, v := range tags {
		if v == "latest.tag" {
			return "latest"
		}
		vtrim := strings.TrimSuffix(v, ".tag")
		vsplit := strings.Split(vtrim, ".")
		if len(vsplit) > 1 {
			b := IsInteger(vsplit[1:])
			if b == true {
				taglist = append(taglist, vsplit)
			}
		} else {
			taglist = append(taglist, vsplit)
		}
	}
	if len(taglist) == 0 {
		return "nonetag"
	}
	biggest = taglist[0]
	for _, v := range taglist {
		biggest = compare(v, biggest)
	}
	fmt.Println("biggessst:", biggest)
	fmt.Println("---------")
	return strings.Join(biggest, ".")
}
