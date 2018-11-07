package validate

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/spf13/cobra"
	"github.com/waterborne-labs/instrument-flight-rules/cmd/lib"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// https://stackoverflow.com/questions/40737122/convert-yaml-to-json-without-struct
func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

func NewValidateCmd(box packr.Box) *cobra.Command {
	var taskSchema bool
	var yamlInput bool

	cmd := &cobra.Command{
		Use:   "validate PATHS...",
		Short: "Validate one or more definition files",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var schemaFile string

			if taskSchema {
				schemaFile = "resource.json"
			} else {
				schemaFile = "task.json"
			}

			schemaContents, err := box.FindString(schemaFile)
			if err != nil {
				panic(err.Error())
			}

			schemaLoader := gojsonschema.NewStringLoader(schemaContents)

			for _, path := range args {
				exists, _ := lib.Exists(path)
				if !exists {
					panic(fmt.Sprintf("File at %s does not exist", path))
				}
				b, err := ioutil.ReadFile(path)
				c := string(b)

				if yamlInput {
					var body interface{}
					if err := yaml.Unmarshal([]byte(c), &body); err != nil {
						panic(err)
					}

					body = convert(body)

					jb, err := json.Marshal(body)
					if err != nil {
						panic(err)
					}

					c = string(jb)
				}

				documentLoader := gojsonschema.NewStringLoader(c)

				result, err := gojsonschema.Validate(schemaLoader, documentLoader)
				if err != nil {
					panic(err.Error())
				}

				if result.Valid() {
					fmt.Printf("The document is valid\n")
				} else {
					fmt.Printf("The document is not valid. see errors :\n")
					for _, desc := range result.Errors() {
						fmt.Printf("- %s\n", desc)
					}
				}
			}
		},
	}

	cmd.Flags().BoolVarP(&taskSchema, "task", "t", true, "Validate a task instead of a resource")
	cmd.Flags().BoolVarP(&yamlInput, "yaml", "y", true, "Parse input files as yaml")

	return cmd
}
