package repos

import (
	"slimlink/core/entities"
	"slimlink/core/exceptions"
	"slimlink/core/ports"
	"slimlink/infrastructure/data"

	"github.com/redis/go-redis/v9"
)

type linkRedisRepo struct {
	db *data.RedisDB
}

func NewLinkRedisRepo(redisDB *data.RedisDB) ports.LinkRepo {
	return &linkRedisRepo{redisDB}
}

func (r *linkRedisRepo) Add(link *entities.Link) error {
	return r.db.Set(link.ID, link.Url)
}

func (r *linkRedisRepo) GetByID(id string) (*entities.Link, error) {
	url, err := r.db.Get(id)
	if err != nil {
		if err == redis.Nil {
			return nil, exceptions.NewAppError[*exceptions.NotFoundError]("key does not exist")
		}
		return nil, err
	}
	return &entities.Link{ID: id, Url: url}, nil
}
