/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"absgo/api"
	"fmt"

	"github.com/spf13/cobra"

	//"bufio"
	"net/http"
)

// TestAPICmd represents the TestAPI command
var TestAPICmd = &cobra.Command{
	Use:   "TestAPI",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("TestAPI called")

		resp, err := http.Get(api.Base_url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("Response status:", resp.Status)

	},
}

func init() {
	rootCmd.AddCommand(TestAPICmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// TestAPICmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// TestAPICmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
