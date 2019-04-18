package process

import (
	"flag"
	"fmt"
)

// Get Flags
func GetFlags() (string, string) {
	accountFlag := flag.String("account", "default", "Lets you select witch AWS account you would like to make changes to")
	regionFlag := flag.String("region", "us-east-1", "Lets you select which region you would like to run Narkwhal on")
	flag.Parse()
	fmt.Printf("account: %s, region: %s\n", *accountFlag, *regionFlag)
	return *accountFlag, *regionFlag
}
