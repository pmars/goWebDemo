package dao

import (
	"encoding/json"
	"errors"

	"github.com/garyburd/redigo/redis"
)

type RedisDao struct {
	Pool *redis.Pool
}

var RedisPool *RedisDao

const RedisOKReturn = "redigo: nil returned" // redis正常返回
func RedisOK(err error) error {
	if err == nil {
		return nil
	} else if err.Error() == RedisOKReturn {
		return nil
	} else {
		return err
	}
}

// TTL key
// 以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live)。
func (pool *RedisDao) TTL(key string) (ttl int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	ttl, err = redis.Int(conn.Do("TTL", key))
	return ttl, RedisOK(err)
}

// 集合操作
// SADD 可以添加多个 返回成功数量
func (pool *RedisDao) SADD(key string, value interface{}) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("SADD", key, value))
	return num, RedisOK(err)
}

func (pool *RedisDao) SMADD(values []interface{}) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("SADD", values...))
	return num, RedisOK(err)
}

// Set 总是成功的
func (pool *RedisDao) Set(key string, value interface{}) (err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	_, err = conn.Do("SET", key, value)
	return RedisOK(err)
}

// 不存在则设置，存在则不设置
func (pool *RedisDao) SetNX(key string, value interface{}) (err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	_, err = conn.Do("SET", key, value, "NX")
	return RedisOK(err)
}
func (pool *RedisDao) SetNX2(key string, value interface{}) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("SETNX", key, value))
	return num, RedisOK(err)
}

// Del 可以删除多个key 返回删除key的num和错误
func (pool *RedisDao) Del(key ...interface{}) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("DEL", key...))
	return num, RedisOK(err)
}

// Get
func (pool *RedisDao) Get(key string) (s string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	s, err = redis.String(conn.Do("GET", key))
	return s, RedisOK(err)
}

// MGET当字段不存在或者 key 不存在时返回nil。
func (pool *RedisDao) MGET(keys []interface{}) (data []string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	data, err = redis.Strings(conn.Do("MGET", keys...))
	return data, RedisOK(err)
}

// MGetInt
func (pool *RedisDao) MGetInt(keys []interface{}) (n []int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	n, err = redis.Ints(conn.Do("MGET", keys...))
	return n, RedisOK(err)
}

// Get
func (pool *RedisDao) GetInt(key string) (n int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	n, err = redis.Int(conn.Do("GET", key))
	return n, RedisOK(err)
}

func (pool *RedisDao) GetInt64(key string) (n int64, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	n, err = redis.Int64(conn.Do("GET", key))
	return n, RedisOK(err)
}

// EXIST
func (pool *RedisDao) EXISTS(key string) (ok bool, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	ok, err = redis.Bool(conn.Do("EXISTS", key))
	return ok, RedisOK(err)
}

// KEYS cz
func (pool *RedisDao) KEYS(pattern string) (keys []string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	keys, err = redis.Strings(conn.Do("KEYS", pattern))
	return keys, RedisOK(err)
}

// SCARD
func (pool *RedisDao) SCARD(key string) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("SCARD", key))
	return num, RedisOK(err)
}

// SPOP 弹出被移除的元素, 当key不存在的时候返回 nil
func (pool *RedisDao) SPOP(key string) (out string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	out, err = redis.String(conn.Do("SPOP", key))
	return out, RedisOK(err)
}

// SRANDMEMBER
func (pool *RedisDao) SRandMember(key string, n int) (s []string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()
	s, err = redis.Strings(conn.Do("SRANDMEMBER", key, n))
	err = RedisOK(err)
	return
}

// SREM
func (pool *RedisDao) SREM(key string, value interface{}) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("SREM", key, value))
	return num, RedisOK(err)
}

// SISMEMBER
func (pool *RedisDao) SISMEMBER(key string, value interface{}) (ok bool, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	ok, err = redis.Bool(conn.Do("SISMEMBER", key, value))
	return ok, RedisOK(err)
}

// SMEMBERS
func (pool *RedisDao) SMEMBERS(key string) (reply []string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	reply, err = redis.Strings(conn.Do("SMEMBERS", key))
	return reply, RedisOK(err)
}

// LPOP
func (pool *RedisDao) LPOP(key string) (out string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	out, err = redis.String(conn.Do("LPOP", key))
	return out, RedisOK(err)
}

// RPOP
func (pool *RedisDao) RPOP(key string) (out string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	out, err = redis.String(conn.Do("RPOP", key))
	return out, RedisOK(err)
}

func (pool *RedisDao) RPush(key string, obj interface{}) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("RPUSH", key, obj))
	return num, err
}

// LPUSH 整型回复: 在 push 操作后的 list 长度。
func (pool *RedisDao) LPUSH(key string, value ...interface{}) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("LPUSH", key, value))
	return num, RedisOK(err)
}

// LINDEX 当 key 位置的值不是一个列表的时候，会返回一个error
func (pool *RedisDao) LINDEX(key string, index int) (out string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	out, err = redis.String(conn.Do("LINDEX", key, index))
	return out, RedisOK(err)
}

func (pool *RedisDao) LRange(key string, start, stop int) (data [][]byte, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	result, err := conn.Do("LRANGE", key, start, stop)
	return redis.ByteSlices(result, err)
}

// HEXISTS
func (pool *RedisDao) HEXISTS(key, field string) (ok bool, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	ok, err = redis.Bool(conn.Do("HEXISTS", key, field))
	return ok, RedisOK(err)
}

// HGET 该字段所关联的值。当字段不存在或者 key 不存在时返回nil。
func (pool *RedisDao) HGET(key, field string) (out string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	out, err = redis.String(conn.Do("HGET", key, field))
	return out, RedisOK(err)
}

// HGET 该字段所关联的值。当字段不存在或者 key 不存在时返回nil。
func (pool *RedisDao) HGETINT(key, field string) (out int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	out, err = redis.Int(conn.Do("HGET", key, field))
	return out, RedisOK(err)
}

// HINCRBY 增值操作执行后的该字段的值。
func (pool *RedisDao) HINCRBY(key, field string, in int) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("HINCRBY", key, field, in))
	return num, RedisOK(err)
}

// HMGETSTRUCT
func (pool *RedisDao) HMGETSTRUCT(key, value interface{}) (err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	v, err := redis.Values(conn.Do("HGETALL", key))
	if err == nil {
		err = redis.ScanStruct(v, value)
	}

	return RedisOK(err)
}

// HMGETMAP
func (pool *RedisDao) HMGETMAP(key string) (map[string]string, error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	m, err := redis.StringMap(conn.Do("HGETALL", key))
	return m, RedisOK(err)
}

func (pool *RedisDao) HMGETINTMAP(key string) (map[string]int, error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	m, err := redis.IntMap(conn.Do("HGETALL", key))
	return m, RedisOK(err)
}

func (pool *RedisDao) HMGETINT64MAP(key string) (map[string]int64, error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	m, err := redis.Int64Map(conn.Do("HGETALL", key))
	return m, RedisOK(err)
}

func (pool *RedisDao) HGETALLMAP(key string) (interface{}, error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	data, err := conn.Do("HGETALL", key)
	return data, RedisOK(err)
}

// set slice to hash
func (pool *RedisDao) HMSETQuestionSlice(key string, data []int, t int) (ok string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	Jsonstr, err := json.Marshal(data)
	if err != nil {
		return
	}
	fieldName := ""
	switch t {
	case 1:
		fieldName = "ModuleContent"
	case 2:
		fieldName = "UserDress"
	case 3:
		fieldName = "UserComm"
	}
	if _, err = conn.Do("HMSET", key, fieldName, Jsonstr); err != nil {
		return
	}
	return
}

// HMSET
func (pool *RedisDao) HMSET(key, value interface{}) (err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	_, err = redis.String(conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(value)...))
	return RedisOK(err)
}

// HMGET
func (pool *RedisDao) HMGET(key, value interface{}) (data []string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	data, err = redis.Strings(conn.Do("HMGET", redis.Args{}.Add(key).AddFlat(value)...))
	return data, RedisOK(err)
}

// HKEYS
func (pool *RedisDao) HKEYS(key string) (data []string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	data, err = redis.Strings(conn.Do("HKEYS", key))
	return data, RedisOK(err)
}

// HMGET
func (pool *RedisDao) HMGET2(key string, feild ...string) (data []string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	data, err = redis.Strings(conn.Do("HMGET", redis.Args{}.Add(key).AddFlat(feild)...))
	return data, RedisOK(err)
}

// HSCAN
// HSET 1如果field是一个新的字段  0如果field原来在map里面已经存在
func (pool *RedisDao) HSET(key, field string, value interface{}) (ok bool, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	ok, err = redis.Bool(conn.Do("HSET", key, field, value))
	return ok, RedisOK(err)
}

// HLEN 哈希集中字段的数量，当 key 指定的哈希集不存在时返回 0
func (pool *RedisDao) HLEN(key string) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("HLEN", key))
	return num, RedisOK(err)
}

// ZREMRANGEBYRANK myzset 0 1  0 -200(保留200名)
func (pool *RedisDao) ZREMRANGEBYRANK(key string, stop int) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("ZREMRANGEBYRANK", key, 0, stop))
	return num, RedisOK(err)
}

// ZADD
func (pool *RedisDao) ZADD(key string, sorce int64, member string) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("ZADD", key, sorce, member))
	return num, RedisOK(err)
}

// ZADD float64
func (pool *RedisDao) ZFADD(key string, sorce float64, member string) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("ZADD", key, sorce, member))
	return num, RedisOK(err)
}

// ZCARD cz
func (pool *RedisDao) ZCARD(key string) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("ZCARD", key))
	return num, RedisOK(err)
}

// ZRANGE cz
func (pool *RedisDao) ZRANGE(key string, start, stop int) (list []string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	list, err = redis.Strings(conn.Do("ZRANGE", key, start, stop))
	return list, RedisOK(err)
}

// ZREVRANGE cz
func (pool *RedisDao) ZREVRANGE(key string, start, stop int) (list []string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	list, err = redis.Strings(conn.Do("ZREVRANGE", key, start, stop))
	return list, RedisOK(err)
}

// ZSCORE cz
func (pool *RedisDao) ZSCORE(key string, member string) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("ZSCORE", key, member))
	return num, RedisOK(err)
}

// ZSCORE cz
func (pool *RedisDao) ZFSCORE(key string, member string) (num float64, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Float64(conn.Do("ZSCORE", key, member))
	return num, RedisOK(err)
}

func (pool *RedisDao) ZREM(key string, member string) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("ZREM", key, member))
	return num, RedisOK(err)
}

// ZREVRANGEBYSCORE 逆序份数  获取的 前N个数据
func (pool *RedisDao) ZREVRANGEBYSCORE(key string, limit int) (list map[string]string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	list, err = redis.StringMap(conn.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "WITHSCORES", "limit", 0, limit))
	return list, RedisOK(err)
}

// ZREVRANGEBYSCORE 逆序份数  获取start len的数据
func (pool *RedisDao) ZREVRANGEBYSCORE2(key string, start, len int) (list map[string]int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	list, err = redis.IntMap(conn.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "WITHSCORES", "limit", start, len))
	return list, RedisOK(err)
}

// ZREVRANGEBYSCORE 逆序份数  获取start len的数据
func (pool *RedisDao) ZREVRANGEBYSCORE3(key string, start, len int) (list map[string]float64, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	list, err = floatMap(conn.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "WITHSCORES", "limit", start, len))
	return list, RedisOK(err)
}

func floatMap(result interface{}, err error) (map[string]float64, error) {
	values, err := redis.Values(result, err)
	if err != nil {
		return nil, err
	}
	if len(values)%2 != 0 {
		return nil, errors.New("redigo: IntMap expects even number of values result")
	}
	m := make(map[string]float64, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].([]byte)
		if !ok {
			return nil, errors.New("redigo: IntMap key not a bulk string value")
		}
		value, err := redis.Float64(values[i+1], nil)
		if err != nil {
			return nil, err
		}
		m[string(key)] = value
	}
	return m, nil
}

// ZREVRANGEBYSCORE 逆序份数  获取的 前N个数据 不要scores
func (pool *RedisDao) GetSearchKeys(key string, limit int) (list []string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	list, err = redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "limit", 0, limit))
	return list, RedisOK(err)
}

// ZREVRANGEBYSCORE 逆序份数  获取的 start,len 不要scores
func (pool *RedisDao) GetSearchKeys2(key string, start, len int) (list []string, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	list, err = redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "limit", start, len))
	return list, RedisOK(err)
}

/**
倒叙获取zset中小于maxScore的len条记录
*/
func (pool *RedisDao) ZrangLessThan(key string, maxScore int64, len int) ([]string, error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	list, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, maxScore, "-inf", "limit", 0, len))
	return list, RedisOK(err)
}

// ZINCRBY +increment  如果没有key 插入
func (pool *RedisDao) ZINCRBY(key string, increment int, member string) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("ZINCRBY", key, increment, member))
	return num, RedisOK(err)
}

// ZRANK 判断一个member 在key中的索引 如果不在 返回nil ,在 返回索引
func (pool *RedisDao) ZRANK(key string, member string) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("ZRANK", key, member))
	return num, RedisOK(err)
}

// ZREVRANK
func (pool *RedisDao) ZREVRANK(key string, member string) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("ZREVRANK", key, member))
	return num, RedisOK(err)
}

// EXPIRE 设置一个key 的过期时间 返回值int 1 如果设置了过期时间 0 如果没有设置过期时间，或者不能设置过期时间
func (pool *RedisDao) EXPIRE(key string, expireTime int) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("EXPIRE", key, expireTime))
	return num, RedisOK(err)
}

// EXPIREAT 设置一个key 的在指定时间过期 返回值：如果生存时间设置成功，返回 1 ;当 key 不存在或没办法设置生存时间，返回 0 。
func (pool *RedisDao) EXPIREAT(key string, expireAtTime int64) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("EXPIREAT", key, expireAtTime))
	return num, RedisOK(err)
}

// SETEX key seconds value
func (pool *RedisDao) SETEX(key string, seconds int, value interface{}) (err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", key, seconds, value)
	return RedisOK(err)
}

// SETEX key seconds value
func (pool *RedisDao) INCR(key string) (err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	_, err = conn.Do("INCR", key)
	return RedisOK(err)
}
func (pool *RedisDao) INCRRET(key string) (num int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	num, err = redis.Int(conn.Do("INCR", key))
	return num, RedisOK(err)
}

func (pool *RedisDao) INCRBY(key string, num int) error {
	conn := pool.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("INCRBY", key, num)
	return RedisOK(err)
}

func (pool *RedisDao) SETBIT(key string, bit, value int) (ret int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	ret, err = redis.Int(conn.Do("SETBIT", key, bit, value))
	return ret, RedisOK(err)
}
func (pool *RedisDao) GETBIT(key string, bit int) (ret int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()

	ret, err = redis.Int(conn.Do("GETBIT", key, bit))
	return ret, RedisOK(err)
}

//set 差集
func (pool *RedisDao) SDIFFSTORE(set1, set2, set3 string) (ret int, err error) {
	conn := pool.Pool.Get()
	defer conn.Close()
	ret, err = redis.Int(conn.Do("SDIFFSTORE", set1, set2, set3))
	return ret, RedisOK(err)
}
