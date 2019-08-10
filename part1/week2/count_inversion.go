package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func merge(left, right []int) ([]int, int) {
	i := 0
	j := 0
	var sort []int
	var count int
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sort = append(sort, left[i])
			i++
			continue
		} else {
			sort = append(sort, right[j])
			j++
			count = count + (len(left) - i)
		}
	}
	sort = append(sort, right[j:]...)
	sort = append(sort, left[i:]...)
	return sort, count
}

func sort(a []int) ([]int, int) {
	var sum int
	if len(a) < 2 {
		return a, sum
	}
	left, num1 := sort(a[:len(a)/2])
	right, num2 := sort(a[len(a)/2:])
	result, count := merge(left, right)
	sum = sum + num1 + num2 + count
	return result, sum

}

func readFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	rows := splitAndTrim(string(r), "\r\n")
	nums := stringToInt(rows)
	return nums, nil

}

func splitAndTrim(oldString, split string) []string {
	arr1 := strings.Split(oldString, split)
	arr2 := make([]string, len(arr1))
	for i := 0; i < len(arr1); i++ {
		b := []byte(arr1[i])
		for j := 0; j < len(b); j++ {
		}
		arr2[i] = strings.TrimSpace(arr1[i])
	}
	var arr3 []string
	for _, item := range arr2 {
		if item != "" {
			arr3 = append(arr3, item)
		}
	}
	return arr3
}

func stringToInt(a []string) []int {
	var nums []int
	for _, item := range a {
		num, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println("err--", err)
		}
		nums = append(nums, num)
	}
	return nums
}

func main() {
	integers, err := readFile("/Users/yangziyun/Downloads/assignment2.txt")
	if err != nil {
		panic(err)
	}
	_, sum := sort(integers)
	fmt.Println("sum--", sum)
}