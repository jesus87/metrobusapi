package sqlxvendor

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jesus87/metrobusapi/src/infrastructure/orm"
	"github.com/jmoiron/sqlx"
)

//SqlxVendor type for persist database
type SqlxVendor struct {
	_engine           string
	_connectionString string
}

var _ orm.OrmManager = (*SqlxVendor)(nil)

//HandleConnection handler for openning connections
func (vendor *SqlxVendor) HandleConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect(vendor._engine, vendor._connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//Select method for select information
func (vendor *SqlxVendor) Select(dest interface{}, query string, arg interface{}) error {
	db, err := vendor.HandleConnection()
	if err != nil {
		return err
	}
	rows, err := db.NamedQuery(query, arg)
	if err != nil {
		return err
	}

	err = sqlx.StructScan(rows, dest)
	if err != nil {
		return err
	}

	return nil
}

//Save method for saving information on database
func (vendor *SqlxVendor) Save(query string, arg interface{}) error {
	db, err := vendor.HandleConnection()
	if err != nil {
		return err
	}
	_, err = db.NamedExec(query, arg)
	if err != nil {
		return err
	}

	return nil
}
//NewSqlxVendor instance for sql type
func NewSqlxVendor(engine string, connectionString string) *SqlxVendor {

	return &SqlxVendor{
		_engine:           engine,
		_connectionString: connectionString,
	}
}
