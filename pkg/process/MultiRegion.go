package process

import "strings"

// MultiRegion will sort through comma seperated regions and make a slice of strings
func MultiRegion(regions string) []string {

	var sliceOfRegions []string

	if regions == "ALL" {
		sliceOfRegions = []string{
			"us-east-1",
			"us-east-2",
			"us-west-1",
			"us-west-2",
			"ap-south-1",
			"ap-northeast-1",
			"ap-northeast-2",
			//"ap-northeast-3",
			"ap-southeast-1",
			"ap-southeast-2",
			"ca-central-1",
			//"cn-north-1",
			//"cn-northwest-1",
			"eu-central-1",
			"eu-west-1",
			"eu-west-2",
			"eu-west-3",
			"eu-north-1",
			"sa-east-1",
		}
		return sliceOfRegions

	}

	if strings.Contains(regions, ",") {
		sliceOfRegions := strings.Split(regions, ",")
		return sliceOfRegions
	}

	sliceOfRegions = append(sliceOfRegions, regions)

	return sliceOfRegions

}
