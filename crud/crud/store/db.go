package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// type RetrieveObjectStoreReq struct {
// 	RetrievedBy string
// 	ObjectType  string
// 	Title       map[types.Language]string
// 	Description map[types.Language]string
// 	ChannelID   string
// 	GroupID     string
// }

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

type getObjectRequest struct {
	ID string
}

type Object struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

// func txStmt(db *DB, stmt *sql.Stmt) *sql.Stmt {
// 	fmt.Println("db.go: Inside txStmt.")
// 	tx, err := db.db.Beginx()
// 	if tx == nil || err != nil {
// 		return stmt
// 	}
// 	return tx.Stmt(stmt)
// }

func (db *DB) RetrieveObject(r string) (string, error) {
	fmt.Println("db.go: Inside RetrieveObject. r = " + r)
	// tx, err := db.db.Beginx()
	// if tx == nil || err != nil {
	// 	return "", err
	// }
	people := []Object{}
	db.db.Select(&people, "SELECT id,name FROM viki_plans where id='"+r+"'")
	first := people[0]
	return first.Name, nil
}

// func RetrieveTitleAndDescription(tx *sqlx.Tx, title map[types.Language]string, description map[types.Language]string, RetrievedBy string) error {
// 	for lang, title := range title {
// 		_, err := tx.Exec("select id,name from viki_plans", title, lang, RetrievedBy)
// 		if err != nil {
// 			tx.Rollback()
// 			return err
// 		}
// 	}

// 	for lang, description := range description {
// 		_, err := tx.Exec("select id,name from viki_plans", description, lang, RetrievedBy)
// 		if err != nil {
// 			tx.Rollback()
// 			return err
// 		}
// 	}
// 	return nil
// }
