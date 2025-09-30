package main

import (
	"fmt"
	"log"

	"github.com/dopaemon/artus/internal/config"
	"github.com/dopaemon/artus/internal/db"
)

func main() {
	db.InitDB()

	var username, password string

	users, err := db.GetAllUsers()
	if err != nil {
		log.Fatal("Lỗi truy vấn user:", err)
	}

	if len(users) == 0 {
		fmt.Println("Chưa có tài khoản, vui lòng đăng ký")

		fmt.Print("Nhập username: ")
		fmt.Scan(&username)

		fmt.Print("Nhập password: ")
		fmt.Scan(&password)

		u, err := db.CreateUser(username, password)
		if err != nil {
			log.Fatal("Không tạo được user:", err)
		}

		fmt.Println("Tạo tài khoản thành công")
		fmt.Println("API Key:", u.APIKey)
	} else {
		fmt.Println("Đăng nhập")

		fmt.Print("Username: ")
		fmt.Scan(&username)

		fmt.Print("Password: ")
		fmt.Scan(&password)

		if db.Authenticate(username, password) {
			fmt.Println("Đăng nhập thành công")
			config.Login = true
		} else {
			fmt.Println("Sai username hoặc password")
		}
	}

	if (config.Login) {
		fmt.Println("You Already Login !!!")
		config.APIKey, _ = db.GetAPIKey()
		fmt.Println(config.APIKey)
	}
}
