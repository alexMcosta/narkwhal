package process

import (
	"encoding/json"
	"fmt"
)

func ParseIt(data []byte) interface{} {

	type Volumes struct {
		State    string `JSON:"State"`
		VolumeId string `JSON:"VolumeId"`
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
