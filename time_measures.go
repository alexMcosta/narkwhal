package main

import "strconv"

func timeMeasures(t string) string {
	//Get the last char to see if it is a 'd'
	lastChar := t[len(t)-1:]

	sNum := t[0 : len(t)-1]
	num, _ := strconv.Atoi(sNum)

	switch lastChar {
	case "d":
		num = num * 24
	case "w":
		num = num * 168
	default:
		return t
	}

	//Convert number to string
	sNum = strconv.Itoa(num) + "h"

	//Add 'h' to end of string
	return sNum

}
