package cmd

import (
	"fmt"
	internal "forge/internal"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		entreis, err := internal.TemplateFS.ReadDir("templates")
		if err != nil {
			fmt.Print("Tempates not found")
			return err
		}
		fmt.Println("Available tempates")
		for _, e := range entreis {
			meta, _ := internal.LoadMetadata(e.Name())
			fmt.Printf("- %s: %s\n", meta.Name, meta.Description)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
