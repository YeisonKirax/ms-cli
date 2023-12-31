// Package create /*
package cmd

import (
	"ms-cli/pkg/commands/create"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new microservice",
	Long: `Create a new microservice using the languages and templates predefined. Additionally, 
generate a base structure.`,
	Run: func(cmd *cobra.Command, args []string) {
		create.CreateMicroservice()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
