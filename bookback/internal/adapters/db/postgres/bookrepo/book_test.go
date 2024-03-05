package bookrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/models"
	mocks2 "github.com/SShlykov/zeitment/bookback/tests/mocks"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func newTestBook() *models.Book {
	return &models.Book{
		ID:          faker.UUIDHyphenated(),
		Title:       faker.Username(),
		Author:      faker.Name(),
		Owner:       faker.UUIDHyphenated(),
		Description: faker.Sentence(),
		IsPublic:    false,
		Publication: models.NewNullTime(time.Now(), true),
		ImageLink:   models.NewNullString(faker.Word(), true),
		MapLink:     models.NewNullString(faker.Word(), true),
		MapParamsID: models.NewNullString(faker.Word(), true),
		Variables:   []string{faker.Word(), faker.Word(), faker.Word()},
	}
}

func rowFromBook(book *models.Book) *mocks2.ScanResult {
	return mocks2.NewScanResult([]interface{}{book.ID, book.CreatedAt, book.UpdatedAt, book.DeletedAt, book.Owner, //nolint:gofmt
		book.Title, book.Author, book.Description, book.IsPublic, book.Publication, book.ImageLink, book.MapLink,
		book.MapParamsID, book.Variables,
	})
}

func inits(ctrl *gomock.Controller) (Repository, *models.Book) {
	client := mocks2.NewMockClient(ctrl)
	db := mocks2.NewMockDB(ctrl)

	testBook := newTestBook()
	row := rowFromBook(testBook)

	client.EXPECT().DB().Return(db)
	db.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(row)

	repo := NewRepository(client)

	return repo, testBook
}

func TestRepository_FindByID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo, testBook := inits(ctrl)

	book, err := repo.FindByID(nil, testBook.ID)
	assert.Empty(t, err)
	assert.Equal(t, testBook, book)
}

func TestRepository_Create(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks2.NewMockClient(ctrl)
	db := mocks2.NewMockDB(ctrl)

	testBook := &models.Book{}

	row := mocks2.NewScanResult([]interface{}{faker.UUIDHyphenated()})

	client.EXPECT().DB().Return(db)
	db.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(row)

	repo := NewRepository(client)
	id, err := repo.Create(nil, testBook)
	assert.Empty(t, err)
	assert.NotEmpty(t, id)
}

func TestRepository_Update(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo, testBook := inits(ctrl)

	book, err := repo.Update(nil, testBook.ID, testBook)
	assert.Empty(t, err)
	assert.Equal(t, testBook, book)
}

func TestRepository_Delete(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo, testBook := inits(ctrl)

	book, err := repo.Delete(nil, testBook.ID)
	assert.Empty(t, err)
	assert.Equal(t, testBook, book)
}
