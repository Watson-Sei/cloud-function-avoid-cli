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
	"github.com/spf13/cobra"
)

//newShowCmd returns cobra.Command instance for show sub-command
func newClaimsCmd() *cobra.Command {
	claimsCmd := &cobra.Command{
		Use:   "claims",
		Short: "Short comment for show sub-command",
		Long:  "Long comment for show sub-command",
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := cmd.Flags().GetInt("integer")
			if err != nil {
				return err
			}
			b, err := cmd.Flags().GetBool("boolean")
			if err != nil {
				return err
			}
			s, err := cmd.Flags().GetString("string")
			if err != nil {
				return err
			}
			cui.Outputln("Integer option value:", i)
			cui.Outputln(" String option value:", s)
			cui.Outputln("Boolean option value:", b)

			return nil
		},
	}

	claimsCmd.Flags().IntP("integer", "i", 0, "integer option")
	claimsCmd.Flags().BoolP("boolean", "b", false, "boolean option")
	claimsCmd.Flags().StringP("string", "s", "", "string option")

	return claimsCmd
}
