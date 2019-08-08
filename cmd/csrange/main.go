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
	split := flag.Int("s", 0, "split output into this many equal size bucket")
	bucket := flag.Int("b", 0, "return the csr for only this bucket")
	flag.Parse()

	csr := stdin()
	// stdin takes precedence over arg
	if len(flag.Args()) > 0 && len(csr) == 0 {
		csr = flag.Args()[0]
	}

	if csr == "" {
		flag.Usage()
		// fmt.Fprintln(os.Stderr, "please provide a comma separated range via arg or stdin")
		os.Exit(1)
	}

	ints, err := csrange.Ints(csr)

	if *bucket > *split {
		fmt.Println("value for bucket (-b) must be <= value for split (-s)")
		os.Exit(1)
	}

	if *split > 0 {
		bckts := buckets(*split, ints)
		if *bucket == 0 {
			for _, b := range bckts {
				fmt.Println(csrange.CSR(b))
			}
		}
		if *bucket > 0 {
			fmt.Println(csrange.CSR(bckts[*bucket-1]))
		}
		return
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing error: %s", err)
		os.Exit(1)
	}
	if *integers {
		sInts := toStrings(ints)
		fmt.Println(strings.Join(sInts, ","))
	} else {
		fmt.Println(csrange.CSR(ints))
	}
	if *count {
		fmt.Printf("\ncount: %d\n", len(ints))
	}

}

func toStrings(ints []int) []string {
	out := []string{}
	for _, i := range ints {
		out = append(out, fmt.Sprintf("%d", i))
	}
	return out
}

func buckets(count int, ints []int) [][]int {
	if count == 0 {
		return make([][]int, 0)
	}

	if count == 1 {
		return [][]int{ints}
	}

	buckets := make([][]int, count)
	for i, v := range ints {
		buckets[i%count] = append(buckets[i%count], v)
	}
	return buckets
}
