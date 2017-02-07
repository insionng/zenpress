package tidb

import (
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/tests"
	"github.com/go-xorm/xorm"
	_ "github.com/pingcap/tidb"
)

var showTestSql = true

// [memory, goleveldb, boltdb]

func newTidbEngine(storeType string) (*xorm.Engine, error) {
	if storeType == "memory" {
		return xorm.NewEngine("tidb", storeType+"://tidb/tidb?parseTime=true")
	}

	os.Remove("./tidb_" + storeType)
	return xorm.NewEngine("tidb", storeType+"://./tidb_"+storeType+"/tidb?parseTime=true")
}

func newTidbDriverDB(storeType string) (*sql.DB, error) {
	if storeType == "memory" {
		return sql.Open("tidb", storeType+"://./tidb/tidb?parseTime=true")
	}

	os.Remove("./tidb_" + storeType)
	return sql.Open("tidb", storeType+"://./tidb_"+storeType+"/tidb?parseTime=true")
}

func newCache() core.Cacher {
	return xorm.NewLRUCacher2(xorm.NewMemoryStore(), time.Hour, 1000)
}

func setEngine(engine *xorm.Engine, useCache bool) {
	if useCache {
		engine.SetDefaultCacher(newCache())
	}
	engine.Logger().ShowSQL(showTestSql)
	engine.Logger().SetLevel(core.LOG_DEBUG)
}

func TestTidbGoLevelDBNoCache(t *testing.T) {
	engine, err := newTidbEngine("goleveldb")
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()
	setEngine(engine, false)

	tests.BaseTestAll(engine, t)
	tests.BaseTestAll2(engine, t)
	tests.BaseTestAll3(engine, t)
}

func TestTidbGoLevelDBWithCache(t *testing.T) {
	engine, err := newTidbEngine("goleveldb")
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	setEngine(engine, true)

	tests.BaseTestAll(engine, t)
	tests.BaseTestAll2(engine, t)
}

func TestTidbMemoryNoCache(t *testing.T) {
	engine, err := newTidbEngine("memory")
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()
	setEngine(engine, false)

	tests.BaseTestAll(engine, t)
	tests.BaseTestAll2(engine, t)
	tests.BaseTestAll3(engine, t)
}

func TestTidbMemoryWithCache(t *testing.T) {
	engine, err := newTidbEngine("memory")
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	setEngine(engine, true)

	tests.BaseTestAll(engine, t)
	tests.BaseTestAll2(engine, t)
}

func TestTidbBoltDBNoCache(t *testing.T) {
	engine, err := newTidbEngine("boltdb")
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()
	setEngine(engine, false)

	tests.BaseTestAll(engine, t)
	tests.BaseTestAll2(engine, t)
	tests.BaseTestAll3(engine, t)
}

func TestTidbBoltDBWithCache(t *testing.T) {
	engine, err := newTidbEngine("boltdb")
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	setEngine(engine, true)

	tests.BaseTestAll(engine, t)
	tests.BaseTestAll2(engine, t)
}

const (
	createTableQl = "CREATE TABLE IF NOT EXISTS `big_struct` (`id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `name` TEXT NULL, `title` TEXT NULL, `age` TEXT NULL, `alias` TEXT NULL, `nick_name` TEXT NULL);"
	dropTableQl   = "DROP TABLE IF EXISTS `big_struct`;"
)

func BenchmarkTidbDriverInsert(t *testing.B) {
	tests.DoBenchDriver(func() (*sql.DB, error) {
		return newTidbDriverDB("goleveldb")
	}, createTableQl, dropTableQl,
		tests.DoBenchDriverInsert, t)
}

func BenchmarkTidbDriverFind(t *testing.B) {
	tests.DoBenchDriver(func() (*sql.DB, error) {
		return newTidbDriverDB("goleveldb")
	}, createTableQl, dropTableQl,
		tests.DoBenchDriverFind, t)
}

func BenchmarkTidbNoCacheInsert(t *testing.B) {
	t.StopTimer()
	engine, err := newTidbEngine("goleveldb")
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	tests.DoBenchInsert(engine, t)
}

func BenchmarkTidbNoCacheFind(t *testing.B) {
	t.StopTimer()
	engine, err := newTidbEngine("goleveldb")
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	//engine.ShowSQL = true
	tests.DoBenchFind(engine, t)
}

func BenchmarkTidbNoCacheFindPtr(t *testing.B) {
	t.StopTimer()
	engine, err := newTidbEngine("goleveldb")
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()
	//engine.ShowSQL = true
	tests.DoBenchFindPtr(engine, t)
}

func BenchmarkTidbCacheInsert(t *testing.B) {
	t.StopTimer()
	engine, err := newTidbEngine("goleveldb")
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	engine.SetDefaultCacher(newCache())
	tests.DoBenchInsert(engine, t)
}

func BenchmarkTidbCacheFind(t *testing.B) {
	t.StopTimer()
	engine, err := newTidbEngine("goleveldb")
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	engine.SetDefaultCacher(newCache())
	tests.DoBenchFind(engine, t)
}

func BenchmarkTidbCacheFindPtr(t *testing.B) {
	t.StopTimer()
	engine, err := newTidbEngine("goleveldb")
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	engine.SetDefaultCacher(newCache())
	tests.DoBenchFindPtr(engine, t)
}
