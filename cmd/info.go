package cmd

import (
	"fmt"
	"forge/internal"

	"github.com/pterm/pterm"
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

		pterm.DefaultSection.Println("Template Info")

		pterm.DefaultTable.
			WithData(pterm.TableData{
				{"Name", meta.Name},
				{"Description", meta.Description},
				{"Language", meta.Language},
				{"Framework", meta.Framework},
			}).
			Render()

		if len(meta.Variables) > 0 {
			pterm.DefaultSection.Println("Variables")

			data := pterm.TableData{
				{"Name", "Flag", "Default", "Description"},
			}

			for name, v := range meta.Variables {
				data = append(data, []string{
					name,
					v.Flag,
					fmt.Sprint(v.Default),
					v.Description,
				})
			}

			pterm.DefaultTable.
				WithHasHeader().
				WithData(data).
				Render()
		}

		if len(meta.Flags) > 0 {
			pterm.DefaultSection.Println("Flags")

			data := pterm.TableData{
				{"Flag", "Type", "Default", "Description"},
			}

			for name, v := range meta.Flags {
				data = append(data, []string{
					name,
					v.Type,
					fmt.Sprint(v.Default),
					v.Description,
				})
			}

			pterm.DefaultTable.
				WithHasHeader().
				WithData(data).
				Render()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
