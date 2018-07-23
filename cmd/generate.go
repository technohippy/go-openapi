package cmd

import (
	"../gen"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var inputFile, outputDir, templateDir, language string

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&inputFile, "input-spec", "i", "", "location of the swagger spec, as URL or file (required)")
	generateCmd.Flags().StringVarP(&outputDir, "output", "o", "./", "where to write the generated files")
	generateCmd.Flags().StringVarP(&templateDir, "template-dir", "t", "", "folder containing the template files")
	generateCmd.Flags().StringVarP(&language, "lang", "l", "", "client language to generate (maybe class name in classpath, required)")
	generateCmd.MarkFlagRequired("input-spec")
	generateCmd.MarkFlagRequired("language")
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate openapi template",
	Long:  `generate openapi template`,
	Run: func(cmd *cobra.Command, args []string) {
		ext := filepath.Ext(inputFile)
		switch ext {
		case ".json":
			fmt.Println("json")
		case ".yaml", ".yml":
			fmt.Println("yaml")
		default:
			fmt.Printf("Invalid filetype (%v): Only json and yaml are acceptalbe.\n", ext)
			// err
			return
		}

		conf, err := gen.ReadConfig(inputFile)
		if err != nil {
			// err
			fmt.Printf("%v\n", err)
			return
		}
		generator := gen.NewGenerator("templates/go") // TODO
		result, err := generator.Generate(conf)
		if err != nil {
			// err
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Println(result)
	},
}
