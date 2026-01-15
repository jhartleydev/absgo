/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"absgo/api"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// ListDataflowCmd represents the ListDataflow command
var ListDataflowCmd = &cobra.Command{
	Use:   "ListDataflow",
	Short: "Lists available dataflows",
	Long: `Lists all available dataflows, unless using the desc flag, which takes a dataflow ID 
and provides a specific description of that dataflow, such as:

absgo ListDataflow --desc WPI

Which returns details on the wage price index.
`,
	Run: func(cmd *cobra.Command, args []string) {
		descTerm, _ := cmd.Flags().GetString("desc")

		if descTerm != "" {
			describeDataflow(descTerm)
		} else {
			getDataflowItems()
		}
	},
}

func init() {
	rootCmd.AddCommand(ListDataflowCmd)

	ListDataflowCmd.PersistentFlags().String("desc", "", "Describe a specific Dataflow Item")

}

func getDataflowItems() {
	client := &http.Client{}

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

		var root Root

		json.Unmarshal([]byte(bodyBytes), &root)

		for _, df := range root.Data.Dataflows {
			fmt.Println(df.Id, ",", df.Name)
		}
	}
}

func filterDataflow(dataflows []DataflowStruct, test func(DataflowStruct) bool) []DataflowStruct {
	var result []DataflowStruct
	for _, p := range dataflows {
		if test(p) {
			result = append(result, p)
		}
	}
	return result
}

func describeDataflow(descTerm string) {
	client := &http.Client{}

	// fmt.Println("ListDataflow called")

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

		var root Root

		json.Unmarshal([]byte(bodyBytes), &root)

		dataflows := root.Data.Dataflows

		filteredDataflow := filterDataflow(dataflows, func(d DataflowStruct) bool {
			return d.Id == descTerm
		})
		fmt.Println(filteredDataflow)

		// for _, df := range root.Data.Dataflows {
		// 	fmt.Println(df.Id, ",", df.Name)
		// }
	}
}
