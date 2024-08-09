package global

import (
	"admin-gin/config"
	"admin-gin/utils/timer"
	"github.com/gorilla/websocket"
	"github.com/olivere/elastic/v7"
	"github.com/qiniu/qmgo"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"sync"
)

var (
	GVA_DB     *gorm.DB
	XTK_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	GVA_REDIS  redis.UniversalClient
	GVA_MONGO  *qmgo.QmgoClient
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	// GVA_LOG    *oplogging.Logger
	GVA_LOG                 *zap.Logger
	GVA_Timer               timer.Timer = timer.NewTimerTask()
	GVA_Concurrency_Control             = &singleflight.Group{}

	ESClient *elastic.Client

	BlackCache local_cache.Cache
	rwLock     sync.RWMutex
	LockMap    sync.Map

	lock sync.Mutex

	WsConnMap = map[uint]*websocket.Conn{}
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	rwLock.RLock()
	defer rwLock.RUnlock()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	rwLock.RLock()
	defer rwLock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

func GetWsConnByUserID(userID uint) *websocket.Conn {
	lock.Lock()
	defer lock.Unlock()
	conn, ok := WsConnMap[userID]
	if !ok {
		return nil
	}
	return conn
}

func SetWsConnByUserID(userID uint, conn *websocket.Conn) bool {
	lock.Lock()
	defer lock.Unlock()
	WsConnMap[userID] = conn
	return true
}
