package zigzagConversion

import "fmt"

func ZigzagConversion(s string, numRows int) string {
	resultList := make([]string, numRows)
	line := 0
	count := 0
	row := 0
	if numRows <= 1 {
		return s
	}
	for _, v := range s {
		count++
		if count >= numRows {
			line++
		}
		if count == numRows*2-1 {
			count = 1
			line = 0
			row = 0
		}
		if line == 0 {
			resultList[row] = resultList[row] + string(v)
		} else {
			resultList[numRows-line] = resultList[numRows-line] + string(v)
		}
		row++

	}
	var result string
	for _, runes := range resultList {
		result = result + runes
	}
	return string(result)
}

func ZigzagConversionPolling(s string, numRows int) string {
	resultList := make([]string, numRows)

	for k, v := range s {
		line := (k) % (numRows + 2)
		println("start line: ", line, k)
		if line >= numRows {
			//line = numRows - (line - numRows) -1
			line = numRows + 0 - (line - numRows + 1) - 1
			println("line-numRows: ", line)
			resultList[line] = resultList[line] + string(v)
			continue
		}
		println("line: ", line)
		resultList[line] = resultList[line] + string(v)

	}
	fmt.Println(resultList)
	var result string
	for _, runes := range resultList {
		result = result + runes
	}
	return string(result)
}
