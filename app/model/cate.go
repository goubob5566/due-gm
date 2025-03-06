package model

import (
	"context"

	"github.com/dobyte/due/cache/redis/v2"
	dbredis "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack"
)

type Cate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func CreateCate(name string) (*Cate, error) {

	c := redis.NewCache()
	r := c.Client().(dbredis.UniversalClient)

	uuid := uuid.New().String()
	cate := &Cate{
		ID:   uuid,
		Name: name,
	}
	cateStr, _ := msgpack.Marshal(&cate)

	_, err := r.HSet(context.Background(), KeysCate, uuid, cateStr).Result()
	if err != nil {
		return nil, err
	}

	return cate, nil
}

func ReadCateOne(id string) (*Cate, error) {

	c := redis.NewCache()
	r := c.Client().(dbredis.UniversalClient)

	cate := &Cate{}
	cateStr, err := r.HGet(context.Background(), KeysCate, id).Result()
	if err != nil {
		return nil, err
	}
	if err == dbredis.Nil {
		return nil, nil
	}
	err = msgpack.Unmarshal([]byte(cateStr), &cate)
	if err != nil {
		return nil, err
	}

	return cate, nil
}

func ReadCateAll() ([]*Cate, error) {
	c := redis.NewCache()
	r := c.Client().(dbredis.UniversalClient)

	cates := make([]*Cate, 0)
	cateMap, err := r.HGetAll(context.Background(), KeysCate).Result()
	if err != nil {
		return nil, err
	}

	for _, str := range cateMap {
		cate := &Cate{}
		err = msgpack.Unmarshal([]byte(str), cate)
		if err != nil {
			continue
		}
		cates = append(cates, cate)
	}

	return cates, nil
}

func UpdateCate(id string, name string) (*Cate, error) {
	c := redis.NewCache()
	r := c.Client().(dbredis.UniversalClient)

	cate := &Cate{
		ID:   id,
		Name: name,
	}
	cateStr, _ := msgpack.Marshal(&cate)

	_, err := r.HSet(context.Background(), KeysCate, id, cateStr).Result()
	if err != nil {
		return nil, err
	}

	return cate, nil
}

func DeleteCate(id string) error {
	c := redis.NewCache()
	r := c.Client().(dbredis.UniversalClient)

	_, err := r.HDel(context.Background(), KeysCate, id).Result()
	if err != nil {
		return err
	}

	am, err := NewArticleModel()
	if err != nil {
		return err
	}
	for _, article := range am.artile {
		if article.CateID == id {
			am.DeleteArticle(article.ID)
		}
	}

	return nil
}
