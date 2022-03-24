package mock

import (
	"time"

	"github.com/eazylaykzy/snippetbox/pkg/models"
)

/*var mockSnippet = &models.Snippet{
	ID:      1,
	Created: time.Now(),
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Expires: time.Date(2023, 03, 16, 22, 46, 04, 00, time.FixedZone("WAT", 1*60*60)),
}*/

var mockSnippet = []*models.Snippet{
	{
		ID:      1,
		Created: time.Now(),
		Title:   "An old silent pond",
		Content: "An old silent pond...",
		Expires: time.Date(2023, 03, 16, 22, 46, 04, 00, time.FixedZone("WAT", 1*60*60)),
	},
	{
		ID:      2,
		Created: time.Now(),
		Title:   "Easy on the world",
		Content: "Let's build a better world together",
		Expires: time.Date(2023, 04, 16, 22, 46, 04, 00, time.FixedZone("WAT", 1*60*60)),
	},
	{
		ID:      3,
		Created: time.Now(),
		Title:   "Great Nigeria",
		Content: "Nigeria will definitely be great again!",
		Expires: time.Date(2023, 05, 16, 22, 46, 04, 00, time.FixedZone("WAT", 1*60*60)),
	},
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet[0], nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return mockSnippet, nil
}
