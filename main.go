package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/pete911/certinfo/pkg/cert"
)

var version = "0.1"

func main() {
	flags, err := ParseFlags()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if flags.Version {
		fmt.Printf("certexpire %s\n", version)
		os.Exit(0)
	}

	certificatesFiles := loadFromArgs(flags.Args)
	err = printCertificates(certificatesFiles, flags.Days, flags.Verbose)
	if err != nil {
		os.Exit(2)
	}
}

func loadFromArgs(args []string) cert.CertificateLocations {
	out := make(chan cert.CertificateLocation)
	go func() {
		var wg sync.WaitGroup
		for _, arg := range args {
			wg.Go(func() {
				addr := toAddress(arg)
				if addr == "" {
					fmt.Printf("%s is not a valid hostname or hostname:port", arg)
				} else {
					out <- cert.LoadCertificatesFromNetwork(addr, "", true)
				}
			})
		}
		wg.Wait()
		close(out)
	}()

	// load certificates from the channel
	certsByArgs := make(map[string]cert.CertificateLocation)
	for location := range out {
		certsByArgs[location.Path] = location
	}

	// sort certificates by input arguments
	var certsSortedByArgs cert.CertificateLocations
	for _, arg := range args {
		arg = toAddress(arg)
		certsSortedByArgs = append(certsSortedByArgs, certsByArgs[arg])
	}
	return certsSortedByArgs
}

func toAddress(arg string) string {
	parts := strings.Split(arg, ":")
	if len(parts) == 1 {
		return arg + ":443"
	}
	if len(parts) != 2 {
		return ""
	}
	if _, err := strconv.Atoi(parts[1]); err != nil {
		return ""
	}
	return arg
}
