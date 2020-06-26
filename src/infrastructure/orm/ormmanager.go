package orm

type OrmManager interface {
	Select(dest interface{}, query string, arg interface{}) error
	Save(query string, arg interface{}) error
}
