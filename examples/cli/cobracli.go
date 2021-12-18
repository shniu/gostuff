package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var endpoints []string

var rootCmd = &cobra.Command{
	Use:        "myctl",
	Short:      "myctl is a custom cli.",
	SuggestFor: []string{"mycli"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("myctl cmd -> ", cmd.Use)
	},
}

func init() {
	rootCmd.PersistentFlags().StringSliceVar(&endpoints, "endpoints", []string{"127.0.0.1:2379"}, "Endpoints for connecting.")

	rootCmd.AddCommand(getCommand())
}

func getCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [options] <key> [range_end]",
		Short: "Gets the key or a range of keys",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("getCommand")
		},
	}

	return cmd
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
