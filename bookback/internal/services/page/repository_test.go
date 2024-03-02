package page

import (
	"github.com/SShlykov/zeitment/bookback/internal/mocks"
	"github.com/SShlykov/zeitment/bookback/internal/models"
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

func rowFromPage(page *models.Page) *mocks.ScanResult {
	return mocks.NewScanResult([]interface{}{page.ID, page.CreatedAt, page.UpdatedAt, page.DeletedAt,
		page.Title, page.Text, page.ChapterID, page.IsPublic, page.MapParamsID})
}

func initPages(ctrl *gomock.Controller) (Repository, *models.Page) {
	client := mocks.NewMockClient(ctrl)
	db := mocks.NewMockDB(ctrl)

	testPage := newTestPage()
	row := rowFromPage(testPage)

	client.EXPECT().DB().Return(db)
	db.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(row)

	repo := NewRepository(client)

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
	client := mocks.NewMockClient(ctrl)
	db := mocks.NewMockDB(ctrl)

	testPage := &models.Page{}

	row := mocks.NewScanResult([]interface{}{faker.UUIDHyphenated()})

	client.EXPECT().DB().Return(db)
	db.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(row)

	repo := NewRepository(client)
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
