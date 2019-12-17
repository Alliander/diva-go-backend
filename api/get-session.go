package api

import (
	"encoding/json"
	"fmt"
	"github.com/alliander/diva-go-backend/session"
	"net/http"
)

// GetSession returns all the session data we have to the frontend
// It's not used by diva-react, so you can remove this handler if you do not need it
//func GetSession(store *sessions.FilesystemStore) http.HandlerFunc {
func GetSession(w http.ResponseWriter, r *http.Request) {
	divaSession, _ := session.Get(r)

	// Convert map[interface{}]interface{} to map[string]string ...
	sessionMap := make(map[string]string)
	for k, v := range divaSession.Values {
		key := fmt.Sprintf("%v", k)
		value := fmt.Sprintf("%v", v)
		sessionMap[key] = value
	}

	enc := json.NewEncoder(w)
	err := enc.Encode(sessionMap)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
