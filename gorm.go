package gormx

import (
	"fmt"
	"github.com/ginx-contribs/dbx"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Options = dbx.Options

func Open(options Options, gormOpts ...gorm.Option) (*gorm.DB, error) {

	var dialector gorm.Dialector
	switch options.Driver {
	case dbx.Mysql, dbx.Postgres:
		// open with *sql.dbx
		sqldb, err := dbx.Open(options)
		if err != nil {
			return nil, err
		}
		if options.Driver == dbx.Mysql {
			dialector = mysql.New(mysql.Config{Conn: sqldb})
		} else {
			dialector = postgres.New(postgres.Config{Conn: sqldb})
		}
	case dbx.Sqlite:
		dialector = sqlite.Open(dbx.SQLiteDsn(options))
	case dbx.Sqlserver:
		dialector = sqlserver.Open(dbx.SQLServerDsn(options))
	default:
		return nil, fmt.Errorf("unsupported driver %s", options.Driver)
	}

	gormdb, err := gorm.Open(dialector, gormOpts...)
	if err != nil {
		return nil, err
	}

	return gormdb, nil
}
