package main

import (
	"encoding/json"
	"fmt"

	"github.com/alexmcosta/narkwhal/ingest"
	"github.com/alexmcosta/narkwhal/process"
)

func main() {
	cats := ingest.GrabAvailableVolumesIDs()
	data, _ := json.Marshal(cats)

	// Output the struct of the marshaled AWS data
	volumes := process.ParseIt(data)

	fmt.Println(volumes)
}
