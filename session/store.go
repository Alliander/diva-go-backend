package session

import (
	"github.com/alliander/diva-go-backend/config"
	"github.com/gorilla/sessions"
	"github.com/quasoft/memstore"
	"net/http"
)

var store *memstore.MemStore

func Init(c *config.Config) {
	// TODO: Add support for more stores
	store = memstore.NewMemStore([]byte(c.CookieSecret))
}

func Get(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "diva-session")
}
