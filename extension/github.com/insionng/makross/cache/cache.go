package cache

import (
	"github.com/insionng/makross/cache"

	"qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/insionng/makross/cache",

	"MakrossCacheStoreKey": cache.MakrossCacheStoreKey,

	"Cacher":     cache.Cacher,
	"DecodeGob":  cache.DecodeGob,
	"EncodeGob":  cache.EncodeGob,
	"EncodeSha1": cache.EncodeSha1,
	"IsExist":    cache.IsExist,
	"Register":   cache.Register,
	"Version":    cache.Version,

	"New":         cache.New,
	"NewTagCache": cache.NewTagCache,
	"Store":       cache.Store,

	"Engine":          spec.StructOf((*cache.Engine)(nil)),
	"FileCacher":      spec.StructOf((*cache.FileCacher)(nil)),
	"NewFileCacher":   cache.NewFileCacher,
	"Item":            spec.StructOf((*cache.Item)(nil)),
	"MemoryCacher":    spec.StructOf((*cache.MemoryCacher)(nil)),
	"NewMemoryCacher": cache.NewMemoryCacher,
	"MemoryItem":      spec.StructOf((*cache.MemoryItem)(nil)),
	"Options":         spec.StructOf((*cache.Options)(nil)),
	"TagCache":        spec.StructOf((*cache.TagCache)(nil)),
	"TagSet":          spec.StructOf((*cache.TagSet)(nil)),
	"NewTagSet":       cache.NewTagSet,
}
