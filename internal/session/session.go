package session

import (
	"errors"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/bagashiz/sea-salon/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// New returns a new session manager instance.
func New(cfg *config.App, sessionStore any) (*scs.SessionManager, error) {
	lifetime, err := time.ParseDuration(cfg.SessionLifetime)
	if err != nil {
		return nil, err
	}

	sessionManager := scs.New()
	sessionManager.Lifetime = lifetime
	sessionManager.Cookie.Secure = cfg.Env == "production"

	if sessionStore != nil {
		store, err := newStore(sessionStore)
		if err != nil {
			return nil, err
		}
		sessionManager.Store = store
	}

	return sessionManager, nil
}

// newStore returns a new session store based on the provided store type.
func newStore(store any) (scs.Store, error) {
	switch store {
	case store.(*pgxpool.Pool):
		return pgxstore.New(store.(*pgxpool.Pool)), nil
		// add more store types here
	}
	return nil, errors.New("unknown session store type")
}
