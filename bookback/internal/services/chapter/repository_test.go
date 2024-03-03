package chapter

import (
	"github.com/SShlykov/zeitment/bookback/internal/models"
	mocks2 "github.com/SShlykov/zeitment/bookback/tests/mocks"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func newTestChapter() *models.Chapter {
	return &models.Chapter{
		ID:          faker.UUIDHyphenated(),
		Title:       faker.Username(),
		IsPublic:    false,
		BookID:      faker.UUIDHyphenated(),
		Number:      7,
		Text:        faker.Sentence(),
		MapLink:     models.NewNullString(faker.Word(), true),
		MapParamsID: models.NewNullString(faker.Word(), true),
	}
}

func rowFromChapter(chapter *models.Chapter) *mocks2.ScanResult {
	return mocks2.NewScanResult([]interface{}{chapter.ID, chapter.CreatedAt, chapter.UpdatedAt, chapter.DeletedAt,
		chapter.Title, chapter.Number, chapter.Text, chapter.BookID, chapter.IsPublic, chapter.MapLink,
		chapter.MapParamsID,
	})
}

func initChapters(ctrl *gomock.Controller) (Repository, *models.Chapter) {
	client := mocks2.NewMockClient(ctrl)
	db := mocks2.NewMockDB(ctrl)

	testChapter := newTestChapter()
	row := rowFromChapter(testChapter)

	client.EXPECT().DB().Return(db)
	db.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(row)

	repo := NewRepository(client)

	return repo, testChapter
}

func TestRepository_FindByID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo, testChapter := initChapters(ctrl)

	book, err := repo.FindByID(nil, testChapter.ID)
	assert.Empty(t, err)
	assert.Equal(t, testChapter, book)
}

func TestRepository_Create(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks2.NewMockClient(ctrl)
	db := mocks2.NewMockDB(ctrl)

	testChapter := &models.Chapter{}

	row := mocks2.NewScanResult([]interface{}{faker.UUIDHyphenated()})

	client.EXPECT().DB().Return(db)
	db.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(row)

	repo := NewRepository(client)
	id, err := repo.Create(nil, testChapter)
	assert.Empty(t, err)
	assert.NotEmpty(t, id)
}

func TestRepository_Update(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo, testChapter := initChapters(ctrl)

	book, err := repo.Update(nil, testChapter.ID, testChapter)
	assert.Empty(t, err)
	assert.Equal(t, testChapter, book)
}

func TestRepository_Delete(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo, testChapter := initChapters(ctrl)

	book, err := repo.Delete(nil, testChapter.ID)
	assert.Empty(t, err)
	assert.Equal(t, testChapter, book)
}
