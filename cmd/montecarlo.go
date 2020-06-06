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
  "os"
  "os/exec"
  "runtime"
  "strconv"

  "gopkg.in/gookit/color.v1"
  "github.com/spf13/cobra"
)

const samples = 10000000

// montecarloCmd represents the montecarlo command
var montecarloCmd = &cobra.Command{
  Use:   "montecarlo",
  Short: "A brief description of your command",
  Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  Run: func(cmd *cobra.Command, args []string) {
    clear := exec.Command("clear")
    clear.Stdout = os.Stdout
    clear.Run()

    greenFG := color.FgGreen.Render
    greenBG := color.New(color.FgBlack, color.BgGreen).Render

    fmt.Println("[COBRA] montecarlo called...\n")

    fmt.Printf("Calculating %v digits of Pi.\n\n", greenBG(" "+strconv.FormatInt(samples, 10)+" "))

    var cpuCount = runtime.NumCPU()
    fmt.Println("CPU Count:\t\t", greenFG(cpuCount), "\n")

    runtime.GOMAXPROCS(runtime.NumCPU())
    rand.Seed(time.Now().UnixNano())

    var piStart               = time.Now().UnixNano()
    var piResult              = PI(samples)
    var piEnd                 = time.Now().UnixNano()
    
    var multipiStart          = time.Now().UnixNano()
    var multipiResult         = MultiPI(samples)
    var multipiEnd            = time.Now().UnixNano()
    
    var piProcessingTime      = (piEnd-piStart)
    var multipiProcessingTime = (multipiEnd-multipiStart)
    var benchmarkTime         = piProcessingTime/multipiProcessingTime

    fmt.Println("PI Result:\t\t", greenFG(piResult))
    fmt.Println("MultiPI Result:\t\t", greenFG(multipiResult), "\n")

    if benchmarkTime <= 1 {
      benchmarkTime = multipiProcessingTime/piProcessingTime
      fmt.Printf("Pi is %s faster than MultiPI.\n\n", greenBG(" "+strconv.FormatInt(benchmarkTime, 10)+"x "))
    } else {
      fmt.Printf("MultiPI is %s faster than PI.\n\n", greenBG(" "+strconv.FormatInt(benchmarkTime, 10)+"x "))
    }
  },
}

func init() {
  rootCmd.AddCommand(montecarloCmd)

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // montecarloCmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // montecarloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func PI(samples int) float64 {
  var inside int = 0

  for i := 0; i < samples; i++ {
    x := rand.Float64()
    y := rand.Float64()
    if (x*x + y*y) < 1 {
      inside++
    }
  }

  ratio := float64(inside) / float64(samples)

  return ratio * 4
}

func MultiPI(samples int) float64 {
  cpus := runtime.NumCPU()

  threadSamples := samples / cpus
  results := make(chan float64, cpus)

  for j := 0; j < cpus; j++ {
    go func() {
      var inside int
      r := rand.New(rand.NewSource(time.Now().UnixNano()))
      for i := 0; i < threadSamples; i++ {
        x, y := r.Float64(), r.Float64()

        if x*x+y*y <= 1 {
          inside++
        }
      }
      results <- float64(inside) / float64(threadSamples) * 4
    }()
  }

  var total float64
  for i := 0; i < cpus; i++ {
    total += <-results
  }

  return total / float64(cpus)
}
