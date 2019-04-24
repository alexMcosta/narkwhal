package main

import (
	"github.com/alexmcosta/narkwhal/pkg/ingest"
	"github.com/alexmcosta/narkwhal/pkg/process"
)

func main() {
	// Flags
	account, regions, time := process.GetFlags()

	// Take the regions flag and break it down by comma seperation
	sliceOfRegions := process.MultiRegion(regions)

	// Get EBS Volume IDs
	mapOfRegions := ingest.GrabAvailableVolumes(account, sliceOfRegions)

	// Filter ID's based on time given
	filteredVolumeIDs := ingest.FilterVolumesByTime(mapOfRegions, account, time)

	// Show and confirm deletion
	process.ListVolumesAndConfirm(filteredVolumeIDs, account, time)
}
