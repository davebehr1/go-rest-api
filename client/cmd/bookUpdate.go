package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lxdAssessmentClient/pkg"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

// bookUpdateCmd represents the bookUpdate command
var bookUpdateCmd = &cobra.Command{
	Use:   "book",
	Short: "update book",
	Run: func(cmd *cobra.Command, args []string) {
		bookJson, err := cmd.Flags().GetString("book")
		if err != nil {
			fmt.Print(err.Error())
		}
		var newBook Book
		json.Unmarshal([]byte(bookJson), &newBook)

		collection, err := cmd.Flags().GetString("collection")
		if err != nil {
			fmt.Print(err.Error())
		}

		bookId, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Print(err.Error())
		}

		client := &http.Client{}
		req, _ := http.NewRequest("PATCH", "http://localhost:8080/1.0/book", bytes.NewBuffer([]byte(bookJson)))
		req.Header.Add("Content-Type", "application/json")

		q := req.URL.Query()
		if collection != "" {
			q.Add("collection", collection)
		}
		if bookId != -1 {
			Id := strconv.Itoa(bookId)
			q.Add("id", Id)
		}
		req.URL.RawQuery = q.Encode()

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

		fmt.Println("updated book:", book)
	},
}

func init() {
	updateCmd.AddCommand(bookUpdateCmd)
	bookUpdateCmd.Flags().Int("id", -1, "id of book you want to update")
	bookUpdateCmd.Flags().String("book", "", "{title:harry potter,author:jk,description:fantasy}")
	bookUpdateCmd.Flags().String("collection", "", "fantasy")
	bookUpdateCmd.MarkFlagRequired("id")
}
