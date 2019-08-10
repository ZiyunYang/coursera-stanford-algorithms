package main

import (
	"coursera/stanford-algorithms/util"
	"fmt"
	"coursera/stanford-algorithms/heap"
	"strconv"
)

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
