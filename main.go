package main

import (
	"github.com/alexmcosta/narkwhal/pkg/ingest"
	"github.com/alexmcosta/narkwhal/pkg/process"
)

func main() {
	// Flags
	acc, reg, time := process.GetFlags()

	// Take the regions flag and break it down by comma seperation
	sliceReg := multiRegion(reg)

	// Get EBS Volume IDs
	mapReg := ingest.GrabAvailableVolumes(acc, sliceReg)

	// Filter ID's based on time given
	filVol := ingest.FilterVolumesByTime(mapReg, acc, time)

	// Show and confirm deletion
	showVolumes(filVol, acc, time)
}
