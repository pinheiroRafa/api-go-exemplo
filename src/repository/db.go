package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"time"

	custom_error "main/src/entities/custom_error"

	_ "github.com/lib/pq"
	"golang.org/x/exp/slices"
)

var db *sql.DB
var normalError error

func OpenDB(setLimits bool) (*sql.DB, error) {
	fmt.Println("abrindo base de dados")
	var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=UTC",
		os.Getenv("HOST_DB"),
		os.Getenv("PORT_DB"),
		os.Getenv("USER_DB"),
		os.Getenv("PASS_DB"),
		os.Getenv("NAME_DB"))
	const PostgresDriver = "postgres"
	db, normalError = sql.Open(PostgresDriver, DataSourceName)
	if normalError != nil {
		return nil, normalError
	}
	if setLimits {
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	normalError = db.PingContext(ctx)
	if normalError != nil {
		return nil, normalError
	}

	return db, nil
}

func Exec[T any](query string, args ...any) (_ bool, err *custom_error.CustomError) {
	var result sql.Result
	result, normalError = db.Exec(query, args...)
	if normalError != nil {
		err := custom_error.New(normalError.Error())
		return false, &err
	}
	var r int64
	r, normalError = result.RowsAffected()
	if normalError != nil {
		err := custom_error.New(normalError.Error())
		return false, &err
	}
	return r > 0, nil
}

func Select[T any](query string, args ...any) (_ []T, err *custom_error.CustomError) {
	list := make([]T, 0)
	sqlStatement, normalError := db.Query(query, args...)
	if normalError != nil {
		err := custom_error.New(normalError.Error())
		return nil, &err
	}
	for sqlStatement.Next() {
		var article T
		s := reflect.ValueOf(&article).Elem()
		t := reflect.TypeOf(&article).Elem()
		var cols, _ = sqlStatement.Columns()
		numCols := len(cols) //s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < s.NumField(); i++ {
			field := s.Field(i)
			idx := slices.IndexFunc(cols, func(c string) bool { return c == t.Field(i).Tag.Get("json") })
			if idx >= 0 {
				columns[i] = field.Addr().Interface()
			}
		}

		normalError := sqlStatement.Scan(columns...)
		if normalError != nil {
			err := custom_error.New(normalError.Error())
			return nil, &err
		}
		list = append(list, article)
	}
	return list, nil
}
