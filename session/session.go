package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxlifeTime int64
}

func NewManager(providerName, cookieName string, maxlifeTime int64) (*Manager, error) {
	provider, ok := provides[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}
	return &Manager{cookieName: cookieName, provider: provider, maxlifeTime: maxlifeTime}, nil
}

var globalSessions *Manager
var provides = make(map[string]Provider)

func init() {
	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGc(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(manager.maxlifeTime),
		}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}
