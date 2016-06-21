package cmd

import (
	"github.com/nii236/btcer/parse"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Begin parsing the blockchain",
	Long: `Begin parsing the blockchain.

This requires a full node BTCD instance.`,
	Run: func(cmd *cobra.Command, args []string) {
		start := viper.GetInt("start")
		end := viper.GetInt("end")
		parse.Fetch(start, end)
	},
}

func init() {
	RootCmd.AddCommand(parseCmd)
	parseCmd.Flags().IntP("start", "s", 0, "Start block")
	parseCmd.Flags().IntP("end", "e", 10000, "End block")
	viper.BindPFlags(parseCmd.Flags())

}
