package main

import (
	"github.com/alexmcosta/narkwhal/pkg/cloud"
)

func main() {
	// Flags
	acc, reg, time := getFlags()

	// Filter time for days and weeks
	time = timeMeasures(time)

	// Take the flag inputs and break it down by comma seperation
	sliceReg := multiRegion(reg)
	sliceAcc := multiAccount(acc)

	// Get EBS Volume IDs
	mapReg := cloud.GetVolumes(sliceAcc, sliceReg)

	// Filter ID's based on time given
	filVol := cloud.NonActiveVolumes(mapReg, sliceAcc, time)

	// Show and confirm deletion
	showVolumes(filVol, sliceAcc, time)
}
