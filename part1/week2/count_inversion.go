package main

import (
	"coursera/stanford-algorithms/util"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// coursera Stanford Algorithms part1-Divide and Conquer, Sorting and Searching, and Randomized Algorithms Assignment 2
//
//Question: Download the following text file:IntegerArray.txt
//
//This file contains all of the 100,000 integers between 1 and 100,000 (inclusive) in some order, with no integer repeated.
//
//Your task is to compute the number of inversions in the file given, where the ith row of the file indicates the ith entry of an array.
//
//Because of the large size of this array, you should implement the fast divide-and-conquer algorithm covered in the video lectures.
//
//The numeric answer for the given input file should be typed in the space below.
//
//So if your answer is 1198233847, then just type 1198233847 in the space provided without any space / commas / any other punctuation marks. You can make up to 5 attempts, and we'll use the best one for grading.
//
//(We do not require you to submit your code, so feel free to use any programming language you want --- just type the final numeric answer in the following space.)
//
//[TIP: before submitting, first test the correctness of your program on some small test files or your own devising. Then post your best test cases to the discussion forums to help your fellow students!]

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
	rows := util.SplitAndTrim(string(r), "\r\n")
	return util.StringToInt(rows)


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
	integers, err := readFile("/Users/xkahj/Documents/code/go/src/coursera/stanford-algorithms/part1/week2/IntegerArray.txt")
	if err != nil {
		panic(err)
	}
	_, sum := sort(integers)
	fmt.Println("sum--", sum)
}