package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"hexarch/utils"
	"time"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.WithField("func", utils.GetCallerInfo()).Fatalf("db connection failure: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.WithField("func", utils.GetCallerInfo()).Fatalf("db ping failure: %v", err)
		return nil, err
	}

	return &Adapter{
		db: db,
	}, nil
}

func (da Adapter) CloseDBConnection() {
	err := da.db.Close()
	if err != nil {
		log.WithField("func", utils.GetCallerInfo()).Fatalf("db close failure: %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation).ToSql()

	if err != nil {
		log.WithField("func", utils.GetCallerInfo()).Debugf("sq query generation failure: %v", err)
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		log.WithField("func", utils.GetCallerInfo()).Debugf("db exec failure: %v", err)
		return err
	}

	return nil
}
