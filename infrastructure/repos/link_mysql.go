package repos

import (
	"database/sql"
	"slimlink/core/entities"
	"slimlink/core/exceptions"
	"slimlink/core/ports"
	"slimlink/infrastructure/data"
)

type linkMySqlRepo struct {
	db *data.MySqlDB
}

func NewLinkMySqlRepo(mySqlDB *data.MySqlDB) ports.LinkRepo {
	return &linkMySqlRepo{mySqlDB}
}

func (r *linkMySqlRepo) Add(link *entities.Link) error {
	_, err := r.db.Exec("INSERT INTO `links` (`id`, `url`) VALUES (?, ?)", link.ID, link.Url)
	return err
}

func (r *linkMySqlRepo) GetByID(id string) (*entities.Link, error) {
	link := entities.Link{ID: id}
	err := r.db.QueryRow("SELECT `url` FROM `links` WHERE `id` = ?", id).Scan(&link.Url)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exceptions.NewAppError[*exceptions.NotFoundError]("no rows found")
		}
		return nil, err
	}
	return &link, nil
}
