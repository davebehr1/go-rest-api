package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lxdAssessmentClient/pkg"
	"net/http"

	"github.com/spf13/cobra"
)

// collectionDeleteCmd represents the collectionDelete command
var collectionDeleteCmd = &cobra.Command{
	Use:   "collection",
	Short: "delete a collection",
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}
		req, _ := http.NewRequest("DELETE", "http://localhost:8080/1.0/collection", nil)

		resp, _ := client.Do(req)

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err.Error())
		}
		var response pkg.HttpResponse
		json.Unmarshal(bodyBytes, &response)

		fmt.Println(response.Payload)
	},
}

func init() {
	deleteCmd.AddCommand(collectionDeleteCmd)
}
