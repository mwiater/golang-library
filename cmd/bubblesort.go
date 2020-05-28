/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"math/rand"
  "time"

	"github.com/spf13/cobra"
)

// bubblesortCmd represents the bubblesort command
var bubblesortCmd = &cobra.Command{
	Use:   "bubblesort",
	Short: "A bubblesort example.",
	Long: `A bubblesort example`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running: bubblesort")

		slice := generateSlice(20)
    fmt.Println("\n--- Unsorted --- \n\n", slice)
    bubblesort(slice)
    fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
	},
}

func init() {
	rootCmd.AddCommand(bubblesortCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bubblesortCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bubblesortCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
 
    slice := make([]int, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - rand.Intn(999)
    }
    return slice
}
  
func bubblesort(items []int) {
    var (
        n = len(items)
        sorted = false
    )
    for !sorted {
        swapped := false
        for i := 0; i < n-1; i++ {
            if items[i] > items[i+1] {
                items[i+1], items[i] = items[i], items[i+1]
                swapped = true
            }
        }
        if !swapped {
            sorted = true
        }
        n = n - 1
    }
}
