package rulesengine

import (
	// "crypto/sha256"
	"encoding/json"
	"sync"
	"time"
)

// Is sync.Map the best choice for our use case? Would it better to use a go map with mutex?

// TTL expiration check every 5 min
// When TTL expires, perform raw JSON hash diff to determine if tree rebuild is needed

type cacheObject struct {
    ts       time.Time
	cfg      json.RawMessage // TODO: change to hash
	ruleSets []cacheRuleSet
}
type cacheRuleSet struct {
	stage       string
	name        string
	modelGroups []cacheModelGroup
}
type cacheModelGroup struct {
	weight       int
	version      string
	analyticsKey string
	defaults     []Function
	root         Node
}

func NewCacheObject(cfg config) (cacheObject, error) {
	return cacheObject{}, nil
}

type cacher interface {
	Get(string) (*cacheObject)
	Set(string, cacheObject)
	Delete(id string)
}

type cache struct {
	*sync.Map
	// cache map[string]cacheObject
}

func (c *cache) Get(id string) (data *cacheObject) {
	if val, ok := c.Map.Load(id); ok {
		return val.(*cacheObject)
	}
	return nil
}

func (c *cache) Save(id string, data cacheObject) {
	c.Map.Store(id, data)
}

func (c *cache) Delete(id string) {
	c.Map.Delete(id)
}