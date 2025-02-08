/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/data_importers"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var initFlg bool
var LoadPath string
var dbHost string = "localhost"
var dbName string
var dbUser string
var dbPass string
var dbPort int16

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "The db connection and initializer",
	Long: `The DB connection sets the database connection paraameters and
also performs the initial database loading
of the state, zip, and area code databases. It will also import the
list of query keywords. If the -I flag is not specified, the database
loading will not take place, and instead the database connection will be
created for use with other parts of the application. When launching the
web service, include the db command and parameters for establishing the
database connection.`,
	Run: func(cmd *cobra.Command, args []string) {
		ghost, _ := cmd.Flags().GetString("host")
		if ghost != "" {
			dbHost = ghost
			fmt.Println("db host:", dbHost)
		}
		if dbPass == "" {
			fmt.Println("dbPass (-P) is required")
			return
		}
		sp := system.GetSystemParams()
		sp.Dbc = &sqldb.DBConnData{DBName: dbName,
			Host:     dbHost,
			User:     dbUser,
			Password: dbPass,
			Port:     dbPort,
		}
		// Import states
		err := sp.Dbc.Connect()
		if err != nil {
			fmt.Printf("Error connecting to DB: %v\n", err)
			return
		}

		if initFlg {
			states := &data_importers.States{DB: sp.Dbc.DB}
			_, err := LoadTables(states)
			if err != nil {

				//fmt.Printf("Error loading tables: %v\n", err)
			}
			//fmt.Printf("Loaded %d states into state table.\n", nload)

			zipcodes := &data_importers.ZCImport{DB: sp.Dbc.DB}
			_, err = LoadTables(zipcodes)
			if err != nil {
				fmt.Printf("Error loading zipcodes: %v\n", err)
			}
			//fmt.Printf("Loaded %d zipcodes into zipcode table.\n", nload)

			areacodes := &data_importers.ACImport{DB: sp.Dbc.DB}
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
	rootCmd.AddCommand(dbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	dbCmd.Flags().BoolVarP(&initFlg, "Init", "I", false, "Initialize database with imported data")
	dbCmd.Flags().StringVarP(&LoadPath, "loadpath", "L", "./data/", "Path to import files")
	dbCmd.Flags().StringVarP(&dbHost, "host", "H", dbHost, "Host")
	dbCmd.Flags().Int16VarP(&dbPort, "port", "p", 5432, "Port")
	dbCmd.Flags().StringVarP(&dbUser, "username", "U", "crawler", "Username")
	dbCmd.Flags().StringVarP(&dbPass, "password", "P", "", "Password")
	dbCmd.Flags().StringVarP(&dbName, "database", "d", "soleirclear", "Database name")

}
