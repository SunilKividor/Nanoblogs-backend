package postgresql

import "github.com/SunilKividor/internal/models"

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
