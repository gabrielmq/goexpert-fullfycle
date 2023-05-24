/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gabrielmq/cobra_cli/internal/database"
	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDB database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  `Create a new category`,
		RunE:  runCreate(categoryDB),
	}
}

func runCreate(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDB.Create(name, description)
		return err
	}
}

func init() {
	db := GetDb()
	categoryDB := GetCategoryDB(db)
	createCmd := newCreateCmd(categoryDB)

	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "")
	createCmd.Flags().StringP("description", "d", "", "")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
