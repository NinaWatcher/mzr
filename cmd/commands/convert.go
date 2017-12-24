package commands

import (
	"io/ioutil"
	"log"
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	YAML "github.com/ghodss/yaml"
	"github.com/mitchellh/go-homedir"
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

	home, err := homedir.Dir()
	yamlDir := home + args[0] + "file.yaml"

	//read file
	yamlFile, err := ioutil.ReadFile(yamlDir)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	//Format the yaml to json
	fmt.Println("Start converting: " + strings.Join(args, " "))
	jsonOut, err := YAML.YAMLToJSON(yamlFile)
	fmt.Println(string(jsonOut))
	err = ioutil.WriteFile("jsonFile.json", jsonOut, 0644)
}

func init() {
	cmdBuild.AddCommand(cmdConvert)
	rootCmd.AddCommand(cmdBuild)
	cmdBuild.PersistentFlags().StringVarP(&outputDir, "convert", "o", "", "filesystem path to write files to")
}
