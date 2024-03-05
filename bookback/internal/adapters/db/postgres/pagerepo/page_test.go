package pagerepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/adapters/db/postgres/bookrepo"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	mocks2 "github.com/SShlykov/zeitment/bookback/tests/mocks"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func newTestPage() *models.Page {
	return &models.Page{
		ID:          "1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       faker.Word(),
		Text:        "Test text",
		ChapterID:   faker.UUIDHyphenated(),
		IsPublic:    true,
		MapParamsID: models.NewNullString(faker.UUIDHyphenated(), true),
	}
}

func rowFromPage(page *models.Page) *mocks2.ScanResult {
	return mocks2.NewScanResult([]interface{}{page.ID, page.CreatedAt, page.UpdatedAt, page.DeletedAt,
		page.Title, page.Text, page.ChapterID, page.IsPublic, page.MapParamsID})
}

func initPages(ctrl *gomock.Controller) (bookrepo.Repository, *models.Page) {
	client := mocks2.NewMockClient(ctrl)
	db := mocks2.NewMockDB(ctrl)

	testPage := newTestPage()
	row := rowFromPage(testPage)

	client.EXPECT().DB().Return(db)
	db.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(row)

	repo := bookrepo.NewRepository(client)

	return repo, testPage
}

func TestRepository_FindByID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo, testPage := initPages(ctrl)

	page, err := repo.FindByID(nil, testPage.ID)

	assert.Empty(t, err)
	assert.Equal(t, testPage, page)
}

func TestRepository_Create(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks2.NewMockClient(ctrl)
	db := mocks2.NewMockDB(ctrl)

	testPage := &models.Page{}

	row := mocks2.NewScanResult([]interface{}{faker.UUIDHyphenated()})

	client.EXPECT().DB().Return(db)
	db.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(row)

	repo := bookrepo.NewRepository(client)
	id, err := repo.Create(nil, testPage)
	assert.Empty(t, err)
	assert.NotEmpty(t, id)
}

func TestRepository_Update(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo, testPage := initPages(ctrl)

	page, err := repo.Update(nil, testPage.ID, testPage)
	assert.Empty(t, err)
	assert.Equal(t, testPage, page)
}

func TestRepository_Delete(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo, testPage := initPages(ctrl)

	page, err := repo.Delete(nil, testPage.ID)
	assert.Empty(t, err)
	assert.Equal(t, testPage, page)
}
