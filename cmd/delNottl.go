/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"RedisDBA/pkg"
	"github.com/spf13/cobra"
)

// delNottlCmd represents the delNottl command
var delNottlCmd = &cobra.Command{
	Use:   "delNottl",
	Short: "Delete keys with no TTL set",
	Long: `Delete keys with no ttl set
For example:
./redisDBA delNottl --config /etc/config.yml
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := pkg.InitClient(c)
		if err != nil {
			panic(err)
		}
		pkg.DelNoTTLPre()
	},
}

func init() {
	rootCmd.AddCommand(delNottlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delNottlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delNottlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
