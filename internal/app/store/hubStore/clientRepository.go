package hubStore

import (
	"Chat/internal/app/model"
	"github.com/google/uuid"
)

// ClientRepository storage of chat users
type ClientRepository interface {
	All() (map[uuid.UUID]*model.Client, error)
	AllUsers() ([]*model.ChatUser, error)
	Find(id uuid.UUID) (*model.Client, error)
	CountByOriginalName(name string) (int, error)
	Add(client *model.Client) (string, error)
	Remove(guid uuid.UUID) error
}