package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aerissecure/csrange"
)

// TODO: permit piping stdin: https://flaviocopes.com/go-shell-pipes/

func main() {
	integers := flag.Bool("i", false, "integers")
	count := flag.Bool("c", false, "count")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, "please provide a comma separated range as an argument")
		os.Exit(1)
	}
	csr := flag.Args()[0]
	ints, err := csrange.Ints(csr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing error: %s", err)
		os.Exit(1)
	}
	if *integers {
		sints := []string{}
		for _, v := range ints {
			sints = append(sints, fmt.Sprintf("%d", v))
		}
		fmt.Println(strings.Join(sints, ","))
	} else {
		fmt.Println(csrange.CSR(ints))
	}
	if *count {
		fmt.Printf("\ncount: %d\n", len(ints))
	}
}
