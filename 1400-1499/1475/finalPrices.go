/*
1475 : https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/description
*/
package main

func finalPrices(prices []int) (result []int) {
o:
	for i, v := range prices {
		for _, w := range prices[i+1:] {
			if w <= v {
				result = append(result, v-w)
				continue o
			}
		}
		result = append(result, v)
	}
	return
}

func main() {

}
