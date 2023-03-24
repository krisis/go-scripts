package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type xlMetaV2VersionHeader struct {
	VersionID string
	ModTime   time.Time
}

type xlMetaV2ShallowVersion struct {
	Header xlMetaV2VersionHeader
}

type xlMetaV2 struct {
	Versions []xlMetaV2ShallowVersion
}

func main() {
	var filename string
	flag.StringVar(&filename, "file", "", "path to xl-meta files")
	flag.StringVar(&filename, "f", "", "path to xl-meta files")
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.Parse()
	if filename == "" {
		log.Fatal("filename is mandatory")
	}

	r, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", filename, err)
	}

	type xlMetas map[string]xlMetaV2
	var xlmetas xlMetas
	d := json.NewDecoder(r)
	err = d.Decode(&xlmetas)
	if err != nil {
		log.Fatal(err)
	}

	versionsMap := make(map[string]time.Time)
	for _, xlmeta := range xlmetas {
		for _, v := range xlmeta.Versions {
			versionsMap[v.Header.VersionID] = v.Header.ModTime
		}
	}

	type xlVersion struct {
		versionID string
		mtime     time.Time
	}
	type xlVersions []xlVersion

	var allVersions xlVersions
	for vid, mtime := range versionsMap {
		allVersions = append(allVersions, xlVersion{vid, mtime})
	}
	sort.Slice(allVersions, func(i, j int) bool {
		return allVersions[i].mtime.Before(allVersions[j].mtime)
	})
	vlabels := make(map[string]string, len(allVersions)) // e.g map[versionID] -> v1
	for i, v := range allVersions {
		vlabels[v.versionID] = fmt.Sprintf("v%d", i+1)
	}

	type driveXL struct {
		drive  string
		labels []string
	}
	var xls []driveXL
	for d, xlmeta := range xlmetas {
		var labels []string
		for _, v := range xlmeta.Versions {
			labels = append(labels, vlabels[v.Header.VersionID])
		}
		xls = append(xls, driveXL{d, labels})
	}
	sort.Slice(xls, func(i, j int) bool {
		return xls[i].drive < xls[j].drive
	})
	if verbose {
		fmt.Println("Versions")
		for i, v := range allVersions {
			fmt.Printf("v%d: %s\n", i+1, v.versionID)
		}
		fmt.Println("")
		fmt.Println("Drives")
		var i int
		for drive := range xlmetas {
			fmt.Printf("d%d: %s\n", i+1, drive)
			i++
		}
		fmt.Println("")
	}
	for i, xl := range xls {
		fmt.Printf("d%d: ", i+1)
		for _, l := range xl.labels {
			fmt.Printf("%s ", l)
		}
		fmt.Println()
	}
}
