package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/pete911/certinfo/pkg/cert"
)

func printCertificates(certificateLocations []cert.CertificateLocation, days int, verbose bool) error {
	var err error = nil
	for _, certificateLocation := range certificateLocations {
		var newErr error
		expiresInDays := 0.0
		if certificateLocation.Error != nil {
			newErr = certificateLocation.Error
		} else {
			expiresInDays, newErr = checkCertificate(certificateLocation.Certificates[0])
		}
		if newErr != nil {
			fmt.Printf("ERROR [%s] %s\n", certificateLocation.Name(), newErr.Error())
		} else if expiresInDays < float64(days) {
			newErr = errors.New("A certificate will expire soon")
			fmt.Printf("ERROR [%s] expires in %d days\n", certificateLocation.Name(), int(expiresInDays))
		} else if verbose {
			fmt.Printf("OK [%s] expires in %d days\n", certificateLocation.Name(), int(expiresInDays))
		}
		if err == nil && newErr != nil {
			err = newErr
		}
	}
	return err
}

func checkCertificate(certificate cert.Certificate) (float64, error) {
	if certificate.Error() != nil {
		return 0, certificate.Error()
	}
	expiresInDays := certificate.NotAfter().Sub(time.Now()).Hours() / 24
	return expiresInDays, nil
}
