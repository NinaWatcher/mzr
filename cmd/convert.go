package cmd

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"errors"
	"github.com/spf13/cobra"
	"github.com/ghodss/yaml"
)

var target string
var outputDir string

var cmdBuild = &cobra.Command{
	Use:   "convert",
	Short: "convert project",
	Run:   nil,
}

var cmdConvert = &cobra.Command{
	Use:   "toJSON",
	Short: "Convert options",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		convert(args)
	},
}

func convert(args []string) {
	 if len(args[0]) > 0{
		 //read file
 		yamlFile, err := ioutil.ReadFile(args[0])
 		AbortIfError(err, "failed to read the yaml file from the provided location")
 		//Format the yaml to json
 		fmt.Println("Start converting: " + strings.Join(args, " "))
 		jsonOut, err := yaml.YAMLToJSON(yamlFile)
 		AbortIfError(err, "failed to convert YAML to JSON")
 		err = ioutil.WriteFile("jsonFile.json", jsonOut, 0644)
		AbortIfError(err, "failed to write converted JSON to file")
		fmt.Println("YAML converted to jsonFile.json")
	 } else {
		 AbortIfError(errors.New(""), "please provide a valid path to the yml file")
	 }
}

func init() {
	cmdBuild.AddCommand(cmdConvert)
	rootCmd.AddCommand(cmdBuild)
	cmdBuild.PersistentFlags().StringVarP(&outputDir, "convert", "o", "", "filesystem path to write files to")
}

func AbortIfError(err error, msg string) {
	if err != nil{
		fmt.Printf("[ERROR] %s", msg)
		if "dev" == os.Getenv("RUN_MODE"){
			fmt.Printf("[ERROR] msg: %s reason: %s\n", msg, err.Error())
		}
		os.Exit(1)
	}
}
