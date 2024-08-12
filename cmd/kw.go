/*
Copyright © 2024 NullNUMMER24
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// kwCmd represents the kw command
var kwCmd = &cobra.Command{
	Use:   "kw",
	Short: "Get some information about the date",
	Run: func(cmd *cobra.Command, args []string) {

		ts := time.Now().UTC().Unix()
		tn := time.Unix(ts, 0)
		parts := strings.Split(fmt.Sprintf("%v", tn), " ")
		year, week := tn.ISOWeek()
		fmt.Printf("Time: %s\nYear: %v\nWeek: %v\n", parts[1], year, week)
	},
}

func init() {
	rootCmd.AddCommand(kwCmd)

}
