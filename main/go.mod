module test/db

go 1.22.1

replace test/pg => ../pg

replace test/tcp => ../tcp

require (
	test/pg v0.0.0-00010101000000-000000000000
	test/tcp v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/lib/pq v1.10.9 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)