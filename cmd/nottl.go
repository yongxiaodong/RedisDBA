/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"RedisDBA/pkg"
	"github.com/spf13/cobra"
)

// nottlCmd represents the nottl command
var nottlCmd = &cobra.Command{
	Use:   "nottl",
	Short: "SCAN No TTL Key",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("nottl called")
		err := pkg.InitClient(c)
		if err != nil {
			panic(err)
		}
		pkg.QueryNoTtlKey()
	},
}

func init() {
	rootCmd.AddCommand(nottlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nottlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nottlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
