package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Paste struct {
	ID      string    `json:"id"`
	Uid     string    `json:"uid"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Expires time.Time `json:"expires"`
}

func NewPaste(
	uid, title, content string, expires int,
) *Paste {
	return &Paste{
		Uid:     uid,
		Title:   title,
		Content: content,
		Created: time.Now(),
		Expires: time.Now().Add(time.Duration(expires) * 24 * time.Hour),
	}
}

type Database struct {
	rdb *redis.Client
	db  *sql.DB
}

func NewDB(rdb *redis.Client) *Database {
	return &Database{
		rdb: rdb,
		db:  nil,
	}
}

func (d *Database) Insert(
	uid string,
	paste *Paste,
) (int, error) {
	// stmt := `INSERT INTO pastes (title, content, created, expires)
	// VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIME-STAMP(), INTERVAL ? DAY))
	// `
	// result, err := p.Db.Exec(stmt, title, content, expires)
	// if err != nil {
	// 	return 0, err
	// }
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return 0, err
	// }
	// return int(id), nil
	ctx := context.Background()
	b, err := json.Marshal(paste)
	if err != nil {
		return 0, err
	}
	status := d.rdb.Set(ctx, uid, string(b), paste.Expires.Sub(time.Time{}))
	if status.Err() != nil {
		return 0, status.Err()
	}
	return 0, nil
}

func (d *Database) Get(
	uid string,
) (*Paste, error) {
	ctx := context.Background()
	status := d.rdb.Get(ctx, uid)
	if status.Err() != nil {
		return nil, status.Err()
	}
	return nil, nil
}

func (d *Database) Latest() ([]*Paste, error) {
	return nil, nil
}
