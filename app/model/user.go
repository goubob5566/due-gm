package model

import (
	"context"

	"github.com/dobyte/due/cache/redis/v2"
	dbredis "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UUID     string `json:"uuid"`
}
type UUser struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
}

var createUserScript = dbredis.NewScript(`
	local userKey = KEYS[1]
	local uuserKey = KEYS[2]
	
	local userStr = ARGV[1]
	local uuserStr = ARGV[2]
	local username = ARGV[3]
	local uuid = ARGV[4]

	local exists = redis.call("HEXISTS", userKey, username)
	if exists == 0 then
		redis.call("HSET", userKey, username, userStr)
		redis.call("HSET", uuserKey, uuid, uuserStr)
		return 1
	end
	
	return 0
`)

func GetUserByUsername(username string) (*User, error) {
	c := redis.NewCache()
	r := c.Client().(dbredis.UniversalClient)

	user := &User{}
	item, err := r.HGet(context.Background(), KeysUser, username).Result()
	if err != nil && err != dbredis.Nil {
		return nil, err
	}
	if err == dbredis.Nil {
		return nil, nil
	}
	err = msgpack.Unmarshal([]byte(item), &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func RegisterUser(username string, password []byte) (*UUser, error) {
	c := redis.NewCache()
	r := c.Client().(dbredis.UniversalClient)

	uuid := uuid.New().String()
	user := &User{
		Username: username,
		Password: string(password),
		UUID:     uuid,
	}
	uuser := &UUser{
		UUID:     uuid,
		Username: username,
	}
	userStr, _ := msgpack.Marshal(&user)
	uuserStr, _ := msgpack.Marshal(&uuser)

	_, err := createUserScript.Run(context.Background(), r,
		[]string{KeysUser, KeysUUser},
		[]interface{}{userStr, uuserStr, username, uuid}).Int()
	if err != nil {
		return nil, err
	}

	return uuser, nil
}
