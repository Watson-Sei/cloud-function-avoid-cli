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
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
	"log"
	"os"
)

//newShowCmd returns cobra.Command instance for show sub-command
func newClaimsCmd() *cobra.Command {
	claimsCmd := &cobra.Command{
		Use:   "claims",
		Short: "Short comment for show sub-command",
		Long:  "Long comment for show sub-command",
		RunE: func(cmd *cobra.Command, args []string) error {
			set, err := cmd.Flags().GetBool("customClaimsSet")
			if err != nil {
				return err
			}
			if set {
				uid, err := cmd.Flags().GetString("uid")
				if err != nil {
					return err
				}

				ctx := context.Background()

				opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
				app, err := firebase.NewApp(context.Background(), nil, opt)
				if err != nil {
					log.Printf("error initializing app: %v", err)
					return err
				}

				client, err := app.Auth(ctx)
				if err != nil {
					log.Fatalf("error getting Auth client: %v\n", err)
					return err
				}

				if err = customClaimsSet(ctx, client, uid); err != nil {
					return err
				}

				cui.Outputln("Custom Claims Set Up")
			}
			return err
		},
	}
	// Note: https://github.com/spf13/pflag/blob/master/flag.go#L362
	claimsCmd.Flags().BoolP("customClaimsSet", "s", false, "Identify Custom Claims command processing.")
	claimsCmd.Flags().StringP("uid", "u", "", "The uid of the user you want to make an Admin.")

	return claimsCmd
}

func customClaimsSet(ctx context.Context, client *auth.Client, uid string) error {
	claims := map[string]interface{}{"admin": true}
	err := client.SetCustomUserClaims(ctx, uid, claims)
	if err != nil {
		log.Fatalf("error setting custom claims %v\n", err)
		return err
	}
	return nil
}