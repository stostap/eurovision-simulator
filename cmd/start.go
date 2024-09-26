/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"eurovision-simulator/controllers"
	"eurovision-simulator/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the DB
		db := utils.InitializeDB()

		// Get the Eurovision session
		eurovision := utils.GetEurovisionSession()

		// Initialize the contest here
		fmt.Println("Starting Eurovision:", eurovision.Year)

		eurovisionController := controllers.NewEurovisionController(db)
		_, err := eurovisionController.StartEurovision()
		if err != nil {
			fmt.Println("Error starting Eurovision:", err)
		} else {
			fmt.Println("Eurovision contest started")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
