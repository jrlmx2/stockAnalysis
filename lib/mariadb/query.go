package mariadb

type queryable interface {
	Table() string
	Data() string
}
