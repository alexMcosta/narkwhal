package main

import "strconv"

func timeMeasures(t string) string {
	//Get the last char to see if it is a 'd'
	lastChar := t[len(t)-1:]

	//If it is a 'd' convert left number to int
	if lastChar != "d" {
		return t
	}
	sNum := t[0 : len(t)-1]
	num, _ := strconv.Atoi(sNum)

	//Multiply number by 24
	num = num * 24

	//Convert number to string
	sNum = strconv.Itoa(num) + "h"

	// //Add 'h' to end of string
	return sNum

}
