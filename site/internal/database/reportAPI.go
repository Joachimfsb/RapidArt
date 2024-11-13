package database

import (
	"database/sql"
	"errors"
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

func GetCountReports(postId int) (int, error) {
	var count = 0

	err := db.QueryRow("SELECT COUNT(PostId) FROM Report WHERE PostId = ?", postId).Scan(&count)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return count, nil
}

func GetAllReportsForPost(postId int) ([]models.Report, error) {
	var reports []models.Report

	rows, err := db.Query("SELECT * FROM Report WHERE PostId = ?", postId)
	if err != nil {
		log.Println(err)
		return []models.Report{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var report models.Report
		err = rows.Scan(&report.ReportId, &report.UserId, &report.PostId, &report.Message, &report.CreationDateTime)
		if err != nil {
			log.Println(err)
			return []models.Report{}, err
		}

		report.CreationDateTime = report.CreationDateTime.Local()
		reports = append(reports, report)
	}
	if errors.Is(err, sql.ErrNoRows) {
		return []models.Report{}, fmt.Errorf("couldnt find post")
	}

	if err != nil {
		log.Println(err)
		return []models.Report{}, err
	}

	return reports, nil
}
