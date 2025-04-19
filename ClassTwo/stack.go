// leetcode Q20:Valid Parentheses
package classtwo

func isValid(s string) bool {
	var stack []string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{" {
			stack = append(stack, string(s[i]))
		}
		if (string(s[i]) == ")" || string(s[i]) == "]" || string(s[i]) == "}") && len(stack) == 0 {
			return false
		}
		if string(s[i]) == ")" {
			if stack[len(stack)-1] != "(" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		if string(s[i]) == "]" {
			if stack[len(stack)-1] != "[" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		if string(s[i]) == "}" {
			if stack[len(stack)-1] != "{" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

//Runtime 0ms Beats 100%
//Memory 4.72MB Beats 9.04%
