package commitsRepositoryModules

import (
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
)

var CommitsRepositoryModules = &commitsRepositoryModules{}

type commitsRepositoryModules struct{}

func (c *commitsRepositoryModules) GetCommits() []*dto.CommitResponse {
	rows, err := db.DB.Query("SELECT commit_id, author, email, commit_date, message FROM cms.commits")
	if err != nil {
		return nil
	}
	var commitLogs []*dto.CommitResponse

	for rows.Next() {
		var commit dto.CommitResponse
		err := rows.Scan(&commit.CommitID, &commit.Author, &commit.Email, &commit.Date, &commit.Message)
		if err != nil {
			continue
		} else {
			commitLogs = append(commitLogs, &commit)
		}

	}
	return commitLogs
}

func (c *commitsRepositoryModules) GetCommitsCount() (int, error) {
	query := "SELECT COUNT(*) FROM cms.commits"
	rows, err := db.DB.Query(query)
	if err != nil {
		return 0, err
	}
	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}
