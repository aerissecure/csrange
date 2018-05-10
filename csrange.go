package csrange

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// csr is comma separated range, e.g.: 1,4-39,199,200-201,400
// because these are ranges using a dash, they are only meant
// for positive integers.
// Returned ints are sorted and unique.
func Ints(csr string) ([]int, error) {
	csr = strings.Trim(csr, ",")
	ints := []int{}
	tokens := strings.Split(csr, ",")
	for _, tok := range tokens {
		switch {
		case strings.Contains(tok, "-"):
			fields := strings.Split(tok, "-")
			if len(fields) != 2 {
				return ints, fmt.Errorf("illegal range entry: %s", tok)
			}
			first, err := strconv.Atoi(fields[0])
			if err != nil {
				return ints, fmt.Errorf("first value is not an integer: %s", tok)
			}
			last, err := strconv.Atoi(fields[1])
			if err != nil {
				return ints, fmt.Errorf("last value is not an integer: %s", tok)
			}
			if last < first {
				return ints, fmt.Errorf("first value is greater than second value: %s", tok)
			}
			for i := first; i <= last; i++ {
				ints = append(ints, i)
			}
		default:
			i, err := strconv.Atoi(tok)
			if err != nil {
				return ints, fmt.Errorf("value is not an integer: %s", tok)
			}
			ints = append(ints, i)
		}
	}
	sort.Ints(ints)
	ii := []int{}
	for _, i := range ints {
		ii = uappend(ii, i)
	}
	return ii, nil
}

// uappend appends the element if it is uniqe
func uappend(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

// fmtRange formats a range or signle integer as a string based on the inputs
func fmtRange(start, end int) string {
	if start == end {
		return fmt.Sprintf("%d", start)
	}
	return fmt.Sprintf("%d-%d", start, end)
}

func CSR(ss []int) string {
	sort.Ints(ss)
	ints := []int{}
	for _, i := range ss {
		ints = uappend(ints, i)
	}
	csr := []string{}
	s := -1
	for i, v := range ints {
		if s == -1 {
			s = v
		}
		if i+1 == len(ints) {
			// append final range
			csr = append(csr, fmtRange(s, v))
			break
		}
		if ints[i+1] != v+1 {
			// end of a sequence
			csr = append(csr, fmtRange(s, v))
			s = -1
		}
	}
	return strings.Join(csr, ",")
}
