package main

import (
	"strconv"
	"errors"
	"fmt"
)
// this method is working with recuirsion
// it splits the operation string when find any supported operation
// operations are ordered from right side to left 
// operations periorities are (- +) then  (* /)

func calculate(operation string) (float64, error) {
	//first validate string to be valid operation string
	val, err := strconv.ParseFloat(operation, 64)
	if err == nil {
		return val, nil
	}
	//look for - or + first
	for i:=len(operation)-1; i>=0; i--{
		if operation[i] == '-' || operation[i] == '+'{
			rightSide, rerr := calculate(operation[:i])
			leftSide, lerr := calculate(operation[i+1:])
			if rerr != nil{
				return -1, rerr
			} else if lerr != nil {
				return -1, lerr
			} else if operation[i] == '-' {
				return rightSide - leftSide ,nil
			} else if operation[i] == '+' {		
				return rightSide + leftSide, nil
			}
		}
	}
	//then look for / or *
	for i:=len(operation)-1; i>=0; i--{
		if operation[i] == '*' || operation[i] == '/'{
			rightSide, rerr := calculate(operation[:i])
			leftSide, lerr := calculate(operation[i+1:])
			if rerr != nil{
				return -1, rerr
			} else if lerr != nil {
				return -1, lerr
			} else if operation[i] == '*' {
				return rightSide * leftSide ,nil
			} else if operation[i] == '/' {		
				return rightSide / leftSide, nil
			}
		}
	}
	return 0, errors.New("operation error near " + operation)
}
