package cache

import (
	"go-gin-server/internal/repository/entity"
)

type TagCache struct {
	cache *Redis
}

// NewTagCache uses the connection allocated in the init of the cache module.
// It does not open each connection per request but it reuses the initial one.
func NewTagCache() (*TagCache, error) {
	cache, err := WithRedis()
	if err != nil {
		return nil, err
	}
	return &TagCache{
		cache: cache,
	}, nil
}

// Check does a dive into the redis cache for an id.
func (tc *TagCache) Get(id string) (entity.Tag, bool, error) {
	return entity.Tag{}, false, nil
}

// Flush purges cache for single element, when it was modified or deleted.
func (tc *TagCache) Flush(id string) error {
	return nil
}

// Purge makes empty the whole cache of the collection.
func (tc *TagCache) Purge() error {
	return nil
}
