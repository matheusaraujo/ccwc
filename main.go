package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	var flagCountBytes bool
	var flagCountWords bool
	var flagCountLines bool
	var flagCountChars bool

	var rootCmd = &cobra.Command{
		Use:   "ccwc [filename]",
		Short: "Coding Challenge - wc",
		Long:  `A solution for the Coding Challenge wc"`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("error: missing required argument")
			}
			if !flagCountBytes && !flagCountLines && !flagCountWords && !flagCountChars {
				return fmt.Errorf("error: at least one of the flags is required")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			result, err := wc(flagCountBytes, flagCountWords, flagCountLines, flagCountChars, args[0])

			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			fmt.Println(result)
		},
	}

	rootCmd.Flags().BoolVarP(&flagCountBytes, "count-bytes", "c", false, "Count bytes in the file")
	rootCmd.Flags().BoolVarP(&flagCountWords, "count-words", "w", false, "Count words in the file")
	rootCmd.Flags().BoolVarP(&flagCountLines, "count-lines", "l", false, "Count lines in the file")
	rootCmd.Flags().BoolVarP(&flagCountChars, "count-characters", "m", false, "Count characters in the file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
