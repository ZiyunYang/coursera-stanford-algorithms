package main

import (
	"coursera/stanford-algorithms/util"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// coursera Stanford Algorithms part1-Divide and Conquer, Sorting and Searching, and Randomized Algorithms Assignment 3
//
//Question: Your task is to compute the total number of comparisons used to sort the given input file by QuickSort. As you know, the number of comparisons depends on which elements are chosen as pivots, so we'll ask you to explore three different pivoting rules.
//
//You should not count comparisons one-by-one. Rather, when there is a recursive call on a subarray of length m, you should simply add m−1 to your running total of comparisons. (This is because the pivot element is compared to each of the other m−1 elements in the subarray in this recursive call.)
//
//WARNING: The Partition subroutine can be implemented in several different ways, and different implementations can give you differing numbers of comparisons. For this problem, you should implement the Partition subroutine exactly as it is described in the video lectures (otherwise you might get the wrong answer).
//
//DIRECTIONS FOR THIS PROBLEM: (1) For the first part of the programming assignment, you should always use the first element of the array as the pivot element. (2) Compute the number of comparisons (as in Problem 1), always using the final element of the given array as the pivot element. Again, be sure to implement the Partition subroutine exactly as it is described in the video lectures. Recall from the lectures that, just before the main Partition subroutine, you should exchange the pivot element (i.e., the last element) with the first element. (3) Compute the number of comparisons (as in Problem 1), using the "median-of-three" pivot rule. [The primary motivation behind this rule is to do a little bit of extra work to get much better performance on input arrays that are nearly sorted or reverse sorted.] In more detail, you should choose the pivot as follows. Consider the first, middle, and final elements of the given array. (If the array has odd length it should be clear what the "middle" element is; for an array with even length 2k 2k, use the kth element as the "middle" element. So for the array 4 5 6 7, the "middle" element is the second one ---- 5 and not 6!) Identify which of these three elements is the median (i.e., the one whose value is in between the other two), and use this as your pivot. As discussed in the first and second parts of this programming assignment, be sure to implement Partition exactly as described in the video lectures (including exchanging the pivot element with the first element just before the main Partition subroutine).
//
//EXAMPLE: For the input array 8 2 4 5 7 1 you would consider the first (8), middle (4), and last (1) elements; since 4 is the median of the set {1,4,8}, you would use 4 as your pivot element.
//
//SUBTLE POINT: A careful analysis would keep track of the comparisons made in identifying the median of the three candidate elements. You should NOT do this. That is, as in the previous two problems, you should simply add m−1 to your running total of comparisons every time you recurse on a subarray with length m.

func readFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	rows := util.SplitAndTrim(string(r), "\r\n")
	return rows, nil

}

func partition(array []int, left, right int, situation string) ([]int, int) {
	var count int
	if left >= right {
		return array, count
	}
	switch situation {
	case "left":
		break
	case "right":
		selectPivot2(array, left, right-1)
	case "median":
		selectPivot3(array, left, right-1)
	}
	pivot := array[left]
	i := left + 1

	for j := left + 1; j < right; j++ {
		count++
		if array[j] < pivot {
			temp := array[j]
			array[j] = array[i]
			array[i] = temp
			i++
		}
	}
	array[left] = array[i-1]
	array[i-1] = pivot
	array, count1 := partition(array, left, i-1, situation)
	count = count + count1
	array, count2 := partition(array, i, right, situation)
	count = count + count2
	return array, count

}

func selectPivot2(array []int, left, right int) {
	swap(array, left, right)
}

func selectPivot3(array []int, left, right int) {
	var mid int = (right + left) / 2
	if (array[mid] > array[left] && array[mid] < array[right]) || (array[mid] > array[right] && array[mid] < array[left]) {
		swap(array, left, mid)
	} else if (array[right] > array[mid] && array[right] < array[left]) || (array[right] > array[left] && array[right] < array[mid]) {
		swap(array, left, right)
	} else {
	}
}

func swap(array []int, i, j int) {
	temp := array[j]
	array[j] = array[i]
	array[i] = temp
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
	rows, err := readFile("/Users/xkahj/Documents/code/go/src/coursera/stanford-algorithms/part1/week3/QuickSort.txt")
	if err != nil {
		fmt.Println("err--", err)
	}
	// left
	nums := stringToInt(rows)
	_, count1 := partition(nums, 0, len(nums), "left")
	fmt.Println("result1---", count1)
	//right
	nums = stringToInt(rows)
	_, count2 := partition(nums, 0, len(nums), "right")
	fmt.Println("result2---", count2)
	//median
	nums = stringToInt(rows)
	_, count3 := partition(nums, 0, len(nums), "median")
	fmt.Println("result3---", count3)
}
