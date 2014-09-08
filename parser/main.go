package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/vierja/logprecios-parsers"
	"os"
)

var url = flag.String("url", "", "Url to fetch")

func main() {
	flag.Parse()
	if *url != "" {

		prod, err := scraping.GetProductData(*url)
		if err != nil {
			fmt.Printf("Error while fetching url: %s: %s", *url, err.Error())
			os.Exit(1)
		}

		prettyJson, err := json.MarshalIndent(prod, "", "  ")
		if err != nil {
			fmt.Printf("Error while generating json: %s", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Product %s:\n %s\n", *url, prettyJson)
	}
}
