/*
3190 : https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/
1. Iterate through each element in the array.
2. Check if the element is not divisible by 3.
3. If not divisible, increment the operation count.
4. Return the total count of operations needed.
*/
package main

func minimumOperations(nums []int) (result int) {
	for _, x := range nums {
		if x%3 != 0 {
			result++
		}
	}
	return
}

func main() {

}
