/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"

	"github.com/spf13/cobra"
)

// dbinitCmd represents the dbinit command
var initFlg bool
var LoadPath string
var dbHost string
var dbName string
var dbUser string
var dbPass string
var dbPort int8

var dbinitCmd = &cobra.Command{
	Use:   "dbinit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dbinit called")
		if initFlg {
			fmt.Printf("Init called!\n")
		}
		if dbPass == "" {
			fmt.Println("dbPass (-P) is required")
			return
		}
		dbcdata := &sqldb.DBConnData{DBName: dbName,
			Host:     dbHost,
			User:     dbUser,
			Password: dbPass,
			Port:     dbPort,
		}
		dbcdata.Connect()
		defer dbcdata.Close()
		// Import states

	},
}

func init() {
	rootCmd.AddCommand(dbinitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbinitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbinitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	dbinitCmd.Flags().BoolVarP(&initFlg, "Init", "I", false, "Initialize database with imported data")
	dbinitCmd.Flags().StringVarP(&LoadPath, "loadpath", "L", "../data/", "Path to import files")
	dbinitCmd.Flags().StringVarP(&dbHost, "host", "H", "localhost", "Host")
	dbinitCmd.Flags().Int8VarP(&dbPort, "port", "p", 5432, "Port")
	dbinitCmd.Flags().StringVarP(&dbUser, "username", "U", "crawler", "Username")
	dbinitCmd.Flags().StringVarP(&dbPass, "password", "P", "", "Password")
	dbinitCmd.Flags().StringVarP(&dbName, "database", "d", "soleirclear", "Database name")

}
