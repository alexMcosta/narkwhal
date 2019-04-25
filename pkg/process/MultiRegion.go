package process

import (
	"strings"
)

// MultiRegion will sort through comma seperated regions and make a slice of strings
func MultiRegion(regions string) []string {

	var sliceOfRegions []string
	var sliceOfInput []string

	// Check for multiple input
	sliceOfInput = checkMultipleInput(regions)

	// Run through slice of inputs to concat all regions desired
	for _, input := range sliceOfInput {
		switch input {
		case "AM":
			sliceOfRegions = append(sliceOfRegions, america(sliceOfRegions)...)
		case "AP":
			sliceOfRegions = append(sliceOfRegions, asiaPacific(sliceOfRegions)...)
		case "EU":
			sliceOfRegions = append(sliceOfRegions, europe(sliceOfRegions)...)
		case "ALL":
			sliceOfRegions = allRegions()
		default:
			sliceOfRegions = append(sliceOfRegions, input)
		}
	}

	return sliceOfRegions

}

func checkMultipleInput(regions string) []string {
	var sliceOfRegions []string

	if strings.Contains(regions, ",") {
		sliceOfRegions = strings.Split(regions, ",")
		return sliceOfRegions
	}

	sliceOfRegions = append(sliceOfRegions, regions)
	return sliceOfRegions

}

func asiaPacific(sliceOfRegions []string) []string {
	regions := []string{
		"ap-south-1",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-southeast-1",
		"ap-southeast-2",
	}
	sliceOfRegions = append(sliceOfRegions, regions...)
	return sliceOfRegions
}

func america(sliceOfRegions []string) []string {
	regions := []string{
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		"ca-central-1",
		"sa-east-1",
	}
	sliceOfRegions = append(sliceOfRegions, regions...)
	return sliceOfRegions
}

func europe(sliceOfRegions []string) []string {
	regions := []string{
		"eu-central-1",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"eu-north-1",
	}
	sliceOfRegions = append(sliceOfRegions, regions...)
	return sliceOfRegions
}

func allRegions() []string {
	var regions []string
	regions = asiaPacific(america(europe(regions)))
	return regions
}
