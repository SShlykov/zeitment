package paragraph

import (
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/go-faker/faker/v4"
	"time"
)

func newTestParagraph() *models.Paragraph {
	return &models.Paragraph{
		ID:        "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Text:      "Test text",
		PageID:    faker.UUIDHyphenated(),
		IsPublic:  true,
	}
}
