package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
	"github.com/miekg/dns"          // DNS library for Go
)

const dbPath = "./dns_records.db"

// Open SQLite database
func openDatabase() (*sql.DB, error) {
	return sql.Open("sqlite3", dbPath)
}

// Query DNS record from the SQLite database
func queryDNSRecord(db *sql.DB, name, qtype string) (string, error) {
	var value string
	query := "SELECT value FROM records WHERE name = ? AND type = ?"
	err := db.QueryRow(query, name, qtype).Scan(&value)
	return value, err
}

// Handle incoming DNS queries
func handleDNSQuery(w dns.ResponseWriter, r *dns.Msg) {
	db, err := openDatabase()
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return
	}
	defer db.Close()

	msg := new(dns.Msg)
	msg.SetReply(r)

	for _, q := range r.Question {
		name := strings.TrimSuffix(q.Name, ".") // Remove trailing dot
		qtype := dns.TypeToString[q.Qtype]

		record, err := queryDNSRecord(db, name, qtype)
		if err != nil {
			log.Printf("No record found for %s %s", name, qtype)
			continue
		}

		rr, err := dns.NewRR(fmt.Sprintf("%s %d IN %s %s", q.Name, 3600, qtype, record))
		if err != nil {
			log.Printf("Failed to create DNS response: %v", err)
			continue
		}
		msg.Answer = append(msg.Answer, rr)
	}

	w.WriteMsg(msg)
}

func main() {
	// Set up DNS server
	dns.HandleFunc(".", handleDNSQuery)

	server := &dns.Server{Addr: ":8053", Net: "udp"}
	log.Printf("Starting DNS server on port 8053...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
