package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func setOptions(countBytes, countWords, countLines, countChars bool) Options {
	if !countBytes && !countWords && !countLines && !countChars {
		return Options{
			CountBytes: true,
			CountWords: true,
			CountLines: true,
			CountChars: false,
		}
	}
	return Options{
		CountBytes: countBytes,
		CountWords: countWords,
		CountLines: countLines,
		CountChars: countChars,
	}
}

func executeWC(args []string, options Options) {
	if len(args) <= 0 {
		result, err := wc(options, nil)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(result)

	} else {
		result, err := wc(options, &args[0])

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(result)
	}

}

func main() {
	var (
		flagCountBytes bool
		flagCountWords bool
		flagCountLines bool
		flagCountChars bool
	)

	rootCmd := &cobra.Command{
		Use:   "ccwc [filename]",
		Short: "Coding Challenge - wc",
		Long:  `A solution for the Coding Challenge wc`,
		Run: func(cmd *cobra.Command, args []string) {
			options := setOptions(flagCountBytes, flagCountWords, flagCountLines, flagCountChars)
			executeWC(args, options)
		},
	}

	rootCmd.Flags().BoolVarP(&flagCountBytes, "count-bytes", "c", false, "Count bytes in the file")
	rootCmd.Flags().BoolVarP(&flagCountWords, "count-words", "w", false, "Count words in the file")
	rootCmd.Flags().BoolVarP(&flagCountLines, "count-lines", "l", false, "Count lines in the file")
	rootCmd.Flags().BoolVarP(&flagCountChars, "count-characters", "m", false, "Count characters in the file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Command execution error: %v\n", err)
		os.Exit(1)
	}
}
