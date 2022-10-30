package postgres

import (
	"blog/app/models"
	"blog/app/storage"

	"github.com/jmoiron/sqlx"
)

type authorRepo struct {
	db *sqlx.DB
}

func NewAuthorRepo(db *sqlx.DB) storage.AuthorRepoI {
	return authorRepo{
		db: db,
	}
}

func (r authorRepo) Create(entity models.Person) (err error) {

	return
}

func (r authorRepo) GetList(query models.Query) (resp []models.Person, err error) {

	return
}

func (r authorRepo) GetByID(ID int) (resp models.Article, err error) {

	return
}

func (r authorRepo) Update(entity models.ArticleUpdateModel) (effectedRowsNum int, err error) {

	return
}

func (r authorRepo) Delete(ID int) (effectedRowsNum int, err error) {

	return
}
