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
	"os"

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
		filename, _ := cmd.Flags().GetString("filename")

		if headerType == "" {
			log.Fatal("no header")
		} else if structureID == "" {
			log.Fatal("no structure ID")
		} else if headerType[0:3] == "CSV" {
			csvbytes := getCSVData(headerType, structureID)
			if filename != "" {
				os.WriteFile(filename+".csv", csvbytes, 0644)
				fmt.Printf("outputting %s.csv to file...", filename)
			}
		} else if headerType == "JSONHeader" {
			getJSONData(headerType, structureID)
		}
	},
}

func init() {
	rootCmd.AddCommand(GetDataCmd)

	GetDataCmd.PersistentFlags().String("header", "", "add a header to the get request")
	GetDataCmd.PersistentFlags().String("structureID", "", "add a structure ID to the request")
	GetDataCmd.PersistentFlags().String("filename", "", "output to file")
}

var Headers = map[string]string{
	"XmlStructureHeader": api.XmlStructureHeader,
	"JSONHeader":         api.JSONHeader,
	"CSVHeader":          api.CSVHeader,
	"CSVLabelHeader":     api.CSVLabelHeader,
}

func getJSONData(headerType string, structureID string) {
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
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var result interface{}

		json.Unmarshal(bodyBytes, &result)

		fmt.Println(result)

		myString := string(bodyBytes[:])
		fmt.Println(myString)
	}
}

func getCSVData(headerType string, structureID string) []byte {
	client := &http.Client{}

	var header string = Headers[headerType]

	var getURL string = api.Base_url + api.DataURL + "," + structureID + ",1.0.0"

	fmt.Println(getURL)

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

	var bodyBytes []byte

	// create structs for JSON responses
	if response.StatusCode == http.StatusOK {
		bodyBytes, err = io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		myString := string(bodyBytes[:])
		fmt.Println(myString)

	}
	return bodyBytes
}
