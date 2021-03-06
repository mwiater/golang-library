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
  "encoding/json"

  "github.com/elastic/go-sysinfo"
  "github.com/spf13/cobra"
)

// sysinfoCmd represents the sysinfo command
var sysinfoCmd = &cobra.Command{
  Use:   "sysinfo",
  Short: "Get infoprmation about your host.",
  Long: `Get infoprmation about your host.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Running: sysinfo")

    // See: https://github.com/elastic/go-sysinfo/blob/master/system_test.go
    host, err := sysinfo.Host()
    if err != nil {
      fmt.Println("ERROR:", err)
    }

    info := host.Info()

    memory, err := host.Memory()
    if err != nil {
      fmt.Println("ERROR:", err)
    }

    cpu, err := host.CPUTime()
    if err != nil {
      fmt.Println("ERROR:", err)
    }
    
    infoJSON, _   := json.MarshalIndent(info, "", "  ")
    memoryJSON, _ := json.MarshalIndent(cpu, "", "  ")
    cpuJSON, _    := json.MarshalIndent(memory, "", "  ")

    fmt.Println("Info:", string(infoJSON), "\n")
    fmt.Println("Memory Info:", string(memoryJSON), "\n")
    fmt.Println("CPU Info:", string(cpuJSON), "\n")
  },
}

func init() {
  rootCmd.AddCommand(sysinfoCmd)

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // sysinfoCmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // sysinfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}