package main

import "strings"

// checkMultiInput checks if multiple regions or special flags are entered.
// If multiple regions are entered it will make a slice seperated at the comma's.
// If one region is entered it will store the single region in a slice
func checkMultiInput(reg string) []string {
	var sliceReg []string

	if strings.Contains(reg, ",") {
		sliceReg = strings.Split(reg, ",")
		return sliceReg
	}

	sliceReg = append(sliceReg, reg)
	return sliceReg

}
