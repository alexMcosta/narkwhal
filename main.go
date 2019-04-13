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
  // First let us configure the session
  // This session will take the default credentials found at ~/.aws/
  sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
    }))

  // Create a service with our session
  svc := ec2.New(sess)

  // Grab the data from AWS and Marshal it
  data, _ := json.Marshal(digest.GrabAllVolumesData(svc))

  // Output the struct of the marshaled AWS data
  volumes := process.ParseIt(data)

  fmt.Println(volumes)

}
