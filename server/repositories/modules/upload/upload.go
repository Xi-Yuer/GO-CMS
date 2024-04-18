package uploadRepositorysModules

import (
	"database/sql"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
	"time"
)

var UploadRepository = &uploadRepository{}

type uploadRepository struct{}

func (u *uploadRepository) CreateRecord(user string, params *dto.UploadFinishRequest) error {
	query := "INSERT INTO file (id, upload_user, file_name, file_size,upload_time) VALUES (?, ?, ?, ?, ?)"
	if _, err := db.DB.Exec(query, params.Identifier, user, params.FileName, params.FileSize, time.Now()); err != nil {
		return err
	}
	return nil
}

func (u *uploadRepository) DeleteRecord(id string) error {
	query := "DELETE FROM file WHERE id = ?"

	if _, err := db.DB.Exec(query, id); err != nil {
		return err
	}
	return nil
}

func (u *uploadRepository) GetRecords(params *dto.Page) ([]dto.UploadRecordResponse, error) {
	query := "SELECT id, upload_user, file_name, file_size,upload_time FROM file LIMIT ? OFFSET ?"
	rows, err := db.DB.Query(query, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var records []dto.UploadRecordResponse
	for rows.Next() {
		var record dto.UploadRecordResponse
		if err := rows.Scan(&record.FileID, &record.UploadUser, &record.FileName, &record.FileSize, &record.UploadTime); err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func (u *uploadRepository) GetFileInfo(id string) (*dto.UploadRecordResponse, error) {
	query := "SELECT id, upload_user, file_name, file_size, upload_time FROM file WHERE id = ?"
	rows, err := db.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var record dto.UploadRecordResponse
	if rows.Next() {
		if err := rows.Scan(&record.FileID, &record.UploadUser, &record.FileName, &record.FileSize, &record.UploadTime); err != nil {
			return nil, err
		}
	}
	return &record, nil
}
