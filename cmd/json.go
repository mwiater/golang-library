/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
  "fmt"
  "io/ioutil"
	"os"
	"log"
	"path"
	"math/rand"
	"time"
	//"strings"
	"strconv"

	"github.com/spf13/cobra"
)

var geneMap map[string]int

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentPath, err := os.Getwd()
		if err != nil {
				log.Println(err)
		}

		data, err := ioutil.ReadFile(path.Join(currentPath, "/cmd/data/onezero.json"))
    if err != nil {
			fmt.Println(err)
			return
    }

		var onesAndZeros []int
    err = json.Unmarshal(data, &onesAndZeros)
    if err != nil {
        fmt.Println(err)
        return
    }
 
		min := 0
    max := len(onesAndZeros)-1

		var colony [][]int
		var maxColonySize = 10000

		geneMap = make(map[string]int)

		for len(colony) < maxColonySize {
			rand.Seed(time.Now().UnixNano())
			var gene []int
			for len(gene) < 8 {
				randomIndex := rand.Intn(max - min) + min
				cell := onesAndZeros[(randomIndex-1):randomIndex][0]
				gene = append(gene, cell)
			}

			colony = append(colony, gene)
		}

		for _, gene := range colony {
			geneString := ""
			for _, cell := range gene {
				geneString = geneString+strconv.Itoa(cell)
			}
			geneMap[geneString] = geneMap[geneString]+1
	}

	totalTally := 0
	for _, element := range geneMap {
		totalTally = totalTally + element
	}

	// VALIDATE
	if totalTally != len(colony) {
		fmt.Println("VALIDATION ERROR: totalTally DOES NOT EQUAL len(colony)")
		os.Exit(1)
	}


	fmt.Println("geneMap",geneMap)
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
