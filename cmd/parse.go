/*
Copyright © 2024 Timur Kulakov totusfloreodev@proton.me
*/
package cmd

import (
	"fmt"
	"github.com/Totus-Floreo/goconstdoc/internal/app"
	"github.com/spf13/cobra"
	"log"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Command for parsing constants from the file",
	Long:  `This command is used to parse constants from the go file and generate documentation html table for them.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		isOutputToTerminal, _ := cmd.Flags().GetBool("cmd")

		if !isOutputToTerminal {
			fmt.Println("parsing constants from the file...")
		}

		table, err := app.ParseGoFile(path)
		if err != nil {
			log.Fatal(err)
		}

		if cmd.Flags().Changed("output") {
			outputPath, err := cmd.Flags().GetString("output")
			if err != nil {
				log.Fatal(err)
			}

			if outputPath != "" {
				overwrite, _ := cmd.Flags().GetBool("overwrite")
				err = app.SaveHTMLToFile(table, outputPath, overwrite)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatal("output path is empty")
			}
		}

		if isOutputToTerminal {
			err = app.WriteToConsole(table)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	parseCmd.PersistentFlags().StringP("path", "p", "", "Path to the go file")
	parseCmd.MarkPersistentFlagRequired("path")

	parseCmd.Flags().Bool("cmd", true, "Output to the command line")

	parseCmd.Flags().StringP("output", "o", "", "Output file for the documentation")
	parseCmd.Flags().Bool("overwrite", false, "Overwrite the file if it exists")
}
