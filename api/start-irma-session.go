package api

import (
	"encoding/json"
	"fmt"
	"github.com/alliander/diva-go-backend/session"
	"github.com/privacybydesign/irmago"
	"github.com/privacybydesign/irmago/server"
	"github.com/privacybydesign/irmago/server/irmaserver"
	"io/ioutil"
	"net/http"
)

func handleIrmaSessionResult(w http.ResponseWriter, r *http.Request) irmaserver.SessionHandler {
	return func(result *server.SessionResult) {
		if result.ProofStatus != irma.ProofStatusValid {
			fmt.Println("Invalid IRMA Proof")
			http.Error(w, "Invalid IRMA Proof", http.StatusUnauthorized)
			return
		}

		divaSession, _ := session.Get(r)
		divaSession.Values["irma-session-result"] = server.ToJson(result)
		err := divaSession.Save(r, w)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

// StartIrmaSession starts a new IRMA session, you can add logic here to check if this request is allowed
func StartIrmaSession(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ### TODO: Check here whether this request is allowed before starting issuance/disclosure! ###
	sessionPointer, token, err := irmaserver.StartSession(request, handleIrmaSessionResult(w, r))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Created session with token ", token)
	// Initialize session
	divaSession, _ := session.Get(r)
	//session, _ := store.Get(r, "diva-session")
	err = divaSession.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "text/json")

	enc := json.NewEncoder(w)
	err = enc.Encode(&server.SessionPackage{
		SessionPtr: sessionPointer,
		Token:      token,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
