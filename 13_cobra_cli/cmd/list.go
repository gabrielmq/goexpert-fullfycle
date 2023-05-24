/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gabrielmq/cobra_cli/internal/database"
	"github.com/spf13/cobra"
)

func newListCmd(categoryDB database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List categories",
		Long:  `List categories`,
		RunE:  runList(categoryDB),
	}
}

func runList(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		categories, err := categoryDB.FindAll()
		if err != nil {
			return err
		}
		for _, category := range categories {
			cmd.Println(category)
		}
		return nil
	}
}

func init() {
	listCmd := newListCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(listCmd)
}
