package main

import (
	"fmt"
	"strings"
)

func RevStr(str string, i int, res string) string {
	if i == len(str) {
		return res
	}

	return RevStr(str, i+1, string(str[i])+res)
}

func IsPolyndrom(str string) bool {
	var isPolyndrome func(string, int, int) bool

	isPolyndrome = func(str string, i, j int) bool {

		if str[i] != str[j] {
			return false
		}
		if i >= j {
			return true
		}
		return isPolyndrome(str, i+1, j-1)
	}

	return isPolyndrome(strings.ToLower(str), 0, len(str)-1)

}

func Power(num, power int) int {
	if power == 0 {
		return 1
	}

	return num * Power(num, power-1)
}

func CountDigits(n int) int {
	if n == 0 {
		return 0
	}
	return CountDigits(n/10) + 1
}

func GCDofNum(num1, num2 int) int {
	if num2 == 0 {
		return num1
	}

	return GCDofNum(num2, num1%num2)
}

func SumRow(nums []int, i int) int {
	if i == len(nums) {
		return 0
	}
	return nums[i] + SumRow(nums, i+1)
}

func SumMatrix(nums [][]int, row int) int {
	if row == len(nums) {
		return 0
	}
	return SumRow(nums[row], 0) + SumMatrix(nums, row+1)
}

func TowerOfHanoi(n int, fromrode string, torode string, auxrode string) {
	if n == 0 {
		return
	}
	TowerOfHanoi(n-1, fromrode, auxrode, torode)
	fmt.Println()
	fmt.Printf(" disk %d moved from %s to %s", n, fromrode, torode)

	TowerOfHanoi(n-1, auxrode, torode, fromrode)
}

func Subset(str string) {
	var subsets func(int)
	res := []byte{}
	n := len(str)
	num := 0

	subsets = func(i int) {
		if i == n {
			num = +1
			fmt.Printf("%d char is:%s \n", num, string(res))
			return
		}
		res = append(res, str[i])
		subsets(i + 1)
		res = res[:len(res)-1]
		subsets(i + 1)
	}

	subsets(0)

}

func SumOfArray(ar []int, i int) int {
	if i == len(ar) {
		return 0
	}
	return SumOfArray(ar, i+1) + ar[i]
}

func Permutaion(s string) {
	res := []string{}

	var permute func([]rune, int)

	permute = func(str []rune, row int) {
		if row == len(str) {
			res = append(res, string(str))
			return
		}

		for i := row; i < len(str); i++ {
			str[i], str[row] = str[row], str[i]
			permute(str, row+1)
			str[i], str[row] = str[row], str[i]
		}
	}
	permute([]rune(s), 0)
	fmt.Println(res)
}

func main() {
	fmt.Println(RevStr("Hello world", 0, ""))

	fmt.Println(IsPolyndrom("mlayalM"))

	fmt.Println(Power(2, 3))
	fmt.Println(CountDigits(1231241))

	fmt.Println(GCDofNum(48, 18))

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Printf("matrix sum:%d", SumMatrix(matrix, 0))

	TowerOfHanoi(3, "A", "B", "C")

	nums := []int{2, 2, 1, 2}
	fmt.Println()

	fmt.Printf("sum of array:%d", SumOfArray(nums, 0))
	fmt.Println()
	Permutaion("ABC")
	fmt.Println()
	Subset("123")

}
