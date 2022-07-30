package database

import (
	"fmt"
)

func SelectMessage() []string {

	rows, err := DB.Query("SELECT text FROM history_msg")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	history_msg := []string{}
	for rows.Next() {
		msgF := ""
		rows.Scan(&msgF)
		history_msg = append(history_msg, msgF)
	}
	return history_msg

}

func InsertMessage(text string) error {
	_, err := DB.Query("INSERT INTO history_msg (text)VALUES ($1)", text)
	if err != nil {
		return err
	}
	return nil
}
