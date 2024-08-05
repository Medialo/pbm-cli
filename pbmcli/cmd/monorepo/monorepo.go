/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package monorepo

import (
	"fmt"
	"github.com/spf13/cobra"
)

// monorepoCmd represents the monorepo.tmpl command
var monorepoCmd = &cobra.Command{
	Use:   "monorepo",
	Short: "All monorepo commands",
	Long:  `Monorepo long`,
}

func NewMonorepoCmd() *cobra.Command {
	return monorepoCmd
}

func init() {
	fmt.Println("monorepo init")
	monorepoCmd.AddCommand(infoCmd)
	monorepoCmd.AddCommand(monorepoInitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monorepoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monorepoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
