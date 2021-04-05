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
		c, err := cmd.Flags().GetString("collection")

		if err != nil {
			fmt.Print(err.Error())
		}
		client := &http.Client{}
		req, _ := http.NewRequest("POST", "http://localhost:8080/1.0/collection", nil)
		req.Header.Add("Content-Type", "application/json")
		q := req.URL.Query()
		if c != "" {
			q.Add("collection", c)
		}
		req.URL.RawQuery = q.Encode()

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

		fmt.Println("created collection:", collection.Name)
	},
}

func init() {
	createCmd.AddCommand(collectionCreateCmd)
	collectionCreateCmd.Flags().String("collection", "", "fantasy")
	collectionCreateCmd.MarkFlagRequired("collection")
}
