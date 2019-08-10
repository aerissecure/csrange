package csrange

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// csr is comma separated range, e.g.: 1,4-39,199,200-201,400
// because these are ranges using a dash, they are only meant
// for positive integers.

// Ints returns the slice of integers represented by the CSR. Any invalid
// input results in an error. The result is a sorted and unique list.
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

// uappend appends the element if it is unique
func uappend(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

// fmtRange formats a range or single integer as a string based on the inputs
func fmtRange(start, end int) string {
	if start == end {
		return fmt.Sprintf("%d", start)
	}
	return fmt.Sprintf("%d-%d", start, end)
}

// CSR produces an ordered, unique, and compact CSR representing
// a slice of integers.
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

// Split divides the ints into n number of buckets
func Split(n int, ints []int) [][]int {
	if n == 0 {
		return make([][]int, 0)
	}

	if n == 1 {
		return [][]int{ints}
	}

	buckets := make([][]int, n)
	for i, v := range ints {
		buckets[i%n] = append(buckets[i%n], v)
	}
	return buckets
}

// SplitContig is like Split only it uses as many contiguous ints as it can to reduce the sice of the csr.
func SplitContig(n int, ints []int) [][]int {
	if n == 0 {
		return make([][]int, 0)
	}

	if n == 1 {
		return [][]int{ints}
	}

	count := int(math.Ceil(float64(len(ints)) / float64(n)))

	buckets := make([][]int, n)
	j := 0
	for i := range buckets {
		for ; j < len(ints); j++ {
			// fmt.Println("i:", i, "j:", j, "count:", count, "mod:", (j+1)%count)
			buckets[i] = append(buckets[i], ints[j])
			if (j+1)%count == 0 {
				j++
				break
			}
		}
	}
	return buckets
}

// SplitString divides CSR string into n number of csr strings
func SplitString(n int, csr string) ([]string, error) {
	// parse csr:
	ints, err := Ints(csr)
	if err != nil {
		return nil, err
	}
	buckets := Split(n, ints)
	out := []string{}
	for _, b := range buckets {
		out = append(out, CSR(b))
	}
	return out, nil
}

// SplitString divides CSR string into n number of csr strings
func SplitStringContig(n int, csr string) ([]string, error) {
	// parse csr:
	ints, err := Ints(csr)
	if err != nil {
		return nil, err
	}
	buckets := SplitContig(n, ints)
	out := []string{}
	for _, b := range buckets {
		out = append(out, CSR(b))
	}
	return out, nil
}
