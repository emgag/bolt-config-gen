package cmd

import (
	"github.com/spf13/cobra"
	"path/filepath"
	"os"
	"fmt"
	"strings"
	"log"
	"github.com/Masterminds/sprig"
	"text/template"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate <env> <dir>",
	Short: "Generate config",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		env := args[0]
		dir := args[1]

		searchSuffix := fmt.Sprintf("_%s.yml.tmpl", env)
		replaceSuffix := "_local.yml"

		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.HasSuffix(path, searchSuffix) {
				fmt.Printf("%s\n", path)

				tpl := template.Must(
					template.New(filepath.Base(path)).Funcs(sprig.TxtFuncMap()).ParseFiles(path),
				)

				outFile := strings.TrimSuffix(path, searchSuffix) + replaceSuffix

				f, err := os.Create(outFile)
				defer f.Close()

				if err != nil {
					return err
				}

				return tpl.Execute(f, nil)
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}

	},
}
