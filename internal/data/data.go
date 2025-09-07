package data

import (
	"database/sql"
	"my-project/internal/conf"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserRepo)

// Data .
type Data struct {
	DB *sql.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(logger)
	db, err := sql.Open("mysql", c.Database.Source)
	if err != nil {
		helper.Errorf("Failed to connect to DB: %v", err)
		return nil, nil, err
	}
	cleanup := func() {
		helper.Info("closing the data resources")
		db.Close()
	}
	return &Data{DB: db}, cleanup, nil
}
