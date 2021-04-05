package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lxdAssessmentClient/pkg"
	"net/http"

	"github.com/spf13/cobra"
)

// bookCreateCmd represents the bookCreate command
var bookCreateCmd = &cobra.Command{
	Use:   "book",
	Short: "create a book",
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}
		req, _ := http.NewRequest("POST", "http://localhost:8080/1.0/book", nil)

		resp, _ := client.Do(req)

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err.Error())
		}
		var book BookResponse
		var response pkg.HttpResponse
		json.Unmarshal(bodyBytes, &response)
		payload, _ := json.Marshal(response.Payload)
		json.Unmarshal(payload, &book)

		fmt.Println("create book:", book)
	},
}

func init() {
	createCmd.AddCommand(bookCreateCmd)
}
