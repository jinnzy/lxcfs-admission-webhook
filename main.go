package main

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"lxcfs-admission-webhook/webhook"
	"os"
)

var Version = "development"

func main() {
	rootCmd := &cobra.Command{
		Use:     "app",
		Version: Version,
	}
	rootCmd.AddCommand(webhook.CmdWebhook)
	if err := rootCmd.Execute(); err != nil {
		klog.Error(err)
		os.Exit(1)
	}
	os.Exit(0)
}
