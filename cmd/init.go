package cmd

import (
	"bytes"
	"fmt"
	"forge/internal"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	port  int
	useTs bool
)

var initCmd = &cobra.Command{
	Use:   "init <template> <project-name>",
	Short: "Create a new project from a template",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		templateName := args[0]
		projectName := args[1]

		src := "templates/" + templateName
		if templateName == "express" && useTs {
			src = "templates/express-ts"
		}
		if _, err := internal.TemplateFS.ReadDir(src); err != nil {
			return fmt.Errorf("template '%s' not found", templateName)
		}

		if _, err := os.Stat(projectName); err == nil {
			return fmt.Errorf("directory '%s' already exists", projectName)
		}

		data := map[string]any{
			"ProjectName": projectName,
			"Port":        port,
			"UseTs":       useTs,
		}

		return copyTemplate(src, projectName, data)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().IntVar(
		&port,
		"port",
		8000,
		"Port number of the server",
	)
	initCmd.Flags().BoolVar(
		&useTs,
		"ts",
		false,
		"Use typescript template",
	)
}

func copyTemplate(src, dest string, data map[string]any) error {
	return fs.WalkDir(internal.TemplateFS, src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath := path[len(src):]
		targetPath := filepath.Join(dest, relPath)

		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		content, err := internal.TemplateFS.ReadFile(path)
		if err != nil {
			return err
		}

		tmpl, err := template.New("file").Parse(string(content))
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			return err
		}

		return os.WriteFile(targetPath, buf.Bytes(), 0644) // lilepath | data | file permission's rw-r--r--
	})
}
