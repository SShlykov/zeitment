package paragraphrepo

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

func newTestParagraph() *models.Paragraph {
	return &models.Paragraph{
		ID:        "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     faker.Word(),
		Text:      faker.Word(),
		Type:      faker.Word(),
		PageID:    faker.UUIDHyphenated(),
		IsPublic:  true,
	}
}

func rowFromParagraph(paragraph *models.Paragraph) *mocks2.ScanResult {
	return mocks2.NewScanResult([]interface{}{paragraph.ID, paragraph.CreatedAt, paragraph.UpdatedAt, paragraph.DeletedAt,
		paragraph.Title, paragraph.Text, paragraph.Type, paragraph.IsPublic, paragraph.PageID})
}

func initParagraphs(ctrl *gomock.Controller) (bookrepo.Repository, *models.Paragraph) {
	client := mocks2.NewMockClient(ctrl)
	db := mocks2.NewMockDB(ctrl)

	testParagraph := newTestParagraph()
	row := rowFromParagraph(testParagraph)

	client.EXPECT().DB().Return(db)
	db.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(row)

	repo := bookrepo.NewRepository(client)

	return repo, testParagraph
}

func TestRepository_FindByID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo, testParagraph := initParagraphs(ctrl)

	paragraph, err := repo.FindByID(nil, testParagraph.ID)

	assert.Empty(t, err)
	assert.Equal(t, testParagraph, paragraph)
}

func TestRepository_Create(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks2.NewMockClient(ctrl)
	db := mocks2.NewMockDB(ctrl)

	testParagraph := &models.Paragraph{}

	row := mocks2.NewScanResult([]interface{}{faker.UUIDHyphenated()})

	client.EXPECT().DB().Return(db)
	db.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(row)

	repo := bookrepo.NewRepository(client)
	id, err := repo.Create(nil, testParagraph)
	assert.Empty(t, err)
	assert.NotEmpty(t, id)
}

func TestRepository_Update(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo, testParagraph := initParagraphs(ctrl)

	paragraph, err := repo.Update(nil, testParagraph.ID, testParagraph)

	assert.Empty(t, err)
	assert.Equal(t, testParagraph, paragraph)
}

func TestRepository_Delete(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo, testParagraph := initParagraphs(ctrl)

	paragraph, err := repo.Delete(nil, testParagraph.ID)

	assert.Empty(t, err)
	assert.Equal(t, testParagraph, paragraph)
}
