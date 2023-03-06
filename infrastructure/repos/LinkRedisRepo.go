package repos

import (
	"errors"
	"slimlink/core/entities"
	"slimlink/core/exceptions"
	"slimlink/core/ports"
	"slimlink/infrastructure/data"

	"github.com/redis/go-redis/v9"
)

type LinkRedisRepo struct {
	redisDB *data.RedisDB
}

func NewLinkRedisRepo(redisDB *data.RedisDB) ports.LinkRepo {
	return &LinkRedisRepo{redisDB}
}

func (linkRedisRepo *LinkRedisRepo) Add(link *entities.Link) error {
	return linkRedisRepo.redisDB.Set(link.ID, link.Url)
}

func (linkRedisRepo *LinkRedisRepo) GetByID(id string) (*entities.Link, error) {
	url, err := linkRedisRepo.redisDB.Get(id)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, exceptions.NewAppError[*exceptions.NotFoundError]("key does not exist")
		}
		return nil, err
	}
	return &entities.Link{ID: id, Url: url}, nil
}
