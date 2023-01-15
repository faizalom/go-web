package models

import (
	"context"
	"database/sql"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

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
	cnt, err := Count(Conn(), "SELECT COUNT(*) FROM `users` WHERE `email` = ?", email)
	if cnt == 0 {
		return false, err
	}
	return true, err
}

func Login(email string, password string) (int64, error) {
	var id int64
	var passwordHash string
	err := Conn().QueryRow("SELECT id, password FROM `users` WHERE `google_Id` IS NULL AND `email` = ?", email).Scan(&id, &passwordHash)
	if err == sql.ErrNoRows {
		return 0, err
	} else if err != nil {
		log.Println(err)
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, nil
	} else if err != nil {
		log.Println(err)
		return 0, err
	}
	return id, err
}

func GetUserById[T Number](id T) (User, error) {
	var user User
	err := Conn().QueryRow("SELECT id, first_name, last_name, email FROM `users` WHERE `id` = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
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
	err := Conn().QueryRow("SELECT id, first_name, last_name, email FROM `users` WHERE `id` = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
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
