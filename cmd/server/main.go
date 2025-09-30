package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dopaemon/artus/internal/db"
)

func main() {
	db.InitDB()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	u, _ := reader.ReadString('\n')
	u = strings.TrimSpace(u)

	fmt.Print("Password: ")
	p, _ := reader.ReadString('\n')
	p = strings.TrimSpace(p)

	user, err := db.CreateUser(u, p)
	if err != nil {
		log.Fatalf("CreateUser error: %v", err)
	}

	fmt.Printf("User created: %s\nAPIKey: %s\n", user.Username, user.APIKey)
}
