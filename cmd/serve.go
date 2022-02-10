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
	"github.com/marlinprotocol/spanchk/serve"
	"github.com/spf13/cobra"
)

var (
	listenAddr    string
	heimdallAddr  string
	validatorAddr string
)

// startCmd represents the generate command
var startCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve queries",
	Long:  `Serve queries`,
	Run: func(cmd *cobra.Command, args []string) {
		serve.Serve(listenAddr, heimdallAddr, validatorAddr)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVarP(&heimdallAddr, "hmdaddr", "a", "127.0.0.1:1317", "Heimdall endpoint to query for span")
	startCmd.Flags().StringVarP(&listenAddr, "listenaddr", "l", "127.0.0.1:13180", "Endpoint to serve queries")
	startCmd.Flags().StringVarP(&validatorAddr, "validatoraddr", "v", "0x0", "Validator to inform about")
}
