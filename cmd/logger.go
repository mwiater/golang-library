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
  "reflect"
  //"encoding/json"

  "github.com/elastic/go-sysinfo"
  log "github.com/sirupsen/logrus"
  "github.com/spf13/cobra"
)

// loggerCmd represents the logger command
var loggerCmd = &cobra.Command{
  Use:   "logger",
  Short: "A brief description of your command",
  Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("logger called")

    fmt.Println("\n----------\n")

    log.WithFields(log.Fields{
      "exampleField01": "walrus",
      "exampleField02": 10,
    }).Info("A log message with fields >>>")

    log.WithFields(log.Fields{
      "exampleField01": true,
      "exampleField02": 122,
    }).Warn("Aother log message with different fields >>>")

    // A common pattern is to re-use fields between logging statements by re-using
    // the logrus.Entry returned from WithFields()
    contextLogger := log.WithFields(log.Fields{
      "reusedField01": "Always logged when called this way...",
      "reusedField02": "Also always logged when called this way...",
    })

    contextLogger.Info("I am a log message with reused fields >>>")
    contextLogger.Info("Reuse example >>>")
    
    // Syslog example
    host, err := sysinfo.Host()
    if err != nil { fmt.Println("ERROR:", err) }

    info        := host.Info()
    //infoJSON, _ := json.MarshalIndent(info, "", "  ")

    // Kinda buried
    // HostInfo Struct: https://github.com/elastic/go-sysinfo/blob/47d31290d0f47c2a72d998ec09e8425a4abacbc7/types/host.go#L30
    // OS: https://github.com/elastic/go-sysinfo/blob/47d31290d0f47c2a72d998ec09e8425a4abacbc7/types/host.go#L50
    sysInfoLogger := log.WithFields(log.Fields{
      "Hostname": string(info.Hostname),
      "OS": string(info.OS.Name)+" "+string(info.OS.Version),
      "Timezone": string(info.Timezone),
    })

    sysInfoLogger.Info("Log message with SysInfo fields >>>")

    fmt.Println("\n----------\n")

    infoStruct       := reflect.ValueOf(info)
    typeOfInfoStruct := infoStruct.Type()
  
    for i := 0; i< infoStruct.NumField(); i++ {
      fmt.Printf("Field: %s\n", typeOfInfoStruct.Field(i).Name)
      fmt.Printf("Value: %s\n", infoStruct.Field(i).Interface())
      fmt.Printf("Type: %s\n", reflect.TypeOf(infoStruct.Field(i).Interface()))
      fmt.Println("----------\n")
    }

    // Fatal: Nothing executes beyond this when it is thrown
    //log.WithFields(log.Fields{
    //  "omg":    true,
    //  "number": 100,
    //}).Fatal("The ice breaks!")
  },
}

func init() {
  rootCmd.AddCommand(loggerCmd)

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // loggerCmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // loggerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
