/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/http"
	"github.com/spf13/cobra"
)

// webfeCmd represents the webfe command
var webhost string = "localhost"
var webport int16 = 9090
var webfeCmd = &cobra.Command{
	Use:   "webfe",
	Short: "webfe launches the web front end",
	Long: `webfe launches the web front end
host and port are required for the front end to launch.`,
	Run: func(cmd *cobra.Command, args []string) {
		ghost, _ := cmd.Flags().GetString("host")
		if ghost == "" {
			webhost = ghost
		}
		fmt.Printf("webfe: web server listening on %s:%d\n", ghost, webport)
		runWS(webhost, webport)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webfeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webfeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	webfeCmd.Flags().StringVarP(&webhost, "webhost", "W", webhost, "Web server Host")
	webfeCmd.Flags().Int16VarP(&webport, "webport", "w", webport, "Web server port")
	fmt.Printf("webfe::init finished.\n")
}

func runWS(host string, port int16) error {
	return http.ServerStart(host, port)
}
