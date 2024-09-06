package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var countBytes bool

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ccwc -c [filename]",
		Short: "Coding Challenge - wc",
		Long:  `A solution for the Coding Challenge wc"`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("error: missing required argument")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			result, err := wc(countBytes, args[0])

			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			fmt.Println(result)
		},
	}

	rootCmd.Flags().BoolVarP(&countBytes, "count-bytes", "c", false, "Count bytes in the file")
	rootCmd.MarkFlagRequired("count-bytes")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
