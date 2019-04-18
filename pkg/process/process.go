package process

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// GetFlags does the work of getting the flags!
func GetFlags() (string, string) {
	accountFlag := flag.String("account", "default", "Lets you select witch AWS account you would like to make changes to")
	regionFlag := flag.String("region", "us-east-1", "Lets you select which region you would like to run Narkwhal on")
	flag.Parse()
	fmt.Printf("account: %s, region: %s\n", *accountFlag, *regionFlag)
	return *accountFlag, *regionFlag
}

// Confirm makes sure that the user either gives a yes or no answer
func Confirm() bool {
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadRune()

	switch text {
	case 'y', 'Y':
		return true
	case 'n', 'N':
		return false
	default:
		fmt.Println("Please type y or n and then press enter: ")
		return Confirm()
	}
}
