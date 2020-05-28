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
	"math/big"

	"github.com/spf13/cobra"
)

// piCmd represents the pi command
var piCmd = &cobra.Command{
	Use:   "pi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running: pi")

		numberofDigits := int64(10000)
		result := Pi(numberofDigits)

		fmt.Println("Calculated", numberofDigits, "digits of Pi:", result)
	},
}

func init() {
	rootCmd.AddCommand(piCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// piCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// piCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func arccot(x int64, unity *big.Int) *big.Int {
	bigx := big.NewInt(x)
	xsquared := big.NewInt(x * x)
	sum := big.NewInt(0)
	sum.Div(unity, bigx)
	xpower := big.NewInt(0)
	xpower.Set(sum)
	n := int64(3)
	zero := big.NewInt(0)
	sign := false

	term := big.NewInt(0)
	for {
		xpower.Div(xpower, xsquared)
		term.Div(xpower, big.NewInt(n))
		if term.Cmp(zero) == 0 {
			break
		}
		if sign {
			sum.Add(sum, term)
		} else {
			sum.Sub(sum, term)
		}
		sign = !sign
		n += 2
	}
	return sum
}

func Pi( ndigits int64 ) string {
	if ( ndigits <=7  ) {
		return "3.141595"
	} else {
		digits := big.NewInt(ndigits + 10)
		unity := big.NewInt(0) // crea un entero tocho
		unity.Exp(big.NewInt(10), digits, nil) // Le asigna valor
		pi := big.NewInt(0)
		four := big.NewInt(4) // Todos deben ser enteros tocho
		
		// Serie de McLaurin
		pi.Mul(four, pi.Sub(pi.Mul(four, arccot(5, unity)), arccot(239, unity)))
		output := fmt.Sprintf("%s.%s",pi.String()[0:1],pi.String()[1:ndigits])
		return output
	}
}