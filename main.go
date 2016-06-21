package main

import (
	"fmt"
	"os"

	"github.com/nii236/btcer/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(-1)
	}
}
