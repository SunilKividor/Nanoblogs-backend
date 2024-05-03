package postgresql

import (
	"context"
	"database/sql"
	"log"

	"github.com/SunilKividor/internal/repository/neo4j"
	"github.com/lib/pq"
)

func AddCategory(categories pq.StringArray) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	for _, category := range categories {
		smt := `INSERT INTO Categories(name) VALUES($1) ON CONFLICT(name) DO NOTHING RETURNING name`
		var cat string
		err = tx.QueryRow(smt, category).Scan(&cat)
		if err != nil {
			if err == sql.ErrNoRows {
				//increase frequency in neo4j db
				log.Println("error : category present")
				err := neo4j.IncreaseCategoryFreq(category)
				if err != nil {
					log.Fatalf(err.Error())
				}
			} else {
				tx.Rollback()
				return err
			}
		}
		//save category in neo4j db
		neo4j.CreateNewCategory(category)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
