package cmd

import (
	"bytes"
	"fmt"
	"forge/internal"
	"text/template"

	"github.com/pterm/pterm"

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
	if len(args) < 2 {
		pterm.Error.Println("project-name not provided")
		pterm.Info.Printfln("Use forge 'info <command>' to see its available flags")
		return nil
	}
	templateName := args[0]
	projectName := args[1]

	meta, err := internal.LoadMetadata(templateName)
	if err != nil {
		return err
	}

	registerFlags(cmd, meta)

	if err := cmd.Flags().Parse(args[2:]); err != nil {
		return err
	}

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

	spinner, _ := pterm.DefaultSpinner.Start("Creating project...")
	err = copyTemplate(src, projectName, data)

	if err != nil {
		spinner.Fail("Failed to create project")
		return err
	}

	spinner.Success("Project files generated")
	postCommandPrint(*meta, data)
	return err
}

func postCommandPrint(meta internal.TemplateMetadata, data map[string]any) {
	if len(meta.PostCreate) > 0 {
		pterm.DefaultSection.Println("Next steps")
		for _, step := range meta.PostCreate {
			tmpl, _ := template.New("step").Parse(step)

			var buf bytes.Buffer
			_ = tmpl.Execute(&buf, data)

			pterm.Info.Printf("  %s\n", buf.String())
		}
	}
}
