package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println("Hello, playground")
	fmt.Println(ParseNumbers("1-8, -4, -7", 1, 10))
	fmt.Println(ParseNumbers("*, -4, -7", 1, 10))
}

func ParseNumbers(s string, start, end int) []int {
	nums := make(map[int]int)

	parts := strings.Split(s, ",")
	for _, part := range parts {
		part = strings.Trim(part, " ")
		if part == "*" {
			for n := start; n <= end; n++ {
				nums[n] = n
			}
			continue
		}
		if n, err := strconv.Atoi(part); err == nil {
			if n >= 0 {
				nums[n] = n
			} else {
				delete(nums, -n)
			}
		} else {
			ranges := strings.Split(part, "-")
			if len(ranges) == 2 {
				min, _ := strconv.Atoi(strings.Trim(ranges[0], " "))
				max, _ := strconv.Atoi(strings.Trim(ranges[1], " "))
				if err1 == nil && err2 == nil {
					for n := min; n <= max; n++ {
						nums[n] = n
					}
				}
			}
		}
	}

	var result []int
	for _, n := range nums {
		result = append(result, n)
	}
	sort.Ints(result)
	return result
}
