package db_test

import (
	"os"
	"strings"
	"testing"

	"github.com/DaveSaah/voting-system/db"
	"github.com/DaveSaah/voting-system/db/models"
)

func TestDatabaseConnection(t *testing.T) {
	conn, err := db.Init()
	if err != nil {
		t.Fatalf("Cannot create the db connection: %s", err)
	}

	defer conn.Close()

	if err := conn.Ping(); err != nil {
		t.Fatalf("Ping Error: %s", err)
	}

	t.Log("Database connected!")
}

func TestInsertData(t *testing.T) {
	conn, err := db.Init()
	if err != nil {
		t.Fatalf("Cannot create the db connection: %s", err)
	}

	defer conn.Close()

	_, err = conn.Exec("insert into dummy values(1, 'First entry')")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAccessData(t *testing.T) {
	var test_dummy models.Dummy

	conn, err := db.Init()
	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close()

	row := conn.QueryRow("select description from dummy")
	if err := row.Scan(&test_dummy.Description); err != nil {
		t.Fatal(err)
	}
}

func TestSanitizeDB(t *testing.T) {
	// load sql file
	file, err := os.ReadFile("./init.sql")
	if err != nil {
		t.Fatal(err)
	}

	// create db connection
	conn, err := db.Init()
	if err != nil {
		t.Fatalf("Cannot create the db connection: %s", err)
	}

	defer conn.Close()

	tx, err := conn.Begin() // start db transaction
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = tx.Rollback() // aborts transaction
	}()

	// run each line in file
	for _, q := range strings.Split(string(file), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}

		if _, err := tx.Exec(q); err != nil {
			t.Fatal(err)
		}
	}

	if err = tx.Commit(); err != nil {
		t.Fatal(err)
	}
}
