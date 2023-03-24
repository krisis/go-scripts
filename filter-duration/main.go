package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/bitfield/script"
)

func main() {
	var duration time.Duration
	var report bool
	flag.DurationVar(&duration, "le", time.Second, "filter (out) durations lesser than given duration")
	flag.BoolVar(&report, "report", false, "print out a short report of min, max and avg duration")
	flag.Parse()

	var idx int
	var max, sum time.Duration
	min := time.Duration(math.MaxInt64)
	script.Stdin().EachLine(func(line string, out *strings.Builder) {
		d, err := time.ParseDuration(line)
		if err != nil {
			log.Fatalf("line no: %d, %v", idx, err)
		}
		if d < duration {
			return
		}

		idx++
		if min > d {
			min = d
		}
		if max < d {
			max = d
		}
		sum += d
		if !report {
			fmt.Fprintln(out, d)
		}

	}).Stdout()
	if report {
		avg := int64(sum) / int64(idx)
		fmt.Println("Report:")
		fmt.Printf("total samples = %d, min = %v, max = %v, avg = %v\n", idx, min, max, time.Duration(avg))
	}
}
