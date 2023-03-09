package repos

import (
	"database/sql"
	"slimlink/core/entities"
	"slimlink/core/exceptions"
	"slimlink/core/ports"
	"slimlink/infrastructure/data"
)

type LinkMySqlRepo struct {
	mySqlDB *data.MySqlDB
}

func NewLinkMySqlRepo(mySqlDB *data.MySqlDB) ports.LinkRepo {
	return &LinkMySqlRepo{mySqlDB}
}

func (linkMySqlRepo *LinkMySqlRepo) Add(link *entities.Link) error {
	_, err := linkMySqlRepo.mySqlDB.Exec("INSERT INTO `links` VALUES (?, ?)", link.ID, link.Url)
	return err
}

func (linkMySqlRepo *LinkMySqlRepo) GetByID(id string) (*entities.Link, error) {
	var link entities.Link
	err := linkMySqlRepo.mySqlDB.QueryRow("SELECT * FROM `links` WHERE `id` = ?", id).Scan(&link.ID, &link.Url)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exceptions.NewAppError[*exceptions.NotFoundError]("no rows found")
		}
		return nil, err
	}
	return &link, nil
}
