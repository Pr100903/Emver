package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type DomainCheckResponse struct {
	Domain      string `json:"domain"`
	HasMX       bool   `json:"hasMX"`
	HasSPF      bool   `json:"hasSPF"`
	SPFRecord   string `json:"spfRecord,omitempty"`
	HasDMARC    bool   `json:"hasDMARC"`
	DMARCRecord string `json:"dmarcRecord,omitempty"`
}

func main() {
	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Add this line

		domain := r.URL.Query().Get("domain")
		if domain == "" {
			http.Error(w, "Domain query parameter is required", http.StatusBadRequest)
			return
		}

		response := checker(domain)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checker(domain string) DomainCheckResponse {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err == nil && len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err == nil {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
				break
			}
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err == nil {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record
				break
			}
		}
	}

	return DomainCheckResponse{
		Domain:      domain,
		HasMX:       hasMX,
		HasSPF:      hasSPF,
		SPFRecord:   spfRecord,
		HasDMARC:    hasDMARC,
		DMARCRecord: dmarcRecord,
	}
}
