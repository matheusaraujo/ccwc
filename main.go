package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ccwc",
		Short: "Coding Challenge - wc",
		Long:  `A solution for the Coding Challenge wc!"`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(wc())
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
