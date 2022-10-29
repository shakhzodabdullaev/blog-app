package storage

import "bootcamp/article/models"

type StorageI interface {
	Article() ArticleRepoI
	Author() AuthorRepoI
}

type ArticleRepoI interface {
	Create(entity models.ArticleCreateModel) (err error)
	GetList(query models.Query) (resp []models.ArticleListItem, err error)
	GetByID(ID int) (resp models.Article, err error)
	Update(entity models.ArticleUpdateModel) (effectedRowsNum int, err error)
	Delete(ID int) (effectedRowsNum int, err error)
}

type AuthorRepoI interface {
	Create(entity models.Person) (err error)
	GetList(query models.Query) (resp []models.Person, err error)
	GetByID(ID int) (resp models.Person, err error)
	Update(entity models.Person) (effectedRowsNum int, err error)
	Delete(ID int) (effectedRowsNum int, err error)
}
