package postgres

import (
	"github.com/jmoiron/sqlx"

	"blog/app/models"
)

type articleRepo struct {
	db *sqlx.DB
}

func NewArticleRepo(db *sqlx.DB) storage.ArticleRepoI {
	return articleRepo{
		db: db,
	}
}

func (r articleRepo) Create(entity models.ArticleCreatedModel) {
	insertQuery := `INSERT INTO article(
title, body, author_id) VALUE (
	$1,
	$2, 
	$3
)`

	_err = r.db.Exec(insertQuery,
		entity.Title,
	)

}
