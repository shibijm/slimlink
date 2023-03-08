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

func NewLinkMySqlRepo(mySqlDB *data.MySqlDB) (ports.LinkRepo, error) {
	_, err := mySqlDB.Exec("CREATE TABLE IF NOT EXISTS `links` (`id` VARCHAR(64) NOT NULL, `url` VARCHAR(2048) NOT NULL, PRIMARY KEY (`id`))")
	if err != nil {
		return nil, err
	}
	return &LinkMySqlRepo{mySqlDB}, nil
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
