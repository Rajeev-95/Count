
package main

import (
	"fmt"
	"strings"
)

func main() {
	var dict_size int
	fmt.Println("Enter the size of the dictionary")
	fmt.Scan(&dict_size)
	dict := make([]string, dict_size)
	for i := 0; i < dict_size; i++ {
		fmt.Printf("Enter a word")
		fmt.Scan(&dict[i])

	}
	var str_size int
	fmt.Println("How many lines are there?")
	fmt.Scan(&str_size)
	str := make([]string, str_size)
	for i := 0; i < str_size; i++ {
		fmt.Printf("Enter a string")
		fmt.Scan(&str[i])
	}
	for c := 0; c < str_size; c++ {
		count := 0
		for i := 0; i < dict_size; i++ {

			s := dict[i]
			mp := make(map[string]int)
			for i := 1; i < len(s)-1; i++ {
				mp[s[i]]++
				st := make(map[string]int)
				for i := 0; i < len(str[c]); i++ {
					if str[c][i] == s[0] && str[c][i+len(s)-1] == s[len(s)-1] && (i+len(s)-1) < len(str[c]) {
						mp1 := make(map[string]int)
						for j := i + 1; j < i+len(s)-1; j++ {
							mp1[str[c][j]]++
							if mp1 == mp {
								var r string = str[c].strings.contains(i, len(s)) // string r is required string
								st.insert(r)

							}

						}
					}
					count += size(st)

				}

			}
		}
		fmt.Println("case #", c+1)
		fmt.Print("%d:", count)
	}
}

