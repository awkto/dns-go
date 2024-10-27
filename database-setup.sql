-- Create the SQLite database and table
CREATE TABLE records (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    value TEXT NOT NULL
);

-- Insert some test records
INSERT INTO records (name, type, value) VALUES ('example.com', 'A', '93.184.216.34');
INSERT INTO records (name, type, value) VALUES ('example.com', 'MX', 'mail.example.com');
INSERT INTO records (name, type, value) VALUES ('example.com', 'TXT', 'v=spf1 include:example.com ~all');
