package util

import (
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadTXT(filePath,split string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	rows := SplitAndTrim(string(r), split)
	////fmt.Println("len(rows)---",len(rows))
	//fmt.Println("row[0]",rows[0])
	//fmt.Println("row[0]",rows[len(rows)-1])
	return rows, nil

}

func SplitAndTrim(oldString, split string) []string {
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

func StringToInt(a []string) ([]int, error) {
	var nums []int
	for _, item := range a {
		num, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func ReadExcel(filePath string) ([][]string, error) {
	file, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	var rowValue [][]string
	for _, sheet := range file.Sheets { // 遍历所有工作表 本题中只有一个工作表
		for _, row := range sheet.Rows { // 遍历所有行
			var cellValue []string
			for _, cell := range row.Cells { // 遍历所有单元格
				if cell.Value != "" {
					cellValue = append(cellValue, cell.Value)
				}
			}
			rowValue = append(rowValue, cellValue)
		}
	}
	return rowValue, nil
}
