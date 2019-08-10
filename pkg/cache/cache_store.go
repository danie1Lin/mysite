package cache

import (
	"github.com/go-redis/redis"
)

type CacheStore interface {
	New(key string, data map[string]interface{}) (err error)
	GetData(key string) (data map[string]string, err error)
	WriteKey(key, field string, value interface{}) (err error)
	ReadKey(key, field string) (value string, err error)
}

type RedisCacheStore struct {
	client *redis.Client
	prefix string
}

func NewRedisCacheStore(addr, password, prefix string) *RedisCacheStore {
	return &RedisCacheStore{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		}),
	}
}

func (s *RedisCacheStore) GetData(key string) (data map[string]string, err error) {
	data, err = s.client.HGetAll(s.prefix + key).Result()
	return
}

func (s *RedisCacheStore) New(key string, data map[string]interface{}) (err error) {
	if data == nil {
		return DataInvalidError
	}

	n, err := s.client.Exists(key).Result()
	if err != nil {
		return err
	}
	if n != 0 {
		return SessionIDDuplicatedError
	}
	_, err = s.client.HMSet(s.prefix+key, data).Result()
	if err != nil {
		return err
	}
	return
}

func (s *RedisCacheStore) WriteKey(key, field string, value interface{}) (err error) {
	_, err = s.client.HSet(s.prefix+key, field, value).Result()
	return
}

func (s *RedisCacheStore) ReadKey(key, field string) (value string, err error) {
	_, err = s.client.HGet(s.prefix+key, field).Result()
	if err != nil {
		return "", err
	}
	return
}
