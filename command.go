// Generated by gen.bash.
// DO NOT EDIT THIS FILE DIRECTLY!

package radix

type Command string

const (
	Append           Command = "append"
	Asking           Command = "asking"
	Auth             Command = "auth"
	Bgrewriteaof     Command = "bgrewriteaof"
	Bgsave           Command = "bgsave"
	Blpop            Command = "blpop"
	Brpop            Command = "brpop"
	Brpoplpush       Command = "brpoplpush"
	ClientCmd        Command = "client"
	Cluster          Command = "cluster"
	Config           Command = "config"
	Dbsize           Command = "dbsize"
	Debug            Command = "debug"
	Decr             Command = "decr"
	Decrby           Command = "decrby"
	Del              Command = "del"
	Discard          Command = "discard"
	Dump             Command = "dump"
	Echo             Command = "echo"
	Eval             Command = "eval"
	Exec             Command = "exec"
	Exists           Command = "exists"
	Expire           Command = "expire"
	Expireat         Command = "expireat"
	Flushall         Command = "flushall"
	Flushdb          Command = "flushdb"
	Get              Command = "get"
	Getbit           Command = "getbit"
	Getrange         Command = "getrange"
	Getset           Command = "getset"
	Hdel             Command = "hdel"
	Hexists          Command = "hexists"
	Hget             Command = "hget"
	Hgetall          Command = "hgetall"
	Hincrby          Command = "hincrby"
	Hincrbyfloat     Command = "hincrbyfloat"
	Hkeys            Command = "hkeys"
	Hlen             Command = "hlen"
	Hmget            Command = "hmget"
	Hmset            Command = "hmset"
	Hset             Command = "hset"
	Hsetnx           Command = "hsetnx"
	Hvals            Command = "hvals"
	Incr             Command = "incr"
	Incrby           Command = "incrby"
	Incrbyfloat      Command = "incrbyfloat"
	Info             Command = "info"
	Keys             Command = "keys"
	Lastsave         Command = "lastsave"
	Lindex           Command = "lindex"
	Linsert          Command = "linsert"
	Llen             Command = "llen"
	Lpop             Command = "lpop"
	Lpush            Command = "lpush"
	Lpushx           Command = "lpushx"
	Lrange           Command = "lrange"
	Lrem             Command = "lrem"
	Lset             Command = "lset"
	Ltrim            Command = "ltrim"
	Mget             Command = "mget"
	Migrate          Command = "migrate"
	Monitor          Command = "monitor"
	Move             Command = "move"
	Mset             Command = "mset"
	Msetnx           Command = "msetnx"
	Multi            Command = "multi"
	Object           Command = "object"
	Persist          Command = "persist"
	Pexpire          Command = "pexpire"
	Pexpireat        Command = "pexpireat"
	Ping             Command = "ping"
	Psetex           Command = "psetex"
	Psubscribe       Command = "psubscribe"
	Pttl             Command = "pttl"
	Publish          Command = "publish"
	Punsubscribe     Command = "punsubscribe"
	Randomkey        Command = "randomkey"
	Rename           Command = "rename"
	Renamenx         Command = "renamenx"
	Restore          Command = "restore"
	Rpop             Command = "rpop"
	Rpoplpush        Command = "rpoplpush"
	Rpush            Command = "rpush"
	Rpushx           Command = "rpushx"
	Sadd             Command = "sadd"
	Save             Command = "save"
	Scard            Command = "scard"
	Script           Command = "script"
	Sdiff            Command = "sdiff"
	Sdiffstore       Command = "sdiffstore"
	Select           Command = "select"
	Set              Command = "set"
	Setbit           Command = "setbit"
	Setex            Command = "setex"
	Setnx            Command = "setnx"
	Setrange         Command = "setrange"
	Shutdown         Command = "shutdown"
	Sinter           Command = "sinter"
	Sinterstore      Command = "sinterstore"
	Sismember        Command = "sismember"
	Slaveof          Command = "slaveof"
	Smove            Command = "smove"
	Sort             Command = "sort"
	Spop             Command = "spop"
	Srandmember      Command = "srandmember"
	Srem             Command = "srem"
	Strlen           Command = "strlen"
	Subscribe        Command = "subscribe"
	Sunion           Command = "sunion"
	Sunionstore      Command = "sunionstore"
	Sync             Command = "sync"
	Time             Command = "time"
	Ttl              Command = "ttl"
	Type             Command = "type"
	Unsubscribe      Command = "unsubscribe"
	Unwatch          Command = "unwatch"
	Watch            Command = "watch"
	Zadd             Command = "zadd"
	Zcard            Command = "zcard"
	Zcount           Command = "zcount"
	Zincrby          Command = "zincrby"
	Zinterstore      Command = "zinterstore"
	Zrange           Command = "zrange"
	Zrangebyscore    Command = "zrangebyscore"
	Zrank            Command = "zrank"
	Zrem             Command = "zrem"
	Zremrangebyrank  Command = "zremrangebyrank"
	Zremrangebyscore Command = "zremrangebyscore"
	Zrevrange        Command = "zrevrange"
	Zrevrangebyscore Command = "zrevrangebyscore"
	Zrevrank         Command = "zrevrank"
	Zscore           Command = "zscore"
	Zunionstore      Command = "zunionstore"
	Smembers         Command = "smembers"
)
