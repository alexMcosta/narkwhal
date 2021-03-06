package main

import (
	"flag"
	"fmt"
)

// GetFlags does the work of getting the flags!
func getFlags() (string, string, string) {
	accFlag := flag.String("account", "default", "Lets you select witch AWS account you would like to make changes to")
	regFlag := flag.String("regions", "us-east-1", "Lets you select which regions you would like to run Narkwhal on")
	timeFlag := flag.String("time", "0s", "Lets you select the amount of time a volume has been available based on MS, seconds, and Hours")
	flag.Parse()

	fmt.Printf("account: %s, regions: %v, Not used within: %s\n", *accFlag, *regFlag, *timeFlag)
	return *accFlag, *regFlag, *timeFlag
}
