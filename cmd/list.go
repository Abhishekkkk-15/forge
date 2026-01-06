package cmd

import (
	internal "forge/internal"
	"sort"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := internal.TemplateFS.ReadDir("templates")
		if err != nil {
			pterm.Error.Println("Templates not found")
			return err
		}

		pterm.DefaultSection.Println("Available Templates")

		tableData := pterm.TableData{
			{"Name", "Description"},
		}

		for _, e := range entries {
			meta, err := internal.LoadMetadata(e.Name())
			if err != nil {
				continue
			}

			tableData = append(tableData, []string{
				meta.Name,
				meta.Description,
			})
		}
		sort.Slice(tableData[1:], func(i, j int) bool {
			return tableData[i+1][0] < tableData[j+1][0]
		})

		pterm.DefaultTable.
			WithHasHeader().
			WithData(tableData).
			Render()
		pterm.Info.Printf("Found %d templates\n", len(tableData)-1)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
