/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"sort"

	"github.com/docker/docker/daemon/logger"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func list() {
	sort.Sort(DFslice(data))
	text, err := PrettyPrintJSONstream(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)

	logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
	// Work with logger
	if disableLogging == false {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}

	slog.SetDefault(logger)
	s := fmt.Sprintf("%d records in total.", len(data))
	logger.Info(s)
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err == nil {
		log.Println(string(b))
	}
	return err
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
