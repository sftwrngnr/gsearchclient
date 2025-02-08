/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/data_importers"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"github.com/spf13/cobra"
)

// dbinitCmd represents the dbinit command
var initFlg bool
var LoadPath string
var dbHost string = "localhost"
var dbName string
var dbUser string
var dbPass string
var dbPort int16

var dbinitCmd = &cobra.Command{
	Use:   "dbinit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ghost, _ := cmd.Flags().GetString("host")
		if ghost != "" {
			dbHost = ghost
			fmt.Println("dbinit host:", dbHost)
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
		// Import states
		err := dbcdata.Connect()
		if err != nil {
			fmt.Printf("Error connecting to DB: %v\n", err)
			return
		}

		if initFlg {
			states := &data_importers.States{DB: dbcdata.DB}
			_, err := LoadTables(states)
			if err != nil {

				//fmt.Printf("Error loading tables: %v\n", err)
			}
			//fmt.Printf("Loaded %d states into state table.\n", nload)

			zipcodes := &data_importers.ZCImport{DB: dbcdata.DB}
			_, err = LoadTables(zipcodes)
			if err != nil {
				fmt.Printf("Error loading zipcodes: %v\n", err)
			}
			//fmt.Printf("Loaded %d zipcodes into zipcode table.\n", nload)

			areacodes := &data_importers.ACImport{DB: dbcdata.DB}
			_, err = LoadTables(areacodes)
			if err != nil {
				//fmt.Printf("Error loading tables: %v\n", err)
				return
			}

		}

	},
}

func LoadTables(myClass data_importers.Importer) (int, error) {
	if myClass.Init(LoadPath) {
		return myClass.Import()
	} else {
		return 0, errors.New("Error with intialization.")
	}
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
	dbinitCmd.Flags().StringVarP(&LoadPath, "loadpath", "L", "./data/", "Path to import files")
	dbinitCmd.Flags().StringVarP(&dbHost, "host", "H", dbHost, "Host")
	dbinitCmd.Flags().Int16VarP(&dbPort, "port", "p", 5432, "Port")
	dbinitCmd.Flags().StringVarP(&dbUser, "username", "U", "crawler", "Username")
	dbinitCmd.Flags().StringVarP(&dbPass, "password", "P", "", "Password")
	dbinitCmd.Flags().StringVarP(&dbName, "database", "d", "soleirclear", "Database name")

}
