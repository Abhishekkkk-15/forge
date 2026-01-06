package cmd

import (
	"fmt"
	"forge/internal"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <template>",
	Short: "Show information about a template",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		meta, err := internal.LoadMetadata(args[0])
		if err != nil {
			return err
		}

		fmt.Println("Name:       ", meta.Name)
		fmt.Println("Description:", meta.Description)
		fmt.Println("Language:   ", meta.Language)
		fmt.Println("Framework:  ", meta.Framework)
		if len(meta.Variables) > 0 {
			fmt.Println("\nVariables:")
			for name, v := range meta.Variables {
				fmt.Printf(
					"  %s (%s) default=%v\n    %s\n",
					name,
					v.Flag,
					v.Default,
					v.Description,
				)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
