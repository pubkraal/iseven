package iseven

import "github.com/pubkraal/isodd/pkg/isodd"

func IsEven(num int) bool {
	odd := isodd.IsOdd(num)
	return !odd
}
