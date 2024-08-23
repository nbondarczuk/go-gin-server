package tag

import (
	"encoding/json"
	"go-gin-server/internal/cache"
	"go-gin-server/internal/repository/entity/tag"
)

const TagEntityName = "tag"

type TagCache struct {
	cache *cache.Redis
}

// NewTagCache uses the connection allocated in the init of the cache module.
// It does not open each connection per request but it reuses the initial one.
func NewTagCache() (*TagCache, error) {
	cache, err := cache.WithRedis()
	if err != nil {
		return nil, err
	}
	return &TagCache{
		cache: cache,
	}, nil
}

// Check does a dive into the redis cache for an id.
func (tc *TagCache) Check(id string) (tag.Tag, bool, error) {
	val, err := tc.cache.Client.Get(TagEntityName).Result()
	if err != nil {
		return tag.Tag{}, false, err
	}
	var tag tag.Tag
	json.Unmarshal([]byte(val), &tag)
	return tag, true, nil
}

// Flush purges cache for single element, when it was modified or deleted.
func (tc *TagCache) Flush(id string) error {
	return nil
}

// Purge makes empty the whole cache of the collection.
func (tc *TagCache) Purge() error {
	return nil
}
