package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lxdAssessmentClient/pkg"
	"net/http"

	"github.com/spf13/cobra"
)

// bookDeleteCmd represents the bookDelete command
var bookDeleteCmd = &cobra.Command{
	Use:   "book",
	Short: "delete a book",
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}
		req, _ := http.NewRequest("DELETE", "http://localhost:8080/1.0/book", nil)

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
	deleteCmd.AddCommand(bookDeleteCmd)
}
