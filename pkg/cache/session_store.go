package cache

import (
	"fmt"

	"github.com/google/uuid"
)

type SessionStore interface {
	New(data map[string]interface{}) (sessionID string, err error)
	GetData(sessionID string) (data map[string]string, err error)
	WriteKey(sessionID, field string, value interface{}) (err error)
	ReadKey(sessionID, field string) (value string, err error)
	HeaderKey() string
	ContextKey() string
}

type SessionRedisStore struct {
	CacheStore
	headerKey, contextKey string
}

func (s *SessionRedisStore) HeaderKey() string { return s.headerKey }

func (s *SessionRedisStore) ContextKey() string { return s.contextKey }

func (s *SessionRedisStore) New(data map[string]interface{}) (sessionID string, err error) {
	uid := uuid.New().String()
	err = s.CacheStore.New(uid, data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return uid, nil
}

func NewSessionRedisStore(cacheStore CacheStore, headerKey, contextKey string) *SessionRedisStore {
	return &SessionRedisStore{cacheStore, headerKey, contextKey}
}
