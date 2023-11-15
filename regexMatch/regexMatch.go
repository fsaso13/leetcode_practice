package main

import "fmt"

/*
 * Support regex match for the full string
 * s - input string
 * p - regex pattern to match:
 *     . - Any char
 *     * - Previous char 0-N times
 * Note - There won't be any cases where * is not valid (i.e., *asd or **)
 */
func isMatch(s string, p string) bool {
	byte_pattern := []rune(p)
	byte_str := []rune(s)

	str_index := 0
	pattern_index := 0
	prev_char := string(byte_str[0])

	if p == ".*" {
		return true
	}

	/*
	 *	Intent - use a sliding window of size 2 to traverse the pattern while matching the string
	 */
	for {
		if str_index >= len(s) {
			// fmt.Printf("%d - %d\n", str_index, pattern_index)
			if pattern_index < len(p) {
				pattern_index += 2
			}
			// fmt.Printf("%d - %d\n", str_index, pattern_index)
			if pattern_index == len(p) {
				return true
			} else {
				return false
			}
		}

		// fmt.Printf("%d - %d\n", pattern_index, str_index)
		if pattern_index >= len(p) && str_index < len(s) {
			return false
		}

		strChar := string(byte_str[str_index])
		pattChar := string(byte_pattern[pattern_index])

		if pattChar == "." {
			if len(p) > 1 && pattern_index != len(p)-1 && string(byte_pattern[pattern_index+1]) == "*" {
				if strChar == prev_char {
					if str_index == len(s)-1 && pattern_index+1 != len(p)-1 {
						// if string is like aaaaaa and pattern is .*a, make sure to use the whole regex to match.
						pattern_index += 2
					} else if len(p)-(pattern_index+2) >= len(s)-str_index {
						if string(byte_str[str_index+1]) == prev_char {
							str_index++
							continue
						}
						pattern_index += 2
						if string(byte_str[str_index+1]) == string(byte_pattern[pattern_index]) {
							str_index++
						}
					} else {
						str_index++
					}
					continue
				} else {
					pattern_index += 2
					continue
				}
			} else {
				str_index++
				pattern_index++
				continue
			}
		} else if pattChar == "*" {
			fmt.Println("I don't think this case should ever get hit")
			return false
		} else {
			if len(p) > 1 && pattern_index != len(p)-1 && string(byte_pattern[pattern_index+1]) == "*" {
				if str_index == len(s)-1 && pattern_index+1 != len(p)-1 {
					// if string is like aaaaaa and pattern is .*a, make sure to use the whole regex to match.
					pattern_index += 2
				} else {
					if strChar != pattChar {
						pattern_index += 2
					} else {
						str_index++
					}
				}
				continue
			} else {
				if pattChar == strChar {
					str_index++
					pattern_index++
					continue
				} else {
					return false
				}
			}
		}
	}
}

func printCase(s string, p string, r string, n *int) {
	fmt.Println(*n)
	*n++
	fmt.Println("===============")
	fmt.Println(isMatch(s, p))
	fmt.Println("Expected: " + r)
	fmt.Println("===============")
	fmt.Println()
}

func main() {

	caseN := 1

	printCase("aa", "a", "False", &caseN)

	printCase("aa", "a*", "True", &caseN)

	printCase("ab", ".*", "True", &caseN)

	printCase("aab", "c*a*b", "True", &caseN)

	printCase("mississippi", "mis*is*ip*.", "True", &caseN)

	printCase("mississippi", "mis*is*p*.", "False", &caseN)

	printCase("ab", ".*c", "False", &caseN)

	printCase("aaa", "aaaa", "False", &caseN)

	printCase("aaba", "ab*a*c*a", "False", &caseN)

	printCase("a", "ab*", "True", &caseN)

	printCase("aaa", "a*a", "True", &caseN)

	printCase("ab", ".*..", "True", &caseN)

	printCase("aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s", "True", &caseN)

	printCase("bbbba", ".*a*a", "True", &caseN)

}
