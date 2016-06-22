package cmd

import (
	"log"

	"github.com/nii236/btcer/parse"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Begin parsing the blockchain",
	Long: `Begin parsing the blockchain.

This requires a full node BTCD instance.

By default, only the relations between types will be generated. Actual
literals will only be generated when the corresponding flag is set.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			log.Fatalln("Invalid arguments")
		}
		start := viper.GetInt("start")
		end := viper.GetInt("end")
		parse.Fetch(start, end)
	},
}

func init() {
	RootCmd.AddCommand(parseCmd)
	parseCmd.Flags().IntP("start", "s", 0, "Start block")
	parseCmd.Flags().IntP("end", "e", 10000, "End block")
	parseCmd.Flags().BoolP("block", "b", false, "Parse block literals")
	parseCmd.Flags().BoolP("tx", "t", false, "Parse transaction literals")
	parseCmd.Flags().BoolP("vin", "i", false, "Parse vin literals")
	parseCmd.Flags().BoolP("vout", "o", false, "Parse vout literals")
	parseCmd.Flags().BoolP("scriptpubkey", "p", false, "Parse scriptpubkey literals")
	parseCmd.Flags().BoolP("address", "a", false, "Parse address literals")
	viper.BindPFlags(parseCmd.Flags())

}
