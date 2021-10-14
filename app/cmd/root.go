package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	Version = "1.0.0"
	rootCmd = &cobra.Command{
		Use: "hathor",
		Version: Version,
		Short: "hathor management cli",
	}
	Env = "alpha"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	flagSet := rootCmd.PersistentFlags()
	flagSet.StringVarP(&Env,"env","e", os.Getenv("ENV"), "环境: dev, test, alpha, prod")
	flagSet.StringVarP(&Version,"version", "v", os.Getenv("VERSION"), "版本: 1.0.0")
}
