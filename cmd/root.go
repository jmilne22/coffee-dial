/*
Copyright © 2025 NAME HERE james@jamesmilne.io
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "coffee-dial",
	Short: "A CLI tool for recording and managing coffee dial-in recipes",
	Long: `coffee-dial is a command-line application for recording, managing, and sharing coffee dial-in recipes.

With coffee-dial, you can log details about your coffee beans, roasters, brewing methods, grinder settings, and more.
This helps you track your dial-in process and reproduce your favorite brews.

Example usage:
		coffee-dial --name "Ethiopia Yirgacheffe" --roaster "Blue Bottle" --method "V60" --grinder "Baratza Encore" --setting "15"

Flags:
		--name      The name of the coffee.
		--roaster   The name of the roaster.
		--method    The method of brewing the coffee.
		--grinder   The name of the grinder.
		--setting   The grinder setting.

For more information, visit the project repository or run 'coffee-dial --help'.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.coffee-dial.yaml)")
	rootCmd.PersistentFlags().StringP("name", "n", "", "The name of the coffee.")
	rootCmd.PersistentFlags().StringP("roaster", "r", "", "The name of the roaster.")
	rootCmd.PersistentFlags().StringP("method", "m", "", "The method of brewing the coffee.")
	rootCmd.PersistentFlags().StringP("grinder", "g", "", "The name of the grinder.")
	rootCmd.PersistentFlags().StringP("setting", "s", "", "The grinder setting.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
