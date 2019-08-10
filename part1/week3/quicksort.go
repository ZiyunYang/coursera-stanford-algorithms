package week3

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	rows := splitAndTrim(string(r), "\r\n")
	return rows, nil

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


//func main() {
//	rows, err := readFile("/Users/yangziyun/Downloads/assignment3.txt")
//	if err != nil {
//		fmt.Println("err--", err)
//	}
//	// left
//	nums := stringToInt(rows)
//	count1 := partition(nums, 0, len(nums), "left")
//	fmt.Println("result1---", count1)
//	//right
//	nums = stringToInt(rows)
//	count2 := partition(nums, 0, len(nums), "right")
//	fmt.Println("result2---", count2)
//	//median
//	nums = stringToInt(rows)
//	count3 := partition(nums, 0, len(nums), "median")
//	fmt.Println("result3---", count3)
//}
