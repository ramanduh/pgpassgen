package main

import (
	"fmt"
	"os"
)

const (
	filename = ".pgpass"
)

func main() {
	var (
		host     string = "*"
		port     int    = 5432
		database string = "*"
		username string = "*"
		password string
	)

	fmt.Print("Host [*]: ")
	fmt.Scanln(&host)
	fmt.Print("Port [5432]: ")
	fmt.Scanln(&port)
	fmt.Print("Database [*]: ")
	fmt.Scanln(&database)
	fmt.Print("Username [*]: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	f, err := os.Create(fmt.Sprintf("%s/%s", dirname, filename))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	mode := os.FileMode(0600)
	f.Chmod(mode)

	output := fmt.Sprintf("%s:%d:%s:%s:%s\n", host, port, database, username, password)
	_, err = f.WriteString(output)
	if err != nil {
		panic(err)
	}

	hint := fmt.Sprintf("psql -h %s -U %s -d %s", host, username, database)
	fmt.Println(fmt.Sprintf("\npgpass created in: %s/%s", dirname, filename))
	fmt.Println(hint)
	// psql -U postgres -h blabla -d dbname
}
