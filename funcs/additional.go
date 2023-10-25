package funcs

import (
	"fmt"
	"strings"
)

func StoreInMap(inp string) map[rune][]string {
	inp = CleanFile(inp)
	mp := make(map[rune][]string)
	i := 1
	for r := 32; r < 127 && i < len(inp); r++ {
		oneline := ""
		for l := 0; l < 8; l++ {

			for inp[i] != '\n' {
				oneline = oneline + string(inp[i])
				i++

			}
			mp[rune(r)] = append(mp[rune(r)], oneline)
			i++
			oneline = ""
		}
		for i < len(inp) && inp[i] != '\n' {
			i++
		}
		i++
	}
	return mp
}

func GetWord(word string, mp map[rune][]string) {
	for j := 0; j < 8; j++ {
		for ind := 0; ind < len(word); ind++ {
			if ind != len(word)-1 {
				fmt.Print(mp[rune(word[ind])][j])
			} else {
				fmt.Println(mp[rune(word[ind])][j])
			}
		}
	}
}

func CleanFile(inp string) string {
	res := ""
	for _, let := range inp {
		if let != 13 {
			res += string(let)
		}
	}
	return res
}

func IsValid(inp []string) bool {
	if len(inp) == 0 || len(inp) > 2 {
		fmt.Println("ERROR: the number of arguments is incorrect")
		return false
	} else if len(inp) == 2 {
		if inp[1] == "shadow" || inp[1] == "thinkertoy" || inp[1] == "standard" {
			return true
		} else {
			fmt.Println("ERROR: the filename is incorrect!")
			return false
		}
	}

	for _, let := range inp[0] {
		if let < 32 || let > 126 {
			if let != 10 {
				fmt.Println("ERROR: the string is invalid (NOT ASCII RANGE)")
				return false
			}
		}
	}
	return true
}

func SepNewLine(s []string) []string {
	var res []string
	for i := 0; i < len(s); i++ {
		res = append(res, strings.Split(s[i], string('\n'))...)
	}
	return res
}
