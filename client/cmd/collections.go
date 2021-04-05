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

type CollectionResponse struct {
	ID         int
	Name       string
	BookAmount int
}

// collectionsCmd represents the collections command
var collectionsCmd = &cobra.Command{
	Use:   "collections",
	Short: "returns a list of collections from the books collection api",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://localhost:8080/1.0/collections", nil)
		req.Header.Add("Accept", "application/json")

		q := req.URL.Query()
		if name != "" {
			q.Add("name", name)
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

		var collections []CollectionResponse
		var response pkg.HttpResponse
		json.Unmarshal(bodyBytes, &response)
		payload, _ := json.Marshal(response.Payload)
		json.Unmarshal(payload, &collections)

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"id", "name", "amount of books"})
		for _, collection := range collections {
			t.AppendRows([]table.Row{
				{collection.ID, collection.Name, collection.BookAmount},
			})
		}
		t.Render()
	},
}

func init() {
	listCmd.AddCommand(collectionsCmd)
	collectionsCmd.Flags().String("name", "", "name of collection")
}
