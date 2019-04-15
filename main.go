package main

import (
	"encoding/json"
	"fmt"

	digest "github.com/alexMcosta/narkwhal/Ingest"
	process "github.com/alexMcosta/narkwhal/Process"
)

func main() {
	// Setup session and service using default credentials found at ~/.aws/
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	// // Service made
	// svc := ec2.New(sess)

	// Grab the data from AWS and Marshal it
	// TODO: Add the marshaling into the function?
	data, _ := json.Marshal(digest.GrabAllVolumesData())

	// Output the struct of the marshaled AWS data
	volumes := process.ParseIt(data)

	//Test to see if I remove all unused EBS volumes on us-east-1
	// TODO: Add the ability to accept flags so when this can be run as a CLI
	// TODO: Possibly create a package that handles the inputs for AWS

	fmt.Println(volumes)
}
