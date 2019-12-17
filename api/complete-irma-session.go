package api

import (
	"github.com/alliander/diva-go-backend/session"
	"net/http"
)

// CompleteIrmaSession returns the result of the IRMA session to the frontend
// diva-react expects a serialized irmaSessionResult object, but you can change
// it before sending to the frontend
func CompleteIrmaSession(w http.ResponseWriter, r *http.Request) {
	divaSession, _ := session.Get(r)
	irmaSessionResult := divaSession.Values["irma-session-result"].(string)

	// TODO: parse result and use it
	//irmaParsedResult := server.SessionResult{}
	//err := json.Unmarshal(irmaSessionResult, &irmaParsedResult)

	// Remove session data after frontend retrieved it
	//session.Values["irma-session-result"] = nil
	//session.Save(r, w)

	w.Header().Add("Content-Type", "text/json")
	w.Write([]byte(irmaSessionResult))
}
