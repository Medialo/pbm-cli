package monorepo

import (
	"fmt"
	"github.com/spf13/cobra"
)

var description string
var version string
var monorepoInitCmd = &cobra.Command{
	Use:   "init [name] (description)",
	Short: "Initialize a new monorepo project",
	Long:  `Initialize a new monorepo project with all configuration files`,
	Args: func(cmd *cobra.Command, args []string) error {
		var minSize = 5

		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}

		if len(args[0]) <= minSize {
			return fmt.Errorf("size name must be at least %d characters", minSize)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//project := prj.Project{
		//	Name:        args[0],
		//	Version:     version,
		//	Description: description,
		//}
		//prj.CreateProject(&project)
	},
}

//func init() {
//	fmt.Println("monorepo init init")
//	createCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the project")
//	createCmd.Flags().StringVarP(&version, "version", "v", "1", "Version of the project monorepo.tmpl")
//	// Here you will define your flags and configuration settings.
//
//	// Cobra supports Persistent Flags which will work for this command
//	// and all subcommands, e.g.:
//	// monorepoCmd.PersistentFlags().String("foo", "", "A help for foo")
//
//	// Cobra supports local flags which will only run when this command
//	// is called directly, e.g.:
//	// monorepoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//}
