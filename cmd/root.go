/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"RedisDBA/pkg"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	configFile string
	outputDir  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "RedisDBA",
	Short: "Please inpu paramter",
	Long: `ERROR, For example:
./RedisDBA nottl`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.RedisDBA.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cobra.OnInitialize(InitConfig)
	//rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file path(default is ./conf/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file path(default is ./conf/config.yaml)")
	//rootCmd.PersistentFlags().StringVar(&outputDir, "output", "", "output directory path(default is ./result/)")
}

var c pkg.Config

func InitConfig() {
	var file string
	if configFile != "" {
		file = configFile
	} else {
		file = fmt.Sprintf("%s/conf/config.yml", pkg.GetExcPath())
	}
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		log.Printf("Read Config File Error: #%v. also can be specified with --config config path", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Println("Parse Config ERROR: ", err)
	}

}
