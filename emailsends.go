package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"net/smtp"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Email string `json:"email"`
	Lang  string `json:"lang"`
}

func sendMailSimple(msg string, email string) {
	auth := smtp.PlainAuth(
		"",
		"mensch.the98@gmail.com", // почта отправитель
		"",                       // 16-ричный пароль от почты отправителя
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"mensch.the98@gmail.com", // почта отправитель
		[]string{email},          // почта получатель
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var msg string
	fmt.Println("Введите текст для отправки")

	//чтение строки с пробелами
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg = scanner.Text()

	//открытие базы данных
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	// загрузка данных по одному
	// insert, err := db.Query("INSERT INTO `users` (`email`, `lang`) VALUES('vadimprog666@gmail.com', 'ru')")
	// if err != nil {
	// 	panic(err)
	// }
	// defer insert.Close()

	// выборка столбцов из таблицы
	res, err := db.Query("SELECT `email`, `lang` FROM `users`")
	if err != nil {
		panic(err)
	}

	// чтение переменных из таблицы
	for res.Next() {
		var user User
		err = res.Scan(&user.Email, &user.Lang)
		if err != nil {
			panic(err)
		}

		email := user.Email

		sendMailSimple(msg, email)
		// fmt.Println(fmt.Sprintf("User: %s and lang %s", user.Email, user.Lang))
	}

	defer db.Close()

}
