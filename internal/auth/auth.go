package auth

import (
	"fmt"
	"github.com/dopaemon/artus/internal/config"
	"github.com/dopaemon/artus/internal/db"
)

func RegisterOrLogin() bool {
	var username, password string
	users, err := db.GetAllUsers()
	if err != nil {
		fmt.Println("Lỗi truy vấn user:", err)
		return false
	}

	if len(users) == 0 {
		fmt.Println("Chưa có tài khoản, vui lòng đăng ký")
		fmt.Print("Nhập username: ")
		fmt.Scan(&username)
		fmt.Print("Nhập password: ")
		fmt.Scan(&password)

		u, err := db.CreateUser(username, password)
		if err != nil {
			fmt.Println("Không tạo được user:", err)
			return false
		}
		fmt.Println("Tạo tài khoản thành công")
		fmt.Println("API Key:", u.APIKey)
		return true
	} else {
		fmt.Println("Đăng nhập")
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)

		if db.Authenticate(username, password) {
			fmt.Println("Đăng nhập thành công")
			config.Login = true
			config.APIKey, _ = db.GetAPIKey()
			return true
		} else {
			fmt.Println("Sai username hoặc password")
			return false
		}
	}
}
