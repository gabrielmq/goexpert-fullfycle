/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(category)
	},
	// Hook que vai executar algo antes do comando
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	cmd.Println("executado antes do Run")
	// },
	// Hook qque vai executar algo depois do comando
	// PostRun: func(cmd *cobra.Command, args []string) {
	// 	cmd.Println("executado depois do Run")
	// },
	// Esse RunE retorna um erro
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	return fmt.Errorf("Executou e retorna uma possivel erro")
	// },
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)

	// associa o valor da flag pra variavel category por referencia
	categoryCmd.Flags().StringVarP(&category, "name", "n", "", "")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
