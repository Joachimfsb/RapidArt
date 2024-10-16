package database

import (
	"fmt"
	"log"
	"rapidart/internal/models"
	"time"
)

func NewReport(report models.Report) error {
	report.CreationDateTime = time.Now()

	sqlInsert := `
		INSERT INTO Report (
		                  UserId,
		                  PostId,
		                  Message,
		                  CreationDateTime
		) VALUES (?, ?, ?, ?);`

	_, err := db.Exec(sqlInsert,
		report.UserId,
		report.PostId,
		report.Message,
		report.CreationDateTime,
	)
	if err != nil {
		log.Println("Error: ", err)
		fmt.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}
	return nil
}
