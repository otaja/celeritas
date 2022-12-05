package session

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
)

type Session struct {
	CookieLifetime string
	CookiePersists string
	CookieName     string
	CookieDomain   string
	CookieSecure   string
	SessionType    string
	DBPool         *sql.DB
	RedisPool      *redis.Pool
}

func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	// how long should sessions last
	minutes, err := strconv.Atoi(c.CookieLifetime)
	if err != nil {
		minutes = 60
	}
	// should Cookies persist?
	if strings.ToLower(c.CookiePersists) == "true" {
		persist = true
	}

	//must Cookies be secure?
	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	}

	// create Session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = c.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// select Session Store
	switch strings.ToLower(c.SessionType) {
	case "redis":
		session.Store = redisstore.New(c.RedisPool)
	case "mysql", "mariadb":
		session.Store = mysqlstore.New(c.DBPool)
	case "postgres", "postgresql":
		session.Store = postgresstore.New(c.DBPool)

	default:
		// cookie
	}

	return session

}