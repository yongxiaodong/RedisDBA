/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"RedisDBA/pkg"
	"github.com/spf13/cobra"
)

// bigkeyCmd represents the bigkey command
var bigkeyCmd = &cobra.Command{
	Use:   "bigkey",
	Short: "scan Big Key TOP 50",
	Long: `scan Big Key TOP 50
and usage of using your command. For example:

./redisDBA bigkey`,
	Run: func(cmd *cobra.Command, args []string) {
		err := pkg.InitClient(c)
		if err != nil {
			panic(err)
		}
		pkg.BigKeyTOP()
	},
}

func init() {
	rootCmd.AddCommand(bigkeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bigkeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bigkeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
