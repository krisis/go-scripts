package main
import (
	"strings"
	"fmt"
	"time"
	"github.com/bitfield/script"
	"log"
)

func main() {
	var (
		curTime time.Time
		idx int
	)
	// "2023-01-30T04:16:06.980055335-08:00"
	//  "Mon Jan _2 15:04:05 2006"
	const layout = "2006-01-02T15:04:05.999999999-07:00"

	// line format: 2023-03-01T08:24:50.001917708-08:00
	script.Stdin().EachLine(func(line string, out *strings.Builder) {
		t, err := time.Parse(layout, line)
		if err != nil {
			log.Fatal(idx, line, err)
		}

		elapsed := time.Duration(0)
		if idx > 0 {
			elapsed = curTime.Sub(t)
		}
		fmt.Fprintf(out, "%d: %v minutes ago\n", idx, elapsed.Round(time.Minute).Minutes())
		idx++
		curTime = t
	}).Stdout()
}
