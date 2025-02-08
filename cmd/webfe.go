/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// webfeCmd represents the webfe command
var webhost string = "localhost"
var webport int16 = 9090
var webfeCmd = &cobra.Command{
	Use:   "webfe",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ghost, _ := cmd.Flags().GetString("host")
		if ghost == "" {
			webhost = ghost
		}
		fmt.Println("webfe called")
	},
}

func init() {
	rootCmd.AddCommand(webfeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webfeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webfeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	webfeCmd.Flags().StringVarP(&webhost, "host", "H", webhost, "Web server Host")
	webfeCmd.Flags().Int16VarP(&webport, "port", "p", webport, "Web server port")

}
