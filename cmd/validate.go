/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/quietstorm/qs/stuff"

	"github.com/spf13/cobra"
)

// Variables to get values from flags
// Local ones only work in the working command
var Pub string
var Priv string

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		stuff.ReadCert(Pub, Priv)
	},
}

func init() {
	tokenCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// works like qs token validate -p /path/to -v ~/home/stuff
	validateCmd.Flags().StringVarP(&Pub, "pub", "p", "", "Path to public key")
	validateCmd.Flags().StringVarP(&Priv, "priv", "v", "", "Path to private key")
}
