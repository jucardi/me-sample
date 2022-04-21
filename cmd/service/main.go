package main

import (
	"fmt"

	"github.com/jucardi/go-titan/configx"
	"github.com/jucardi/go-titan/info"
	"github.com/jucardi/go-titan/logx"
	"{{.golang.module_path}}/{{.service_name}}/server"
	"github.com/spf13/cobra"
)

const (
	Usage = `
Docker Swarm manager for AWS Autoscaling groups
  - Version: %s
  - Built: %s
`
)

var rootCmd = &cobra.Command{
	Use:   "ms-sample",
	Short: "Docker Swarm manager for AWS Autoscaling groups",
	Long:  fmt.Sprintf(Usage, info.Version, info.Built),
	Run:   run,
}

func main() {
	// Loads the config related flags to the cobra command.
	configx.LoadFlagsToCommand(rootCmd)

	// Sub-Commands
	//
	//    In this section, add any additional sub-commands for cobra, for example:
	//       $ ms-sample bgprocess
	//    The "bgprocess" CLI command would be added by creating a new *cobra.Command
	//    and them doing:    rootCmd.AddCommand(bgProcessCmd)

	if err := rootCmd.Execute(); err != nil {
		rootCmd.Println()
		rootCmd.Println(rootCmd.UsageString())
	}
}

func run(cmd *cobra.Command, _ []string) {
	logx.WithObj(
		configx.FromCommand(cmd),
	).Fatal("failed to load configuration")

	server.Run()
}
