package main

import (
	"encoding/json"

	"github.com/alexmcosta/narkwhal/ingest"
	"github.com/alexmcosta/narkwhal/process"
)

func main() {
	// Grab all available volumes and marshal them
	volumeIDs := ingest.GrabAvailableVolumesIDs()
	data, _ := json.Marshal(volumeIDs)

	// Process the ID's And make them structs
	volumes := process.GetStructOfVolumes(data)

	//Remove any availble structs
	ingest.RemoveAvailableEBS(volumes)
}
