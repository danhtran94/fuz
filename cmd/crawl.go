// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/danhtran94/fuz/pkg/dbcare"
	"github.com/danhtran94/fuz/pkg/template"

	"github.com/spf13/cobra"
)

var dbConnectString string

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Generate patterns for dest table",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		tableName := args[0]

		dbcare.SetupDBClient(dbConnectString, true)
		cols, err := dbcare.GetColTypes(tableName)
		if err != nil {
			log.Println(err)
			return
		}

		template := template.NewTemplate(tableName, cols)
		jsonDefine, err := json.MarshalIndent(template, "", "  ")
		if err != nil {
			log.Println(err)
			return
		}

		err = ioutil.WriteFile(fmt.Sprintf("./templates/%s.json", tableName), jsonDefine, 0644)
		if err != nil {
			log.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crawlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	crawlCmd.Flags().StringVar(&dbConnectString, "db", "", "Postgres connect string")
}
