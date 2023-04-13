package slice

import (
	"fmt"
	"sort"
	"strings"
)

// leetcode 189. Rotate Array
// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
func rotate1(nums []int, k int) {
	// O(N)/O(N)
	if k == 0 || len(nums) == 1 {
		return
	}

	validK := k % len(nums)
	splitIndex := len(nums) - validK
	rightSlice := make([]int, splitIndex)
	copy(rightSlice, nums[:splitIndex])
	leftSlice := nums[splitIndex:]

	i := 0
	for i < len(leftSlice) {
		nums[i] = leftSlice[i]
		i++
	}

	j := 0
	for i < len(nums) {
		nums[i] = rightSlice[j]
		i++
		j++
	}
}

func rotate2(nums []int, k int) {
	// O(N)/O(1)
	validK := k % len(nums)
	// reverse and rotate
	if len(nums) <= 1 || validK == 0 {
		return
	}

	reverse(nums)
	reverse(nums[:validK])
	reverse(nums[validK:])
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i <= j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func wordCount(str string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(str) {
		m[w]++
	}
	return m
}

func sortByCount(m map[string]int) []string {
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	count := func(i, j int) bool { return m[s[i]] > m[s[j]] }
	sort.Slice(s, count)
	return s
}

func PrintSlice() {
	// init
	s1 := make([]int, 5)
	s2 := make([]int, 3, 5)
	// modify value of element
	s1[0] = 3
	fmt.Println(s1)
	// copy
	s3 := make([]int, 3)
	copy(s3, s1[0:3])
	fmt.Println(s1, s2, s3)

	// exercise 1
	testSlice1 := []int{1, 2, 3, 4, 5, 6, 7}
	testSlice2 := []int{-1, -100, 3, 99}
	k1, k2 := 3, 2
	rotate2(testSlice1, k1)
	rotate1(testSlice2, k2)
	fmt.Printf("rotate %d steps to the right: %v\n", k2, testSlice1)
	fmt.Printf("rotate %d steps to the right: %v\n", k1, testSlice2)

	// exercise 2
	testStr := "hello fuck ok ok hello go go big big big at at at at go go go"
	fmt.Println("test str is", testStr)
	fmt.Println("sort slice is", sortByCount(wordCount(testStr)))
}
