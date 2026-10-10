package main

import (
	"errors"
	"flag"
	"fmt"
)

type Flags struct {
	Args	[]string
	Days	int
	Verbose bool
	Version bool
}

func ParseFlags() (Flags, error) {
	// Change the default "Usage" massage.
	flag.CommandLine.Usage = func() {
		fmt.Println("Usage:\n    certexpire [flags] host1 [host2:port2 ...]\nwith [flags] in:")
		flag.PrintDefaults()
	}

	var flags Flags
	flag.BoolVar(&flags.Verbose, "verbose", false, "Display info even when no problem arises")
	flag.BoolVar(&flags.Version, "version", false, "Display the version and quit")
	flag.IntVar(&flags.Days, "days", 10, "Warn if cert expires soon, in days of 24h")
	flag.Parse()
	flags.Args = flag.Args()

	if !flags.Version && len(flags.Args) == 0 {
		return Flags{}, errors.New("At least one host (optionally with :port) is required")
	}
	if flags.Days < 1 {
		return Flags{}, errors.New("The minimal expiration value is 1 (day)")
	}

	return flags, nil
}
