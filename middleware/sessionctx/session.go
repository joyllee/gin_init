package sessionctx

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"local/gin_init/model/db"
)

type SessionCtx struct {
	TenantId    string
	Dbx         *sqlx.DB
	RedisClient *redis.Client
}

func NewSessionCtx(c *gin.Context) {
	c.Set("session", &SessionCtx{
		TenantId:    "",
		Dbx:         db.SqlDb,
		RedisClient: db.RedisClient,
	})
}

func (s *SessionCtx) GetDbx() *sqlx.DB {
	return s.Dbx
}

func (s *SessionCtx) GetRedisClient() *redis.Client {
	return s.RedisClient
}

func (s *SessionCtx) GetTenantId() string {
	return s.TenantId
}

func (s *SessionCtx) SetTenantId(tenantId string) {
	s.TenantId = tenantId
}
