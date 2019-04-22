package main

import (
	"github.com/alexmcosta/narkwhal/pkg/ingest"
	"github.com/alexmcosta/narkwhal/pkg/process"
)

func main() {
	// Flags
	account, region, time := process.GetFlags()

	//Get EBS Volume Data
	EBSVolumes := ingest.GrabAvailableVolumes(account, region)

	// Filter data for the Volume ID's
	volumeIDs := process.GetSliceOfIDs(EBSVolumes)

	// Filter ID's based on time given
	filteredVolumeIDs := ingest.FilterVolumesByTime(volumeIDs, account, region, time)

	// Show and confirm deletion
	process.ListVolumesAndConfirm(filteredVolumeIDs, account, region, time)
}
