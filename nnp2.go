package main

import "errors"

func GetLastButOne(list []int) (int, error){
	if(list == nil){
		return 0, errors.New("nilError")
	}
	l := len(list)
	if(l < 2){
		return 0, errors.New("EmptyError")
	}
	return list[l-2], nil
}

