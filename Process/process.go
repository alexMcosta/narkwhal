package process

import (
  "fmt"
  "encoding/json"
  "time"
)

func ParseIt(data []byte) interface{}{
  // Turn data into a string for easier ingesting
  //var dataByteArray = []byte(data.GoString())

  //Make a Volume struct for JSON parsing
  type Attachments struct {
    AttachTime          *time.Time
    DeleteOnTermination bool
    Device              string
    InstanceId          string
    State               string
    VolumeId            string
  }

  type Tags struct {
    Key   string `JSON:"Key"`
    Value string `JSON:"Value"`
  }

  type Volumes struct {
    Attachments      []Attachments
    AvailabilityZone string     `JSON:"AvailabilityZone"`
    CreateTime       *time.Time `JSON:"CreateTime"`
    Encrypted        bool       `JSON:"Encrypted"`
    Iops             int        `JSON:"Iops"`
    Size             int        `JSON:"Size"`
    SnapshotId       string     `JSON:"SnapshotId"`
    State            string     `JSON:"State"`
    Tags             []Tags
    VolumeId         string     `JSON:"VolumeId"`
    VolumeType       string     `JSON:"VolumeType"`
  }

  type Volume struct {
    Volumes []Volumes
	}

  var dataVolumes Volume

  err := json.Unmarshal(data, &dataVolumes)
  if err != nil {
    fmt.Println("error:", err)
  }

  //fmt.Printf("%+v", dataVolumes)
  return dataVolumes
}
