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

// lscodesCmd represents the lscodes command
var lscodesCmd = &cobra.Command{
	Use:   "lscodes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

This command will reference the API datastructure, and return a list of dimensions from the codelist`,
	Run: func(cmd *cobra.Command, args []string) {
		dataflow, _ := cmd.Flags().GetString("dataflow")

		fmt.Println("*** Codes ***")
		fmt.Println("-------------")
		if dataflow != "" {
			lscodes(dataflow)
		} else {
			fmt.Println("No dataflow provided")
		}

	},
}

func init() {
	rootCmd.AddCommand(lscodesCmd)
	lscodesCmd.PersistentFlags().String("dataflow", "", "Returns code list for this dataflow ID")

}

func lscodes(dataflow string) {
	client := &http.Client{}

	// change this to accept a dataflow to list dimensions of
	getURL := api.Datastructure + dataflow + "/" + api.Codelist

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

		var DSroot Root

		json.Unmarshal([]byte(bodyBytes), &DSroot)

		//var firstDS Codelist = DSroot.Data.Codelists[0]

		//fmt.Println("codelist Id and name:", firstDS.Id, firstDS.Name)
		count := 0
		for _, i := range DSroot.Data.Codelists {
			count += 1
			fmt.Println(i.Id, ",", i.Name)
			for _, e := range i.Codes {
				fmt.Println("--", e.Id, e.Name)
				// for _, annotation := range e.Annotations {
				// 	fmt.Println("--", annotation.Type, annotation.Text)
				// }
			}

		}
		fmt.Println(count)

	}
}

// func getDimensions(dataflow string) {
// 	client := &http.Client{}

// 	// change this to accept a dataflow to list dimensions of
// 	getURL := api.Datastructure + dataflow + "/" + api.Codelist

// 	req, err := http.NewRequest("GET", getURL, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	req.Header.Set("Accept", "application/vnd.sdmx.structure+json")

// 	response, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer response.Body.Close()

// 	if response.StatusCode == http.StatusOK {
// 		bodyBytes, err := io.ReadAll(response.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		var DSroot Root

// 		json.Unmarshal([]byte(bodyBytes), &DSroot)

// 		var firstDS DataStructure = DSroot.Data.DataStructures[0]

// 		fmt.Println("first Id:", firstDS.Id)

// 		for _, i := range firstDS.DsComponents.DimList.Dimensions {
// 			fmt.Println(i.Id, ",", i.Position, ",", i.Type)

// 		}

// 		for _, ds := range firstDS.DsComponents.AttributeList.Attributes {
// 			fmt.Println("attribute ID:", ds.Id)

// 		}

// 	}
//}
