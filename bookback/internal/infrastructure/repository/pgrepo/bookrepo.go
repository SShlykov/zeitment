package pgrepo

import (
	"context"
	entity2 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/postgres"
	"github.com/jackc/pgx/v5"
)

// BookRepo описывает репозиторий для работы с книгами.
//
//go:generate mockgen -destination=../../../../tests/mocks/domain/repository/pgrepo/book_repo_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo BookRepo
type BookRepo interface {
	Repository[entity2.Book]
	GetTOCSectionsFromChapters(ctx context.Context, bookID string) ([]*entity2.Section, error)
	GetTOCSectionsFromPages(ctx context.Context, bookID string) ([]*entity2.Section, error)
}

type bookRepo struct {
	repository[entity2.Book]
}

func NewBookRepository(db postgres.Client) BookRepo {
	return &bookRepo{
		repository: repository[entity2.Book]{
			Name:   "BookRepository",
			entity: entity2.Book{},
			db:     db,
		},
	}
}

func (br *bookRepo) GetTOCSectionsFromChapters(ctx context.Context, bookID string) ([]*entity2.Section, error) {
	query, args, err := br.db.Builder().
		Select("id", "title", "number", "is_public").
		From("chapters").
		Where("book_id "+"= ?", bookID).
		OrderBy("number ASC").
		ToSql()

	if err != nil {
		return nil, err
	}
	q := postgres.Query{Name: br.Name + ".GetTOCSectionsFromChapters", Raw: query}

	var rows pgx.Rows
	rows, err = br.db.DB().QueryRawContextMulti(ctx, q, args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	sections := make([]*entity2.Section, 0)
	for rows.Next() {
		var section entity2.Section
		section.Level = "chapter"
		err = rows.Scan(&section.ID, &section.Title, &section.Order, &section.IsPublic)
		if err != nil {
			return nil, err
		}
		sections = append(sections, &section)
	}

	return sections, nil
}

func (br *bookRepo) GetTOCSectionsFromPages(ctx context.Context, bookID string) ([]*entity2.Section, error) {
	query, args, err := br.db.Builder().
		Select("pages.id", "pages.chapter_id", "pages.title", "pages.number", "pages.is_public").
		From(br.entity.TableName()).
		Join("chapters on books.id = chapters.book_id").
		Join("pages on chapters.id = pages.chapter_id").
		Where("books.id "+"= ?", bookID).
		OrderBy("chapter_id ASC", "number ASC").
		ToSql()

	if err != nil {
		return nil, err
	}

	q := postgres.Query{Name: br.Name + ".GetTOCSectionsFromPages", Raw: query}

	var rows pgx.Rows
	rows, err = br.db.DB().QueryRawContextMulti(ctx, q, args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	sections := make([]*entity2.Section, 0)
	for rows.Next() {
		var section entity2.Section
		section.Level = "page"
		err = rows.Scan(&section.ID, &section.ParentID, &section.Title, &section.Order, &section.IsPublic)
		if err != nil {
			return nil, err
		}
		sections = append(sections, &section)
	}

	return sections, nil
}
