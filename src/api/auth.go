package api

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/google/uuid"
	"log"
	"math/big"
	"net/http"
)

// AuthNZ functionality is here only for demonstration purposes.
// In the actual setup, this will be provided by the API Gateway product (e.g., Apigee or AWS APIGateway).
// So that we only have to focus on implementing logic specific to our use cases (implementing REST resources).

// https://medium.com/@satyendra.jaiswal/securing-apis-oauth-2-0-and-api-keys-best-practices-6d779b00d934

func NewOAuth2Manager() (*server.Server, *store.ClientStore, error) {
	manager := manage.NewDefaultManager()

	// client memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore()) // in-memory token store

	// demonstration purposes
	clientId := "000000"
	secret := "999999"
	domain := "http://localhost"

	clientStore := store.NewClientStore()
	err := clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: secret,
		Domain: domain,
	})
	if err != nil {
		return nil, nil, err
	}

	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.UserAuthorizationHandler = func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		return clientId, nil
	}

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("internal error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("response error:", re.Error.Error())
	})

	return srv, clientStore, nil
}

func (c Configuration) register(w http.ResponseWriter, r *http.Request) {

	clientId := uuid.New().String()

	byteSize := 32
	secret, err := randomString(byteSize)
	if err != nil {
		log.Println("error: /register", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = c.clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: secret,
		Domain: "http://localhost",
	})

	if err != nil {
		log.Println("error: /register", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	type entry struct {
		ClientId string `json:"clientId"`
		Secret   string `json:"secret"`
	}

	e := entry{
		ClientId: clientId,
		Secret:   secret,
	}

	err = json.NewEncoder(w).Encode(e)
	if err != nil {
		log.Println("error: /register", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func randomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func randomStringUrlSafe(n int) (string, error) {
	b, err := randomString(n)
	return base64.URLEncoding.EncodeToString([]byte(b)), err
}

func (c Configuration) authorize(w http.ResponseWriter, r *http.Request) {
	err := c.oauth2Server.HandleAuthorizeRequest(w, r)
	if err != nil {
		log.Println("error: /authorize", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (c Configuration) token(w http.ResponseWriter, r *http.Request) {
	err := c.oauth2Server.HandleTokenRequest(w, r)
	if err != nil {
		log.Println("error: /oauth/token", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
