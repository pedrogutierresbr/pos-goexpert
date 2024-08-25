/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/pedrogutierresbr/pos-goexpert/16-cobra_cli/internal/database"

	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  "Create a new category",
		RunE:  runCreate(categoryDb),
	}
}

// // createCmd represents the create command
// var createCmd = &cobra.Command{
// 	Use:   "create",
// 	Short: "A brief description of your command",
// 	Long:  `A longer description that spans multiple lines and likely contains examples`,

// 	RunE: runCreate(GetCategoryDB(GetDb())), // ta chamando tudo por injeção de dependência, fica mais fácil de testar
// }

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.Flags().StringP("description", "d", "", "Description of the category")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}

// criando tables
// entrar no sqlite: sqlite3 data.db

// Categories: create table categories (id string, name string, description string);
// limpar database: delete from categories;

// Comando para criar uma categoria via CLI
// go run main.go category create -n={Nome categotia} -d={nome da descrição}
