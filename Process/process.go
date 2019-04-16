package process

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // Create a struct of Volume that is a volume of volumes. Right now all we need are IDs
// // Yes the naming is bad but I am at a loss of what to name a volume of volumes
// type Volumes struct {
// 	VolumeId string `JSON:"VolumeId"`
// }

// type Volume struct {
// 	Volumes []Volumes
// }

// // GeStructOfVolumes Gets a struct of volumes
// func GetStructOfVolumes(data []byte) []Volumes {

// 	var dataVolumes Volume

// 	err := json.Unmarshal(data, &dataVolumes)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}

// 	return dataVolumes.Volumes
// }
