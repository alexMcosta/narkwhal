package process

import (
	"flag"
	"fmt"
)

// GetFlags does the work of getting the flags!
func GetFlags() (string, string, string) {
	accountFlag := flag.String("account", "default", "Lets you select witch AWS account you would like to make changes to")
	regionsFlag := flag.String("regions", "us-east-1", "Lets you select which regions you would like to run Narkwhal on")
	timeFlag := flag.String("time", "0s", "Lets you select the amount of time a volume has been available based on MS, seconds, and Hours")
	flag.Parse()

	fmt.Printf("account: %s, region: %v, Not used within: %s\n", *accountFlag, *regionsFlag, *timeFlag)
	return *accountFlag, *regionsFlag, *timeFlag
}
