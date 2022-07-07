package speed_test

import (
	"fmt"
	"time"
)

func SpeedTest(f func(int) int, testcase []int) {
	fmt.Println("--------------------")
	for _, v := range testcase {
		t0 := time.Now()
		res := f(v)
		elapsed := time.Since(t0) * time.Microsecond
		fmt.Println("inp: ", v, ": res:  ", res, "\t. ExecTime: \t", elapsed)
	}
	fmt.Println("--------------------")
}
