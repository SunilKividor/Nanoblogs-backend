package postgresql

import (
	"log"

	"github.com/SunilKividor/internal/models"
)

func PostBlogQuery(blog models.Blog) error {
	smt := `INSERT INTO blogs(userid,title,content,category,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6) RETURNING id`
	_, err := db.Exec(smt, blog.UserId, blog.Title, blog.Content, blog.Category, blog.Created_At, blog.Updated_At)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBlogQuery(blog models.UpdateBlogReqModel) error {
	smt := `UPDATE blogs SET title = $1,content = $2,category = $3 WHERE id = $4`
	_, err := db.Exec(smt, blog.Title, blog.Content, blog.Category, blog.BlogId)
	return err
}

func GetAllUserBlogsQuery(user_id models.GetBlogsReqModel) ([]models.Blog, error) {
	var blogs []models.Blog
	smt := `SELECT * FROM blogs WHERE userid = $1`
	rows, err := db.Query(smt, user_id.UserId)
	if err != nil {
		return blogs, err
	}
	defer rows.Close()
	for rows.Next() {
		var blog models.Blog
		err := rows.Scan(&blog.Id, &blog.UserId, &blog.Title, &blog.Content, &blog.Category, &blog.Created_At, &blog.Updated_At)
		if err != nil {
			return blogs, err
		}
		log.Println(blog)
		blogs = append(blogs, blog)
	}
	if rows.Err() != nil {
		return blogs, err
	}
	log.Println(blogs)
	return blogs, nil
}
