package process

import (
	"encoding/json"
	"fmt"
)

type Volumes struct {
	VolumeId string `JSON:"VolumeId"`
}

type Volume struct {
	Volumes []Volumes
}

func (v Volumes) getVolumeId() string {
	return v.VolumeId
}

func ParseIt(data []byte) interface{} {

	var dataVolumes Volume

	err := json.Unmarshal(data, &dataVolumes)
	if err != nil {
		fmt.Println("error:", err)
	}

	//fmt.Printf("%+v", dataVolumes)
	return dataVolumes
}
