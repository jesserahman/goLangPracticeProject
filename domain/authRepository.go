package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/jesserahman/goLangPracticeProject/logger"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct {
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	verifyUrl := buildVerifyURL(token, routeName, vars)
	if response, err := http.Get(verifyUrl); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false
	} else {
		m := map[string]bool{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth server" + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

// Sample: localhost:8081/auth/verify?token=aaaa.bbbb.cccc&routeName=GetCustomers&customer_id=2000
func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{
		Scheme: "http",
		Host:   os.Getenv("DOCKER_AUTH_SERVICE_HOST"),
		Path:   os.Getenv("AUTH_SERVICE_VERIFY_PATH"),
	}

	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)

	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}
