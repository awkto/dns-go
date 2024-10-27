# Step 1: Create the SQLite Database and DNS Records Table
This part involves setting up a local SQLite database to store the DNS records that the Go app will serve. Here’s what you need to do:

## Step 1.1: Install SQLite (if not already installed)
- Linux/macOS:
You can install **SQLite** via your package manager:

```bash
sudo apt-get install sqlite3        # Ubuntu/Debian
sudo pacman -S sqlite               # Arch Linux
brew install sqlite                 # macOS
```
- Windows:
Download the SQLite tools from the official website:
https://sqlite.org/download.html

## Step 1.2: Create the SQLite Database
Create a new SQLite database file (if it doesn't exist yet). Open your terminal or command prompt and run:

```bash
sqlite3 dns_records.db
```
This command starts the SQLite shell with the database file dns_records.db. If the file doesn’t exist, it will be created in the current directory.

## Step 1.3: Create the Records Table
Inside the SQLite shell (you'll see a prompt like sqlite>), run the following SQL command to create a table for DNS records:

```sql
CREATE TABLE records (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    value TEXT NOT NULL
);
```
- id: A unique identifier for each record (auto-incremented).
- name: The domain name (e.g., example.com).
- type: The DNS record type (e.g., A, MX, TXT).
- value: The corresponding value for the record (e.g., IP address for A record).


## Step 1.4: Insert Sample DNS Records
While still inside the SQLite shell, insert some example DNS records for testing:

```sql
INSERT INTO records (name, type, value) VALUES ('example.com', 'A', '93.184.216.34');
INSERT INTO records (name, type, value) VALUES ('example.com', 'MX', 'mail.example.com');
INSERT INTO records (name, type, value) VALUES ('example.com', 'TXT', 'v=spf1 include:example.com ~all');
```
- A Record: Maps the domain example.com to the IP 93.184.216.34.
- MX Record: Specifies the mail exchange server mail.example.com.
- TXT Record: Provides a sample SPF entry for the domain.

## Step 1.5: Verify the Records
To confirm the records were added, run the following query:

```sql
SELECT * FROM records;
You should see output like this:
```
```less
1|example.com|A|93.184.216.34
2|example.com|MX|mail.example.com
3|example.com|TXT|v=spf1 include:example.com ~all
```

## Step 1.6: Exit the SQLite Shell
Once you’re done inserting records, type:

```sql
.exit
```

This will exit the SQLite shell and save the changes to the **dns_records.db** file.