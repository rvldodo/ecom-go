package cookies

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dodo/ecom/lib/jwt"
)

type CookieData struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}

var cookieName = "Cookie"

func SetCookie(value *CookieData, w http.ResponseWriter, r *http.Request) error {
	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	data, err := jwt.CreateTokenCookie(string(v))
	if err != nil {
		return err
	}

	if c.Value == "" {
		c.Name = cookieName
		c.Value = data
		c.Expires = time.Now().Add(24 * time.Hour)
		http.SetCookie(w, c)
	}

	return nil
}
