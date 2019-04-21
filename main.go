package main

import (
	"github.com/alexmcosta/narkwhal/pkg/process"
)

func main() {
	// Flags
	account, region, time := process.GetFlags()

	// if time == "0s" {

	// }

	// ingest.FilterOldVolumes(account, region, time)

	process.ListVolumesAndConfirm(account, region, time)
}
