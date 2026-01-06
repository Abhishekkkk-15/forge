package cmd

import (
	"bytes"
	"forge/internal"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:                "init <template> <project-name>",
	Short:              "Create a new project from a template",
	DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
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
