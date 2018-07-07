package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aerissecure/csrange"
)

func stdin() string {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return ""
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return ""
	}

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(buf))
}

func main() {
	integers := flag.Bool("i", false, "integers")
	count := flag.Bool("c", false, "count")
	flag.Parse()

	csr := stdin()
	// stdin taxes precedence over arg
	if len(flag.Args()) > 0 && len(csr) == 0 {
		csr = flag.Args()[0]
	}

	if csr == "" {
		flag.Usage()
		// fmt.Fprintln(os.Stderr, "please provide a comma separated range via arg or stdin")
		os.Exit(1)
	}

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
