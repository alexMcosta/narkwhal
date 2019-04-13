package main

import (
  "fmt"
  "encoding/json"

  "github.com/alexMcosta/narkwhal/Digest"
  "github.com/alexMcosta/narkwhal/Process"

  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
  )

func main() {
  // Setup session and service using default credentials found at ~/.aws/
  sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
  }))

  // Service made
  svc := ec2.New(sess)

  // Grab the data from AWS and Marshal it
  // TODO: Add the marshaling into the function?
  data, _ := json.Marshal(digest.GrabAllVolumesData(svc))

  // Output the struct of the marshaled AWS data
  volumes := process.ParseIt(data)
  fmt.Println(volumes)

  //Test to see if I remove all unused EBS volumes on us-east-1
  // TODO: Add the ability to accept flags so when this can be run as a CLI
  // TODO: Possibly create a package that handles the inputs for AWS
  input := &ec2.DeleteVolumeInput{
    AvailabilityZone: aws.String("us-east-1"),
  }
  digest.RemoveOldEBS(svc, input)

  fmt.Println(volumes)
}
