package postgres

import (
	"blog/app/storage"
	"log"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db      *sqlx.DB
	article storage.ArticleRepoI
	author  storage.AuthorRepoI
}

func NewPostgres(psqlConnString string) storage.StorageI {
	db, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		log.Panic(err)
	}

	return &Store{
		db: db,
	}
}

func (s *Store) Article() storage.ArticleRepoI {
	if s.article == nil {
		s.article = &articleRepo{db: s.db}
	}
	return s.article
}

func (s *Store) Author() storage.AuthorRepoI {
	if s.author == nil {
		s.author = &authorRepo{db: s.db}
	}
	return s.author
}
