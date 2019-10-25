package cmd

import (
	"log"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"bufio"
	"encoding/csv"
	"encoding/json"

	"fmt"
	"os"

	grok "github.com/vjeantet/grok"
)

var (
	config             grok.Config
	pattern            string
	multiLinePattern   string
	outputType         string
	additionalPatterns []string
)

func init() {
	log.SetFlags(0)
	parseCmd.PersistentFlags().BoolVar(&config.NamedCapturesOnly, "named-captures-only", true, "only include named capture groups in returned data")
	parseCmd.PersistentFlags().BoolVar(&config.SkipDefaultPatterns, "skip-default-patterns", false, "skip default patterns")
	parseCmd.PersistentFlags().BoolVar(&config.RemoveEmptyValues, "remove-empty-values", true, "do not include empty values in returned data")
	parseCmd.PersistentFlags().StringArrayVarP(&config.PatternsDir, "patterns-dir", "d", nil, "directory to with additional grok patterns")
	parseCmd.PersistentFlags().StringVarP(&pattern, "pattern", "p", "", "pattern to match")
	parseCmd.PersistentFlags().StringVarP(&multiLinePattern, "multi-line-pattern", "m", "", "pattern to mark the beginning of a multiline grok")
	parseCmd.PersistentFlags().StringVarP(&outputType, "output-type", "o", "json", "output type csv or json")

	// TODO: add support for "--additional-pattern"
	parseCmd.PersistentFlags().StringArrayVarP(&additionalPatterns, "additional-pattern", "a", nil, "additional grok pattern to reference")

	rootCmd.AddCommand(parseCmd)

}

// process the parse command
var parseCmd = &cobra.Command{
	Use:   "parse <files>",
	Short: "parse log files using grok",
	Run:   doWork,
}

func doWork(cmd *cobra.Command, args []string) {
	g, err := grok.NewWithConfig(&config)
	check(err)

	for _, pattern := range additionalPatterns {
		words := strings.Fields(pattern)
		reg := regexp.MustCompile(`^\S* `)
		res := reg.ReplaceAllString(pattern, "")
		g.AddPattern(words[0], res)
	}

	for _, arg := range args {

		file, err := os.Open(arg)
		check(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		str := ""
		for scanner.Scan() {
			var values map[string]string
			if multiLinePattern != "" {
				v, err := g.Parse(multiLinePattern, scanner.Text())
				check(err)
				if len(v) > 0 {
					values, err = g.Parse(pattern, str)
					check(err)
					str = scanner.Text()
				} else {
					str = str + scanner.Text()
					continue
				}
			} else {
				values, err = g.Parse(pattern, str)
			}

			// TODO: enum these values
			if outputType == "json" {
				outputJSON(values)
			} else if outputType == "csv" {
				outputCSV(values)
			} else {
				log.Fatalf("unknown output type: %s", outputType)
			}

		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

}

func outputJSON(values map[string]string) {
	jsonString, err := json.Marshal(values)
	check(err)
	fmt.Println(string(jsonString))
}

func outputCSV(values map[string]string) {
	writer := csv.NewWriter(os.Stdout)
	var record []string

	record = append(record, values["timestamp"])
	record = append(record, values["level"])
	record = append(record, values["thread"])
	record = append(record, values["class"])
	record = append(record, values["details"])
	writer.Write(record)
	writer.Flush()
}

// check for errors
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
