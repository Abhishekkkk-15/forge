package cmd

import (
	"fmt"
	"forge/internal"

	"github.com/spf13/cobra"
)

func registerFlags(cmd *cobra.Command, meta *internal.TemplateMetadata) {
	for _, variable := range meta.Variables {
		if variable.Source != "flag" {
			continue
		}

		switch variable.Type {
		case "string":
			cmd.Flags().String(
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
		cmd.Flags().Bool(
			name,
			flag.Default.(bool),
			flag.Description,
		)
	}
}

func runInit(cmd *cobra.Command, args []string) error {
	if args[1] == " " {
		return fmt.Errorf("project name not provided")
	}
	templateName := args[0]
	projectName := args[1]

	meta, err := internal.LoadMetadata(templateName)
	if err != nil {
		return err
	}

	// 1️⃣ Register flags dynamically
	registerFlags(cmd, meta)

	// 2️⃣ Now parse flags manually
	if err := cmd.Flags().Parse(args[2:]); err != nil {
		return err
	}

	// 3️⃣ Build template data
	data := map[string]any{
		"ProjectName": projectName,
		"Description": meta.Description,
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
