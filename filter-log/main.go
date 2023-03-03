package main
import (
	"strings"
	"fmt"
	"time"
	"flag"
	"github.com/bitfield/script"
	"log"
)

type logTime time.Time

func (l *logTime) Set(s string) error {
	const layout = "2006-01-02T15:04:05.99.999"
	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	*l = logTime(t)
	return nil
}
func (l *logTime) String() string {
	if l != nil {
		return fmt.Sprint(*l)
	}
	return ""
}

func main() {
	since := logTime{}
	until := logTime{}
	flag.Var(&since,"since", "filter logs since this time")
	flag.Var(&since,"s", "filter logs since this time")
	flag.Var(&until, "until", "filter logs until this time")
	flag.Var(&until, "u", "filter logs until this time")
	flag.Parse()

	// 2023-03-02T20:48:27.781
	const layout = "2006-01-02T15:04:05"

	// line format: 2023-03-02T20:48:27.781 [404 Not Found] s3.HeadObject cdc1and1-ss-0-3.cdc1and1-hl.ns-sy-ros-minio.svc.cluster.local:9000/fmaas/checkpoints/checkpointdir1n1.cdc.oss.apic.alert/fm-streaming/offsets/47684 2a00:fbc:2110:14b3:b412:22:2:0                  19.785ms     ↑ 153 B ↓ 291 B
	var idx int
	script.Stdin().EachLine(func(line string, out *strings.Builder) {
		tokens := strings.Split(line, " ")
		if len(tokens) <= 1 {
			log.Fatal(line, "not enough tokens")
		}
		t, err := time.Parse(layout, tokens[0])
		if err != nil {
			log.Fatal(idx, tokens[0], err)
		}

		if !time.Time(since).IsZero() &&  t.Before(time.Time(since)) {
			return
		}
		if !time.Time(until).IsZero() && t.After(time.Time(until)) {
			return
		}
		fmt.Fprintln(out, line)
		idx++

	}).Stdout()
}
