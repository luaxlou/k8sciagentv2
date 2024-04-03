/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a new app version",
	Long: `Usage:
		k8sciagentv2 deploy -a appName -i imageName -v version -e env`,
	Run: func(cmd *cobra.Command, args []string) {

		version, _ := cmd.Flags().GetString("version")
		image, _ := cmd.Flags().GetString("image")
		app, _ := cmd.Flags().GetString("app")
		env, _ := cmd.Flags().GetString("env")

		Deploy(app, image, version, env)
	},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k8sciagentv2",
	Short: "k8s agent",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringP("version", "v", "", "new app version")
	deployCmd.Flags().StringP("image", "i", "", "docker image name")
	deployCmd.Flags().StringP("app", "a", "", "app name")
	deployCmd.Flags().StringP("env", "e", "", "env")
}
