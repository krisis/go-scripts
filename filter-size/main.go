package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/bitfield/script"
	"github.com/dustin/go-humanize"
)

func main() {
	var report bool
	flag.BoolVar(&report, "report", false, "print out a short report of min, max and avg duration")
	flag.Parse()

	var idx int
	var max, sum uint64
	var min uint64 = math.MaxUint64
	script.Stdin().EachLine(func(line string, out *strings.Builder) {
		sz, err := humanize.ParseBytes(line)
		if err != nil {
			log.Fatalf("line no: %d, %v", idx, err)
		}

		idx++
		if min > sz {
			min = sz
		}
		if max < sz {
			max = sz
		}
		sum += sz
		if !report {
			fmt.Fprintln(out, sz)
		}

	}).Stdout()
	if report {
		avg := sum / uint64(idx)
		fmt.Println("Report:")
		fmt.Printf("total samples = %d, min = %v, max = %v, avg = %s\n", idx, humanize.IBytes(min), humanize.IBytes(max), humanize.IBytes(avg))
	}
}
