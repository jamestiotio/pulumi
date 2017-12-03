// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/pulumi/pulumi/pkg/backend/cloud"
	"github.com/pulumi/pulumi/pkg/util/cmdutil"
)

func newLoginCmd() *cobra.Command {
	var cloudURL string
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Log into the Pulumi Cloud",
		Long:  "Log into the Pulumi Cloud.  You can script by using PULUMI_ACCESS_TOKEN environment variable.",
		Args:  cmdutil.NoArgs,
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			if cloudURL == "" {
				// If no URL was specified, assume it's the default URL.
				cloudURL = cloud.DefaultURL()
			}
			return cloud.Login(cloudURL)
		}),
	}
	cmd.PersistentFlags().StringVarP(&cloudURL, "cloud-url", "c", "", "A cloud URL to log into")
	return cmd
}

func newLogoutCmd() *cobra.Command {
	var cloudURL string
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Log out of the Pulumi Cloud",
		Long:  "Log out of the Pulumi Cloud.  Deletes stored credentials on the local machine.",
		Args:  cmdutil.NoArgs,
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			if cloudURL == "" {
				// If no URL was specified, assume it's the default URL.
				cloudURL = cloud.DefaultURL()
			}
			return cloud.Logout(cloudURL)
		}),
	}
	cmd.PersistentFlags().StringVarP(&cloudURL, "cloud-url", "c", "", "A cloud URL to log out of")
	return cmd
}
