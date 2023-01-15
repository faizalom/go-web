package models

import (
	"context"

	"github.com/faizalom/go-web/config"

	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type AccessToken struct {
	ID        int64
	UserID    int64
	Token     string
	ExpiredAt time.Time
}

func GenerateAccessToken(userId int64) (string, error) {
	db := Conn()

	id := uuid.New()
	token := id.String()

	query := "INSERT INTO access_token(user_id, token, expired_at) VALUES (?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer stmt.Close()

	expiredAt := time.Now().Add(time.Minute * time.Duration(config.SessionLifetime)).Format("2006-01-02 15:04:05")
	_, err = stmt.ExecContext(ctx, userId, token, expiredAt)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return token, err
}

func GetAuthUser(token string) (User, error) {
	user := User{}
	id, userId, err := VerifyAccessToken(token)
	if userId == 0 {
		return user, err
	}
	err = UpdateAccessToken(id)
	if err != nil {
		return user, err
	}
	return GetUserById(userId)
}

func VerifyAccessToken(token string) (int64, int64, error) {
	var id int64
	var userID int64
	expiredAt := time.Now().Format("2006-01-02 15:04:05")
	err := Conn().QueryRow("SELECT id, user_id FROM `access_token` WHERE `token` = ? AND `expired_at` >= ?", token, expiredAt).Scan(&id, &userID)
	if err == sql.ErrNoRows {
		return 0, 0, nil
	} else if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	return id, userID, nil
}

func UpdateAccessToken(id int64) error {
	db := Conn()

	query := "UPDATE `access_token` set expired_at = ? WHERE id = ?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	expiredAt := time.Now().Add(time.Minute * time.Duration(config.SessionLifetime)).Format("2006-01-02 15:04:05")
	_, err = stmt.ExecContext(ctx, expiredAt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func DeleteAccessToken(token string) {
	db := Conn()

	query := "DELETE from `access_token` WHERE token = ?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, token)
	if err != nil {
		log.Println(err)
	}
}
