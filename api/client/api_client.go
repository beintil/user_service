package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	lue "github.com/beintil/user_service/pkg/lu_error"
	"net/http"
)

// Config Connected config
type Config struct {
	// Host server
	Host string
	// Port server
	Port int
}

// Service struct
type Service struct {
	// Http request
	r *http.Request
	// cfg
	cfg Config
}

func NewService(r *http.Request, cfg Config) *Service {
	return &Service{r: r, cfg: cfg}
}

// apiResponseUser There struct returns from http functions
type apiResponseUser struct {
	// Response message
	Message string `json:"message"`
	// Response data
	Data User `json:"data"`
}

func (s *Service) client(dest any, method string, url string, requestBody []byte) lue.IError {
	httpClient := http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return lue.New(err.Error(), lue.BadRequest).SetHttpCode(http.StatusBadRequest)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return lue.New(err.Error(), lue.BadRequest).SetHttpCode(http.StatusBadRequest)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&dest)
	if err != nil {
		return lue.New(err.Error(), lue.InternalServerError).SetHttpCode(http.StatusInternalServerError)
	}
	return nil
}

func (s *Service) CheckAuthData(email string, password string) (user *User, e lue.IError) {
	type requestAuth struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	rAuth := requestAuth{email, password}

	request, err := json.Marshal(rAuth)
	if err != nil {
		return nil, lue.New(err.Error(), lue.InternalServerError).SetHttpCode(http.StatusInternalServerError)
	}
	res := apiResponseUser{}
	e = s.client(&res, http.MethodPost, fmt.Sprintf("%s:%d/user-service/api/v1/verify", s.cfg.Host, s.cfg.Port), request)
	if e != nil {
		return nil, e
	}
	return &res.Data, nil
}
