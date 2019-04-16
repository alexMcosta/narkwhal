package main

import (
	"encoding/json"
	"fmt"

	"github.com/alexmcosta/narkwhal/ingest"
	"github.com/alexmcosta/narkwhal/process"
)

func main() {
	volumeIDs := ingest.GrabAvailableVolumesIDs()
	data, _ := json.Marshal(volumeIDs)

	// Output the struct of the marshaled AWS data
	volumes := process.GetMapOfIDs(data)

	fmt.Println(volumes)

	ingest.RemoveAvailableEBS(volumes)
}
