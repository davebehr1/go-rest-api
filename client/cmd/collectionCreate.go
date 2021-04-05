package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lxdAssessmentClient/pkg"
	"net/http"

	"github.com/spf13/cobra"
)

// collectionCreateCmd represents the collectionCreate command
var collectionCreateCmd = &cobra.Command{
	Use:   "collection",
	Short: "Create a collection",
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}
		req, _ := http.NewRequest("POST", "http://localhost:8080/1.0/collection", nil)

		resp, _ := client.Do(req)

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err.Error())
		}
		var collection CollectionResponse
		var response pkg.HttpResponse
		json.Unmarshal(bodyBytes, &response)
		payload, _ := json.Marshal(response.Payload)
		json.Unmarshal(payload, &collection)

		fmt.Println("create collection:", collection)
	},
}

func init() {
	createCmd.AddCommand(collectionCreateCmd)
}
