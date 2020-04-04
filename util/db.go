package util

import (
	gopg "github.com/go-pg/pg"
)

func NewGoPgDb(dataSourceName string) (*gopg.DB, error) {
	options, err := gopg.ParseURL(dataSourceName)
	if err != nil {
		return nil, err
	}

	//check if connection is up
	GoPgDB := gopg.Connect(options) //this is only structured this way for testing gopg
	_, err = GoPgDB.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return GoPgDB, nil
}
