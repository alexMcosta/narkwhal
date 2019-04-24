package process

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// MultiRegion will sort through comma seperated regions and make a slice of strings
func MultiRegion(regions string) []string {

	var sliceOfRegions []string

	if strings.Contains(regions, ",") {
		sliceOfRegions := strings.Split(regions, ",")
		return sliceOfRegions
	}

	sliceOfRegions = append(sliceOfRegions, regions)

	return sliceOfRegions

}

// confirm makes sure that the user either gives a yes or no answer
func confirm() bool {
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadRune()

	switch text {
	case 'y', 'Y':
		return true
	case 'n', 'N':
		return false
	default:
		fmt.Println("Please type y or n and then press enter: ")
		return confirm()
	}
}
