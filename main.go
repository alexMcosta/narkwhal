package main

import (
	"encoding/json"

	"github.com/alexmcosta/narkwhal/ingest"
	"github.com/alexmcosta/narkwhal/process"
)

func main() {
	volumeIDs := ingest.GrabAvailableVolumesIDs()
	data, _ := json.Marshal(volumeIDs)

	// Output the struct of the marshaled AWS data
	volumes := process.GetMapOfIDs(data)
	ingest.RemoveAvailableEBS(volumes)
}
