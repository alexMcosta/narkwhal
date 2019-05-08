package main

import (
	"github.com/alexmcosta/narkwhal/pkg/cloud"
)

func main() {
	// Flags
	acc, reg, time := getFlags()

	// Take the regions flag and break it down by comma seperation
	sliceReg := multiRegion(reg)

	// Get EBS Volume IDs
	mapReg := cloud.GetVolumes(acc, sliceReg)

	// Filter ID's based on time given
	filVol := cloud.NonActiveVolumes(mapReg, acc, time)

	// Show and confirm deletion
	showVolumes(filVol, acc, time)
}
