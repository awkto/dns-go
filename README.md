# Getting things up and Running

## Create the SQLite database:

```bash
sqlite3 dns_records.db < database-setup.sql
```

## Run the Go application

```bash
go run main.go
```

## Test the DNS Server
Use dig to test the DNS server locally:

```bash
dig @localhost -p 8053 example.com A
dig @localhost -p 8053 example.com MX
dig @localhost -p 8053 example.com TXT
```

### Explanation
- **SQLite Backend** : The DNS server queries SQLite for records based on the name and type.
- **DNS Handling** : The Go server listens on port 8053 and uses the miekg/dns library to respond to DNS queries.
- **Testing** : You can use dig to verify if the DNS server is returning correct records.
