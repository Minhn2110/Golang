package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use: "Hello",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			rootCmd.PersistentFlags().String("author", args[1], "Author name for copyright attribution")
			_ = viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))

			fmt.Println("Root Command", args)
		},
		Example: "go run main.go author Minh",
	}
}

func Execute() error {
	return rootCmd.Execute()
}
