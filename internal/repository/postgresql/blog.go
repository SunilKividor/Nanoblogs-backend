package postgresql

import (
	"context"
	"database/sql"
	"log"

	"github.com/SunilKividor/internal/models"

	"github.com/google/uuid"
)

func PostBlogQuery(blog models.Blog) error {
	smt := `INSERT INTO blogs(userid,title,content,category,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6)`
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	_, err = tx.Exec(smt, blog.UserId, blog.Title, blog.Content, blog.Category, blog.Created_At, blog.Updated_At)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = AddCategory(blog.Category)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func UpdateBlogQuery(blog models.UpdateBlogDBModel) error {
	smt := `UPDATE blogs SET title = $1,content = $2,category = $3 WHERE id = $4`
	_, err := db.Exec(smt, blog.Title, blog.Content, blog.Category, blog.BlogId)
	return err
}

func GetAllUserBlogsQuery(user_id uuid.UUID) ([]models.GetBlogResBody, error) {
	var blogs []models.GetBlogResBody
	smt := `SELECT id,title,content,category,created_at,updated_at FROM blogs WHERE userid = $1`
	rows, err := db.Query(smt, user_id)
	if err != nil {
		return blogs, err
	}
	defer rows.Close()
	for rows.Next() {
		var blog models.GetBlogResBody
		err := rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Category, &blog.Created_At, &blog.Updated_At)
		if err != nil {
			return blogs, err
		}
		blogs = append(blogs, blog)
	}
	if rows.Err() != nil {
		return blogs, err
	}
	log.Println(blogs)
	return blogs, nil
}

func GetAllBlogsQuery() ([]models.GetBlogResBody, error) {
	var blogs []models.GetBlogResBody
	smt := `SELECT id,title,content,category,created_at,updated_at FROM blogs`
	rows, err := db.Query(smt)
	if err != nil {
		return blogs, err
	}
	defer rows.Close()
	for rows.Next() {
		var blog models.GetBlogResBody
		err := rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Category, &blog.Created_At, &blog.Updated_At)
		if err != nil {
			return blogs, err
		}
		blogs = append(blogs, blog)
	}
	if rows.Err() != nil {
		return blogs, err
	}
	log.Println(blogs)
	return blogs, nil
}

func DeleteBlogQuery(blod_id uuid.UUID, user_id uuid.UUID) error {
	log.Println(blod_id)
	log.Println(user_id)
	smt := "DELETE FROM blogs WHERE id=$1 AND userid=$2"
	_, err := db.Exec(smt, blod_id, user_id)
	return err
}
