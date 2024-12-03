/*
2109 : https://leetcode.com/problems/adding-spaces-to-a-string/description/?envType=daily-question&envId=2024-12-03
*/
package main

func addSpaces(s string, spaces []int) string {
	spaces = append(spaces, len(s)) // Add last index of s to process it
	result := []byte(s[:spaces[0]])

	for i := 1; i < len(spaces); i++ {
		result = append(result, ' ')
		result = append(result, s[spaces[i-1]:spaces[i]]...)
	}
	return string(result)
}

func main() {

}
