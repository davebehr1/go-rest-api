package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lxdAssessmentClient/pkg"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

type Book struct {
	Title       string
	Description string
	Author      string
	PublishedAt time.Time
}

// bookCreateCmd represents the bookCreate command
var bookCreateCmd = &cobra.Command{
	Use:   "book",
	Short: "create a book",
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

		client := &http.Client{}
		req, _ := http.NewRequest("POST", "http://localhost:8080/1.0/book", bytes.NewBuffer([]byte(bookJson)))
		req.Header.Add("Content-Type", "application/json")

		q := req.URL.Query()
		if collection != "" {
			q.Add("collection", collection)
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

		fmt.Println("created book:", book)
	},
}

func init() {
	createCmd.AddCommand(bookCreateCmd)
	bookCreateCmd.Flags().String("book", "", "{title:harry potter,author:jk,description:fantasy,edition:1}")
	bookCreateCmd.Flags().String("collection", "", "fantasy")
	bookCreateCmd.MarkFlagRequired("book")
}
