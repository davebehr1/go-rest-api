package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lxdAssessmentClient/pkg"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

// collectionUpdateCmd represents the collectionUpdate command
var collectionUpdateCmd = &cobra.Command{
	Use:   "collection",
	Short: "update collection",
	Run: func(cmd *cobra.Command, args []string) {
		cId, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Print(err.Error())
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Print(err.Error())
		}
		client := &http.Client{}
		req, _ := http.NewRequest("PATCH", "http://localhost:8080/1.0/collection", nil)
		req.Header.Add("Content-Type", "application/json")
		q := req.URL.Query()
		if cId != -1 {
			Id := strconv.Itoa(cId)
			q.Add("id", Id)
		}
		if name != "" {
			q.Add("name", name)
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

		fmt.Println("updated collection:", collection.Name)
	},
}

func init() {
	rootCmd.AddCommand(collectionUpdateCmd)
	updateCmd.AddCommand(collectionUpdateCmd)
	collectionUpdateCmd.Flags().Int("id", -1, "id of collection to update")
	collectionCreateCmd.Flags().String("name", "", "fantasy")
	collectionCreateCmd.MarkFlagRequired("id")
	collectionCreateCmd.MarkFlagRequired("name")
}
