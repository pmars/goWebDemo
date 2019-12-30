package inits

import (
	"time"

	"goWebDemo/conf"
	"goWebDemo/dao"

	"github.com/garyburd/redigo/redis"
	"github.com/pmars/beego/logs"
)

func initRedisEngine() {
	logs.Debug("Init Redis Pool Now...")

	if conf.Config.Redis.IsAuth {
		dao.RedisPool = newRedisPoolAuth(
			conf.Config.Redis.Host,
			conf.Config.Redis.Password,
			conf.Config.Redis.MaxIdle,
			conf.Config.Redis.MaxActive,
			time.Duration(conf.Config.Redis.DialTimeout)*time.Second)
	} else {
		dao.RedisPool = newRedisPool(
			conf.Config.Redis.Host,
			conf.Config.Redis.MaxIdle,
			conf.Config.Redis.MaxActive,
			time.Duration(conf.Config.Redis.DialTimeout)*time.Second)
	}
	if err := dao.RedisPool.Pool.TestOnBorrow(dao.RedisPool.Pool.Get(), time.Now()); err != nil {
		panic(err)
	}
	logs.Debug("Init Redis Pool Done!!!")

	// 初始化销量缓存（暂时不需要，第一次查询的时候加载到缓存）
	// models.MInitSalesVolume()
	// logs.Debug("Init SalesVolume Done!!!")
}

func newRedisPoolAuth(server, auth string, maxIdle, maxActive int, dialTimeout time.Duration) *dao.RedisDao {
	return &dao.RedisDao{
		Pool: &redis.Pool{
			MaxIdle:     maxIdle,
			MaxActive:   maxActive,
			Wait:        true,
			IdleTimeout: dialTimeout,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", server)
				if err != nil {
					return nil, err
				}
				if _, err := c.Do("AUTH", auth); err != nil {
					c.Close()
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
	}
}

func newRedisPool(server string, maxIdle, maxActive int, dialTimeout time.Duration) *dao.RedisDao {
	return &dao.RedisDao{
		Pool: &redis.Pool{
			MaxIdle:     maxIdle,
			MaxActive:   maxActive,
			Wait:        true,
			IdleTimeout: dialTimeout,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", server)
				if err != nil {
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
	}
}
