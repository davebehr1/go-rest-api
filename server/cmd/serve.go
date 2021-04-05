package cmd

import (
	"fmt"
	"log"
	"lxdAssessmentServer/pkg/routes"
	"net"
	"net/http"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Spins up http server",
	Run: func(cmd *cobra.Command, args []string) {

		router := routes.RouteHandlers()
		lhttp, err := net.Listen("tcp", ":8080")

		if err != nil {
			log.Fatal("listen error:", err)
		}

		fmt.Println("listening on port 8080")

		http.Serve(lhttp, router)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
