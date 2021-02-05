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
	"context"
	"fmt"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go"
	"github.com/spf13/cobra"
)

var Quote string

// quoteCmd represents the quote command
var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "Ask for quote",
	Long:  `Quote for a specific stock`,
	Run: func(cmd *cobra.Command, args []string) {
		finnhubClient := finnhub.NewAPIClient(finnhub.NewConfiguration()).DefaultApi
		auth := context.WithValue(context.Background(), finnhub.ContextAPIKey, finnhub.APIKey{
			Key: Token,
		})

		quote, _, err := finnhubClient.Quote(auth, Quote)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Current price: %+v\n", quote.C)
		fmt.Printf("Highest price of the day: %+v\n", quote.H)
		fmt.Printf("Lowest price of the day: %+v\n", quote.L)
		fmt.Printf("Open price: %+v\n", quote.O)
		fmt.Printf("Previous Close price: %+v\n", quote.Pc)
	},
}

func init() {
	rootCmd.AddCommand(quoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	quoteCmd.Flags().StringVarP(&Quote, "quote", "q", "", "The stock to look for")
	quoteCmd.MarkFlagRequired("quote")
}
