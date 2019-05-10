package main

var asiaPac = []string{
	"ap-south-1",
	"ap-northeast-1",
	"ap-northeast-2",
	"ap-southeast-1",
	"ap-southeast-2",
}

var america = []string{
	"us-east-1",
	"us-east-2",
	"us-west-1",
	"us-west-2",
	"ca-central-1",
	"sa-east-1",
}

var europe = []string{
	"eu-central-1",
	"eu-west-1",
	"eu-west-2",
	"eu-west-3",
	"eu-north-1",
}

// multiRegion will sort through comma seperated regions and make a slice of strings.
func multiRegion(reg string) []string {

	var sliceReg []string
	sliceInp := checkMultiInput(reg)

	// Run through slice of inputs to concat all regions desired
	for _, input := range sliceInp {
		switch input {
		case "AM":
			sliceReg = append(sliceReg, america...)
		case "AP":
			sliceReg = append(sliceReg, asiaPac...)
		case "EU":
			sliceReg = append(sliceReg, europe...)
		case "ALL":
			asiaAmer := append(asiaPac, america...)
			asiaAmerEur := append(asiaAmer, europe...)
			sliceReg = append(sliceReg, asiaAmerEur...)
		default:
			sliceReg = append(sliceReg, input)
		}
	}

	return sliceReg

}
