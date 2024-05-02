/*
Copyright © 2024 Timur Kulakov totusfloreodev@proton.me
*/
package cmd

import (
	"fmt"
	"github.com/Totus-Floreo/goconstdoc/internal/app"
	"github.com/Totus-Floreo/goconstdoc/internal/constant"
	"github.com/Totus-Floreo/goconstdoc/internal/domain"
	"github.com/spf13/cobra"
	"log"
)

var InteractionEnum = &domain.EnumValue{Allowed: []string{constant.Builtin, constant.Merge, constant.Overwrite}, Value: constant.Builtin}

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Command for parsing constants from the file",
	Long:  `This command is used to parse constants from the go file and generate documentation html table for them.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		isOffOutputToTerminal, _ := cmd.Flags().GetBool("nocmd")
		interaction := InteractionEnum.String()

		if isOffOutputToTerminal {
			fmt.Println("parsing constants from the file...")
		}

		table, err := app.ParseGoFile(path, interaction)
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

		if !isOffOutputToTerminal {
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

	parseCmd.Flags().Bool("nocmd", false, "Disable output to the command line")

	parseCmd.Flags().StringP("output", "o", "", "Output file for the documentation")
	parseCmd.Flags().Bool("overwrite", false, "Overwrite the file if it exists")

	parseCmd.Flags().VarP(InteractionEnum, "interaction", "i", "Type of interaction with built-in values\nallowed values are builtin, merge, overwrite")
}
