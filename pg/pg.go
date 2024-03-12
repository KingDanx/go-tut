package pg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	_ "github.com/lib/pq"
)

// ? This function uses the built in go database/sql
func Test() {
	connStr := "connectionString"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id FROM usrdevice ORDER BY id ASC;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name: %s\n", name)
	}
}

// ? this function is utilizing the pgx lib
func TestPGX() {
	conn, err := pgx.Connect(context.Background(), "connectionString")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var id int
	var name string
	var description pgtype.Text //? A string that could be null
	var enabled bool
	var parentGroup int
	err = conn.QueryRow(context.Background(), "SELECT * FROM usrgroup WHERE id=$1", 4000).Scan(&id, &name, &description, &enabled, &parentGroup)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("id:", id, "name:", name, "enabled:", enabled, "parentgroup:", parentGroup)

	monitorDBConnection(context.Background(), &conn)

}

// ? Goes with the PGX version
func monitorDBConnection(ctx context.Context, dbConn **pgx.Conn) {
	ticker := time.NewTicker(1 * time.Second) // Regular ping interval
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := (*dbConn).Ping(ctx); err != nil {
				fmt.Println("Lost database connection, attempting to reconnect...")
				newConn, reconnErr := reconnect(ctx)
				if reconnErr != nil {
					fmt.Println("Reconnection failed:", reconnErr)
					ticker.Reset(5 * time.Second)
				} else {
					(*dbConn).Close(ctx) // Close the old connection
					*dbConn = newConn    // Update the pointer to use the new connection
					ticker.Reset(1 * time.Second)
				}
			} else {
				fmt.Println("Database connection is healthy")
			}
		case <-ctx.Done():
			return // Exit the function if the context is canceled
		}
	}
}

// ? Goes with the PGX version
func reconnect(ctx context.Context) (*pgx.Conn, error) {
	dbConn, err := pgx.Connect(ctx, "connectionString")
	if err != nil {
		return nil, err
	}
	fmt.Println("Reconnected to the database successfully.")
	return dbConn, nil
}
