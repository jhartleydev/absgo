/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"absgo/api"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// ListDataflowCmd represents the ListDataflow command
var ListDataflowCmd = &cobra.Command{
	Use:   "ListDataflow",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		client := &http.Client{}

		fmt.Println("ListDataflow called")

		getURL := api.Base_url + api.DataflowURL

		req, err := http.NewRequest("GET", getURL, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Accept", "application/vnd.sdmx.structure+json")

		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			data := string(bodyBytes)
			fmt.Println(data)
		}
	},
}

func init() {
	rootCmd.AddCommand(ListDataflowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ListDataflowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ListDataflowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
