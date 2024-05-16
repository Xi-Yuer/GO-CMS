package logsRepositoryModules

import (
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
)

var LogsRepository = logsRepository{}

type logsRepository struct{}

func (l *logsRepository) CreateLogRecord(params *dto.CreateLogRecordRequest) error {
	query := "INSERT INTO logs (id, user_id, user_account, ip, method, path, status, duration) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.DB.Exec(query, params.ID, params.UserID, params.UserName, params.UserIP, params.RequestMethod, params.RequestPath, params.RequestStatus, params.RequestDuration)
	return err
}

func (l *logsRepository) GetLogRecords(params *dto.Page) (*dto.HasTotalResponseData, error) {
	count := "SELECT COUNT(*) FROM logs"
	var total int
	r, err := db.DB.Query(count)
	if err != nil {
		return nil, err
	}
	for r.Next() {
		err := r.Scan(&total)
		if err != nil {
			return nil, err
		}
	}
	query := "SELECT id, user_id, user_account, ip, method, path, status, duration, create_time FROM logs ORDER BY create_time DESC LIMIT ?, ?"
	var logs []*dto.GetLogRecordResponse
	rows, err := db.DB.Query(query, params.Offset, params.Limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var log dto.GetLogRecordResponse
		err := rows.Scan(&log.ID, &log.UserID, &log.UserName, &log.UserIP, &log.RequestMethod, &log.RequestPath, &log.RequestStatus, &log.RequestDuration, &log.CreateTime)
		if err != nil {
			return nil, err
		}
		logs = append(logs, &log)
	}
	return &dto.HasTotalResponseData{
		List:  logs,
		Total: total,
	}, nil
}
