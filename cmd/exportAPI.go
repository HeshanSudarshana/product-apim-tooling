// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/menuka94/wso2apim-cli/utils"
)

var exportAPIName string
var exportAPIVersion string
var exportEnvironment string

// ExportAPICmd represents the exportAPI command
var ExportAPICmd = &cobra.Command{
	Use:   "exportAPI (--name <name-of-the-api> --version <version-of-the-api> --environment <environment-to-which-the-api-should-be-exported>)",
	Short: utils.ExportAPICmdLongDesc,
	Long:  utils.ExportAPICmdLongDesc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exportAPI called")

		fmt.Println("Name:", exportAPIName)
		fmt.Println("Version:", exportAPIVersion)
		fmt.Println("Environment:", exportEnvironment)

		m := utils.GetOAuthTokens("admin", "admin")
		accessToken := m["access_token"]
		refreshToken := m["refresh_token"]
		tokenType := m["token_type"]
		expiresIn := m["expires_in"]

		fmt.Println("AccessToken:", accessToken)
		fmt.Println("RefreshToken:", refreshToken)
		fmt.Println("TokenType:", tokenType)
		fmt.Println("ExpiresIn:", expiresIn)
	},
}

func init() {
	RootCmd.AddCommand(ExportAPICmd)
	ExportAPICmd.Flags().StringVarP(&exportAPIName, "name", "n", "", "Name of the API to be exported")
	ExportAPICmd.Flags().StringVarP(&exportAPIVersion, "version", "s", "", "Version of the API to be exported")
	ExportAPICmd.Flags().StringVarP(&exportEnvironment, "environment", "e", "", "Environment to which the API should be exported")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ExportAPICmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ExportAPICmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
