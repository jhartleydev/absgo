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
	Short: "Lists codes for a dataflow",
	Long: `Takes a single dataflow id as an argument following the "dataflow" flag, and returns a codelist of available codes 
	for further querying. For example:

usage: absgo lscodes --dataflow WPI

`,
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
	lscodesCmd.PersistentFlags().StringP("dataflow", "d", "", "Returns code list for this dataflow ID")

}

func lscodes(dataflow string) {
	client := &http.Client{}

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

		count := 0
		for _, i := range DSroot.Data.Codelists {
			count += 1
			fmt.Println(i.Id, ",", i.Name)
			for _, e := range i.Codes {
				fmt.Println("--", e.Id, e.Name)
			}

		}
		fmt.Println(count)

		fmt.Println(len(DSroot.Data.DataStructures[0].DsComponents.DimList.Dimensions) + len(DSroot.Data.DataStructures[0].DsComponents.DimList.TimeDimensions))

		for _, i := range DSroot.Data.DataStructures[0].DsComponents.DimList.Dimensions {
			fmt.Println(i.Id, ", position: ", i.Position)

		}

		for _, i := range DSroot.Data.DataStructures[0].DsComponents.DimList.TimeDimensions {
			fmt.Println(i.Id, ", position: ", i.Position)

		}

	}
}
