package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lxdAssessmentClient/pkg"
	"net/http"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

type BookResponse struct {
	ID            int
	Title         string
	Description   string
	Author        string
	Collection    string
	Edition       int
	PublishedDate string
}

// booksCmd represents the books command
var booksCmd = &cobra.Command{
	Use:   "books",
	Short: "returns a list of books from the books collection api",

	Run: func(cmd *cobra.Command, args []string) {
		author, _ := cmd.Flags().GetString("author")
		title, _ := cmd.Flags().GetString("title")
		fromDate, _ := cmd.Flags().GetString("from_date")
		toDate, _ := cmd.Flags().GetString("to_date")

		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://localhost:8080/1.0/books", nil)
		req.Header.Add("Accept", "application/json")

		q := req.URL.Query()
		if author != "" {
			q.Add("author", author)
		}
		if title != "" {
			q.Add("title", title)
		}
		if fromDate != "" {
			q.Add("fromDate", fromDate)
		}
		if toDate != "" {
			q.Add("toDate", toDate)
		}

		req.URL.RawQuery = q.Encode()

		resp, err := client.Do(req)
		if err != nil {
			fmt.Print(err.Error())
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err.Error())
		}

		var books []BookResponse
		var response pkg.HttpResponse
		json.Unmarshal(bodyBytes, &response)
		payload, _ := json.Marshal(response.Payload)
		json.Unmarshal(payload, &books)
		fmt.Println(books)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"id", "title", "description", "author", "edition", "publishedDate", "collection"})
		for _, book := range books {
			t.AppendRows([]table.Row{
				{book.ID, book.Title, book.Description, book.Author, book.Edition, book.PublishedDate, book.Collection},
			})
		}
		t.Render()
	},
}

func init() {
	listCmd.AddCommand(booksCmd)
	booksCmd.Flags().String("author", "", "author of book")
	booksCmd.Flags().String("title", "", "title of book")
	booksCmd.Flags().String("from_date", "", "published from this date")
	booksCmd.Flags().String("to_date", "", "published to this date")
}
