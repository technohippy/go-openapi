package gen

import (
	"fmt"
	"github.com/cbroglie/mustache"
	"io/ioutil"
	//"os"
	"path/filepath"
)

type Generator struct {
	TemplateDir string
}

func NewGenerator(templateDir string) *Generator {
	return &Generator{templateDir}
}

func (g *Generator) Generate(config *OpenAPIObject) (string, error) {
	fileInfos, err := ioutil.ReadDir(g.TemplateDir)
	if err != nil {
		return "", err
	}
	for _, fileInfo := range fileInfos {
		if filepath.Ext(fileInfo.Name()) == ".mustache" {
			content, err := g.GenerateFromFile(fileInfo.Name(), config)
			if err != nil {
				return "", err
			}
			/*
				file, err := os.Create(fmt.Sprintf("./%v.generated.txt", fileInfo.Name()))
				if err != nil {
					return "", err
				}
				defer file.Close()
				file.Write(([]byte)(content))
			*/
			fmt.Println(content)
		}
	}
	return "", nil // TODO
}

func (g *Generator) GenerateFromFile(filename string, config *OpenAPIObject) (string, error) {
	fmt.Printf("generating %v...\n", filename)
	if filename != "README.mustache" { // TODO
		return "", nil
	}
	buf, err := ioutil.ReadFile(filepath.Join(g.TemplateDir, filename))
	if err != nil {
		return "", err
	}
	template := string(buf)
	renderer := &Renderer{}
	return renderer.Render(template, config)
}

type Renderer struct {
}

func (r *Renderer) Render(template string, config *OpenAPIObject) (string, error) {
	return mustache.Render(template, config.TemplateVariables())
}
