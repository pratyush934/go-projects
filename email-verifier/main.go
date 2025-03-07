package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("domain, hasMx, hasSPF, sprRecord, hasDMARC, dmarcRecord")

	fmt.Println()
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("There exist an Error while taking input ", err)
	}
}

func checkDomain(domain string) {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mx, err := net.LookupMX(domain)

	if err != nil {
		log.Println("there exist an error in LookUpMx", err)
	}

	if len(mx) > 0 {
		hasMx = true
	}

	txt, err := net.LookupTXT(domain)

	lengthOfXXX := len(txt)
	fmt.Println(" This is the length of xxx ", lengthOfXXX)

	if err != nil {
		log.Println("there exist in error in LookUpTXT", err)
	}

	for _, value := range txt {
		if strings.HasPrefix(value, "v=spf1") {
			hasSPF = true
			spfRecord = value
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Println("Error exist in dmarcRecords ", err)
	}

	lengthOfdmarcRecords := len(dmarcRecords)
	fmt.Println("this is the length of xx part 2", lengthOfdmarcRecords)

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Println(domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
