package cmd

import (
	"github.com/pinkhello/hathor/core"
	"github.com/spf13/cobra"
)

var (
	httpCmd = &cobra.Command{
		Use:   "http",
		Short: "start http",
		Run:   initHttp,
	}
)

func initHttp(cmd *cobra.Command, args []string) {
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		panic(err)
	}
	core.Start(env)
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
