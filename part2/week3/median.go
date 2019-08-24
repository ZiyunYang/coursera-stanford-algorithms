package main

import (
	"coursera/stanford-algorithms/util"
	"fmt"
	"coursera/stanford-algorithms/heap"
	"strconv"
)

//The goal of this problem is to implement the "Median Maintenance" algorithm (covered in the Week 5 lecture on heap applications). The text file contains a list of the integers from 1 to 10000 in unsorted order; you should treat this as a stream of numbers, arriving one by one. Letting xi denote the ith number of the file, the kth median mk is defined as the median of the numbers x1,…,xk. (So, if k is odd, then mk is ((k+1)/2)th smallest number among x1,…,xk; if k is even, then mk is the (k/2)th smallest number among x1,…,xk.)
//
//In the box below you should type the sum of these 10000 medians, modulo 10000 (i.e., only the last 4 digits). That is, you should compute (m1+m2+m3+⋯+m10000)mod10000.
//
//OPTIONAL EXERCISE: Compare the performance achieved by heap-based and search-tree-based implementations of the algorithm.

type Item struct {
	Value int
}

func (x Item) Less(than heap.Item) bool {
	return x.Value < than.(Item).Value
}
func main() {
	medians, err := util.ReadTXT("/Users/xkahj/Documents/code/go/src/coursera/stanford-algorithms/part2/week3/median.txt", "\n")
	if err != nil {
		fmt.Println("read data err---", err)
	}
	minH := heap.NewMin()
	maxH := heap.NewMax()
	var allMedian []int
	var sum int
	for i, item := range medians {
		median, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println("str to int err---", err)
			return
		}
		medianItem := Item{Value: median}
		if maxH.Len() == 0 {
			maxH.Insert(medianItem)
			fmt.Println(i+1, median)
			sum = sum + median
			allMedian = append(allMedian, median)
			continue
		}
		if medianItem.Less(maxH.Get(0)) {
			maxH.Insert(medianItem)
		} else {
			minH.Insert(medianItem)
		}
		if maxH.Len()-minH.Len() > 1 {
			h := maxH.Extract()
			minH.Insert(h)
		}
		if minH.Len()-maxH.Len() > 1 {
			h := minH.Extract()
			maxH.Insert(h)
		}
		var target int
		if maxH.Len() >= minH.Len() {
			target = maxH.Get(0).(Item).Value
		} else {
			target = minH.Get(0).(Item).Value
		}
		sum = sum + target

		allMedian = append(allMedian, target)
	}
	fmt.Println("sum---", sum%10000)
}
