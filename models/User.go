package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64  `json:"id"`
	GoogleID  string `json:"-"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

func InsertUser(user User) (int64, error) {
	db := Conn()

	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, user.Email, user.Password)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return result.LastInsertId()
}

func InsertUserWithDetails(user User) (int64, error) {
	db := Conn()

	query := "INSERT INTO users(email, password, google_id, first_name, last_name) VALUES (?, ?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, user.Email, user.Password, user.GoogleID, user.FirstName, user.LastName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return result.LastInsertId()
}

func IsEmailExists(email string) (bool, error) {
	cnt, err := Count(Conn(), "SELECT COUNT(*) FROM users WHERE email = ?", email)
	if cnt == 0 {
		return false, err
	}
	return true, err
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := Conn().QueryRow("SELECT id, first_name, last_name, email, password FROM users WHERE google_id IS NULL AND email = ?", email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return user, err
	} else if err != nil {
		log.Println(err)
		return user, err
	}
	return user, err
}

func Login(email string, password string) (User, string) {
	user, err := GetUserByEmail(email)
	if err == sql.ErrNoRows {
		return user, "Username or password is incorrect"
	} else if err != nil {
		return user, "Internal Server Error"
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return user, "Username or password is incorrect"
	} else if err != nil {
		log.Println(err)
		return user, "Internal Server Error"
	}
	return user, ""
}

func GetAuthUser(auth *sessions.Session) (User, error) {
	id := fmt.Sprintf("%d", auth.Values["user_id"])
	var user User
	err := Conn().QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err == sql.ErrNoRows {
		return user, nil
	} else if err != nil {
		log.Println(err)
		return user, err
	}
	return user, err
}

func GetUserByGoogleId(id string) (User, error) {
	var user User
	err := Conn().QueryRow("SELECT id, first_name, last_name, email FROM users WHERE google_id = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err == sql.ErrNoRows {
		return user, nil
	} else if err != nil {
		log.Println(err)
		return user, err
	}
	return user, err
}

func UpdateUser(user User) (int64, error) {
	db := Conn()

	query := "UPDATE users set first_name = ?, last_name = ? WHERE id = ?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, user.FirstName, user.LastName, user.ID)
	if err != nil {
		log.Println(err)
	}
	return result.RowsAffected()
}
