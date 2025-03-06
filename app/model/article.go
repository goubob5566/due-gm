package model

import (
	"context"

	"github.com/dobyte/due/cache/redis/v2"
	dbredis "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack"
)

type Article struct {
	ID     string `json:"id"`
	CateID string `json:"cateid"`
	Title  string `json:"title"`
}

type ArticleModel struct {
	artile map[string]*Article
}

func (am *ArticleModel) GetArticle() map[string]*Article {
	return am.artile
}

func (am *ArticleModel) CreateArticle(cateid string, title string) (*Article, error) {

	article := &Article{
		ID:     uuid.New().String(),
		CateID: cateid,
		Title:  title,
	}

	am.artile[article.ID] = article

	articleStr, _ := msgpack.Marshal(&article)

	c := redis.NewCache()
	r := c.Client().(dbredis.UniversalClient)

	_, err := r.HSet(context.Background(), KeysArticle, article.ID, articleStr).Result()
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (am *ArticleModel) UpdateArticle(id string, cateid string, title string) (*Article, error) {

	article := &Article{
		ID:     id,
		CateID: cateid,
		Title:  title,
	}

	am.artile[article.ID] = article

	articleStr, _ := msgpack.Marshal(&article)

	c := redis.NewCache()
	r := c.Client().(dbredis.UniversalClient)

	_, err := r.HSet(context.Background(), KeysArticle, article.ID, articleStr).Result()
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (am *ArticleModel) DeleteArticle(id string) error {
	delete(am.artile, id)
	c := redis.NewCache()
	r := c.Client().(dbredis.UniversalClient)

	_, err := r.HDel(context.Background(), KeysArticle, id).Result()
	if err != nil {
		return err
	}

	return nil
}

func NewArticleModel() (*ArticleModel, error) {
	am := &ArticleModel{
		artile: make(map[string]*Article),
	}

	c := redis.NewCache()
	r := c.Client().(dbredis.UniversalClient)

	articleMap, err := r.HGetAll(context.Background(), KeysArticle).Result()
	if err != nil {
		return nil, err
	}

	for _, str := range articleMap {
		article := &Article{}
		err = msgpack.Unmarshal([]byte(str), article)
		if err != nil {
			continue
		}
		am.artile[article.ID] = article
	}
	return am, nil
}
