package process

import (
	"encoding/json"
	"fmt"
)

// Create a struct of Volume that is a volume of volume IDs
// Yes the naming is bad but I am at a loss of what to name a volume of volumes
type Volumes struct {
	VolumeId string `JSON:"VolumeId"`
}

type Volume struct {
	Volumes []Volumes
}

// func (v Volumes) getVolumeId() string {
// 	return v.VolumeId
// }

// GetMapOfID Gets a map of the IDS given
func GetMapOfIDs(data []byte) []Volumes {

	var dataVolumes Volume

	err := json.Unmarshal(data, &dataVolumes)
	if err != nil {
		fmt.Println("error:", err)
	}

	//fmt.Printf("%+v", dataVolumes)

	return dataVolumes.Volumes
}
