package db_game

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const createMails = `-- name: CreateMails :execresult
INSERT INTO mail (account_user_id, sender, ` + "`" + `type` + "`" + `, ` + "`" + `status` + "`" + `, delete_all, attachment, title, message, expired_at) VALUES %s
`

func (q *Queries) CreateMails(ctx context.Context, mails []CreateMailParams) (sql.Result, error) {
	params := []string{}
	for _, mail := range mails {
		attachment := "NULL"
		if mail.Attachment.Valid {
			attachment = fmt.Sprintf("'%s'", mail.Attachment.String)
		}
		message := "NULL"
		if mail.Message.Valid {
			message = fmt.Sprintf("'%s'", mail.Message.String)
		}
		expiredAt := "NULL"
		if mail.ExpiredAt.Valid {
			expiredAt = fmt.Sprintf("'%s'", mail.ExpiredAt.Time.Format("2006-01-02 15:04:05"))
		}
		params = append(params, fmt.Sprintf("(%d, '%s', %d, %d, %d, %s, '%s', %s, %s)", mail.AccountUserID, mail.Sender, mail.Type, mail.Status, mail.DeleteAll, attachment, mail.Title, message, expiredAt))
	}
	query := fmt.Sprintf(createMails, strings.Join(params, ","))

	return q.db.ExecContext(ctx, query)
}

const updateMailsStatus = `-- name: UpdateMailsStatus :execresult
UPDATE mail SET status = %d WHERE id IN (%s)
`

func (q *Queries) UpdateMailsStatus(ctx context.Context, status int32, mailsId []int32) (sql.Result, error) {
	ids := []string{}
	for _, id := range mailsId {
		ids = append(ids, fmt.Sprintf("%d", id))
	}

	query := fmt.Sprintf(updateMailsStatus, status, strings.Join(ids, ","))

	return q.db.ExecContext(ctx, query)
}
