package cmd

import (
	"fmt"
	"forge/internal"

	"github.com/spf13/cobra"
)

func registerFlags(cmd *cobra.Command, meta *internal.TemplateMetadata) {
	for name, variable := range meta.Variables {
		if variable.Source != "flag" {
			continue
		}
		value := variable.Default
		dynamicFlages[name] = &value

		switch variable.Type {
		case "string":
			cmd.Flags().StringVar(
				dynamicFlages[name].(*string),
				variable.Flag,
				fmt.Sprint(variable.Default),
				variable.Description,
			)
		case "number":
			cmd.Flags().Int(
				variable.Flag,
				int(variable.Default.(float64)),
				variable.Description,
			)

		}
	}
	for name, flag := range meta.Flags {
		value := flag.Default
		dynamicFlages[name] = &value
		cmd.Flags().Bool(
			name,
			flag.Default.(bool),
			flag.Description,
		)
	}
}

func runInit(cmd *cobra.Command, args []string) error {
	templateName := args[0]
	projectName := args[1]

	meta, err := internal.LoadMetadata(templateName)
	if err != nil {
		return err
	}
	data := map[string]any{
		"ProjectName": projectName,
	}
	for name, variable := range meta.Variables {
		if variable.Source == "flag" {
			data[name] = cmd.Flag(variable.Flag).Value.String()
		}
	}
	for name := range meta.Flags {
		val, _ := cmd.Flags().GetBool(name)
		data[name] = val
	}

	src := "templates/" + templateName
	return copyTemplate(src, projectName, data)

}
