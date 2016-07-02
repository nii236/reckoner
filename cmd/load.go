package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type quad struct {
	Subj string `json:"subject"`
	Pred string `json:"predicate"`
	Obj  string `json:"object"`
}

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		file := viper.GetString("file")
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		bb := bytes.NewBuffer(b)
		s := bufio.NewScanner(bb)
		qs := []quad{}
		for s.Scan() {
			sarr := strings.Split(s.Text(), " ")
			if len(sarr) != 4 {
				log.Fatalln("Wrong length for RDF data")
			}
			qs = append(qs, quad{
				Subj: sarr[0],
				Pred: sarr[1],
				Obj:  sarr[2],
			})
		}
		b2, err := json.Marshal(qs)
		log.Println(qs)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := http.Post("http://127.0.0.1:64210/api/v1/write", "application/json", strings.NewReader(string(b2)))
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		var responseBody []byte
		resp.Body.Read(responseBody)
		log.Println(string(responseBody))
	},
}

func init() {
	RootCmd.AddCommand(loadCmd)
	loadCmd.Flags().StringP("file", "f", "", "A file to load")
	viper.BindPFlags(loadCmd.Flags())
}
