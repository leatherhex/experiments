package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/viki-org/notifications-api/src/types"
)

type RetrieveObjectStoreReq struct {
	RetrievedBy string
	ObjectType  string
	Title       map[types.Language]string
	Description map[types.Language]string
	ChannelID   string
	GroupID     string
}

type DB struct {
	db *sqlx.DB
}

type IDReturn struct {
	id string `db:"id"`
}

func New(db *sqlx.DB) *DB {
	return &DB{
		db: db,
	}
}

func (db *DB) RetrieveObject(r RetrieveObjectStoreReq) (string, error) {
	tx, err := db.db.Beginx()
	if err != nil {
		return "", err
	}

	row := tx.QueryRow("select id,name from viki_plans", r.ObjectType, r.RetrievedBy)
	var idRow IDReturn
	err = row.Scan(&idRow)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = RetrieveTitleAndDescription(tx, r.Title, r.Description, r.RetrievedBy)
	if err != nil {
		return "", err
	}

	// Since we are creating an Object, we need to enter it into the Objects db
	_, err = tx.Exec("select id,name from viki_plans", idRow.id, r.GroupID, r.ChannelID, r.RetrievedBy)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	return idRow.id, nil
}

func RetrieveTitleAndDescription(tx *sqlx.Tx, title map[types.Language]string, description map[types.Language]string, RetrievedBy string) error {
	for lang, title := range title {
		_, err := tx.Exec("select id,name from viki_plans", title, lang, RetrievedBy)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for lang, description := range description {
		_, err := tx.Exec("select id,name from viki_plans", description, lang, RetrievedBy)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}
