/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"absgo/api"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// GetDataCmd represents the GetData command
var GetDataCmd = &cobra.Command{
	Use:   "GetData",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		headerType, _ := cmd.Flags().GetString("header")
		structureID, _ := cmd.Flags().GetString("structureID")

		if headerType == "" {
			log.Fatal("no header")
		} else if structureID == "" {
			log.Fatal("no structure ID")
		} else {
			getData(headerType, structureID)
		}
	},
}

func init() {
	rootCmd.AddCommand(GetDataCmd)

	GetDataCmd.PersistentFlags().String("header", "", "add a header to the get request")
	GetDataCmd.PersistentFlags().String("structureID", "", "add a structure ID to the request")
}

var Headers = map[string]string{
	"XmlStructureHeader": api.XmlStructureHeader,
	"JSONHeader":         api.JSONHeader,
	"CSVHeader":          api.CSVHeader,
	"CSVLabelHeader":     api.CSVLabelHeader,
}

func getData(headerType string, structureID string) {
	client := &http.Client{}

	header := Headers[headerType]

	dataFlowId := structureID

	getURL := api.Base_url + api.DataURL + "," + dataFlowId + ",1.0.0"

	req, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", header)

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// create structs for JSON responses
	if response.StatusCode == http.StatusOK {
		if header == "CSVHeader" || header == "CSVLabelHeader" {
			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			println(bodyBytes)

			myString := string(bodyBytes[:])

			println(myString)
		} else if header == "JSONHeader" {
			println("json not ready")
		} else {
			println("do nothing")
		}

		// var root Root

		// json.Unmarshal([]byte(bodyBytes), &root)

		// dataflows := root.Data.Dataflows

		// filteredDataflow := filterDataflow(dataflows, func(d DataflowStruct) bool {
		// 	return d.Id == descTerm
		// })
		// fmt.Println(filteredDataflow)

		// for _, df := range root.Data.Dataflows {
		// 	fmt.Println(df.Id, ",", df.Name)
		// }
	}
}
