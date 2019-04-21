package main

import (
	"github.com/alexmcosta/narkwhal/pkg/process"
)

func main() {
	// Flags
	account, region, time := process.GetFlags()

	// TODO: Refactor Possibly remove the bottom function and bring it back to main. Keep Main() the structure of the application.
	process.ListVolumesAndConfirm(account, region, time)
}
