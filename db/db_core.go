package db

import (
	dbManager "database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb" // sql server
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"          // postgreSQL
	_ "github.com/sijms/go-ora/v2" // oracle
	"io"
	_ "modernc.org/sqlite"
	"strings"
)

const (
	dbSqlite    = "sqlite"
	dbMysql     = "mysql"
	dbOracle    = "oracle"
	dbSqlServer = "sqlserver"
	dbPostgres  = "postgres"
)

func Pg(username, password, url string) *DataSource {
	// e.g. postgres://postgres:12345678@192.168.8.200:5432/douyin?sslmode=disable
	source := fmt.Sprintf("postgres://%v:%v@%v", username, password, url)
	return connDB(dbPostgres, source)
}

func Sqlserver(username, password, url string) *DataSource {
	seperatorIndex := strings.Index(url, "/")
	var host, port, dbName string
	if seperatorIndex < 0 || seperatorIndex == len(url)-1 {
		runtimeExcption("error: database name is empty!")
	} else {
		netAddress := url[:seperatorIndex]
		colonIndex := strings.Index(netAddress, ":")
		if colonIndex < 0 || colonIndex == len(netAddress)-1 {
			runtimeExcption("host address format is error!")
		} else {
			host = netAddress[:colonIndex]
			port = netAddress[colonIndex+1:]
		}
		dbName = url[seperatorIndex+1:]
	}
	// e.g. "server=192.168.1.103;port=1433;database=STG;user id=SA;password=root@123"
	sourceName := fmt.Sprintf("server=%v;port=%v;database=%v;user id=%v;password=%v", host, port, dbName, username, password)
	return connDB(dbSqlServer, sourceName)
}

func Oracle(username, password, url string) *DataSource {
	seperatorIndex := strings.Index(url, "/")
	var netAddress, uri string
	if seperatorIndex < 0 {
		netAddress = url
		uri = "?charset=utf8"
	} else {
		netAddress = url[:seperatorIndex]
		uri = url[seperatorIndex:]
	}
	// e.g. oracle://user:pass@server/service_name
	sourceName := fmt.Sprintf("oracle://%v:%v@%v/%v", username, password, netAddress, uri)
	return connDB(dbOracle, sourceName)
}

func Mysql(username, password, url string) *DataSource {
	seperatorIndex := strings.Index(url, "/")
	var netAddress, uri string
	if seperatorIndex < 0 {
		netAddress = url
		uri = "?charset=utf8"
	} else {
		netAddress = url[:seperatorIndex]
		uri = url[seperatorIndex:]
	}
	// e.g. root:root@tcp(192.168.1.103:3306)/tx?charset=utf8
	sourceName := fmt.Sprintf("%v:%v@tcp(%v)%v", username, password, netAddress, uri)
	return connDB(dbMysql, sourceName)
}

func Sqlite(dbName string) *DataSource {
	return connDB(dbSqlite, dbName)
}

func connDB(driverName, sourceName string) *DataSource {
	db, err := dbManager.Open(driverName, sourceName)
	assert(err != nil, "failed to build db connection:", err)

	ds := &DataSource{driverName,
		sourceName,
		db}
	return ds
}

// 数据库连接对象
type DataSource struct {
	driver string
	source string
	db     *dbManager.DB
}

func (ds *DataSource) Exec(args ...any) int64 {
	sql, vals := ds.parseArgs("update", args)

	execResult, err := ds.db.Exec(sql, vals...)
	assert(err != nil, "failed to execute sql:", sql, err)

	affected, err := execResult.RowsAffected()
	assert(err != nil, "failed to update:", sql, err)
	return affected
}

func (ds *DataSource) Insert(args ...any) any {
	sql, vals := ds.parseArgs("insert", args)

	stmt, err := ds.db.Prepare(sql)
	assert(err != nil, "failed to prepare:", err)
	defer cl(stmt)

	execResult, err := stmt.Exec(vals...)
	assert(err != nil, "failed to insert:", err)

	id, err := execResult.LastInsertId()
	assert(err != nil, "failed to insert:", err)
	return id
}

func (ds *DataSource) Update(args []any) int64 {
	sql, vals := ds.parseArgs("update", args)

	stmt, err := ds.db.Prepare(sql)
	assert(err != nil, "failed to prepare:", err)
	defer cl(stmt)

	execResult, err := stmt.Exec(vals...)
	assert(err != nil, "failed to update:", err)

	affected, err := execResult.RowsAffected()
	assert(err != nil, "failed to update:", err)
	return affected
}

// 记录数统计
func (ds *DataSource) Count(obj string) int {
	sql := "select count(*) from " + obj
	rows, err := ds.db.Query(sql)
	assert(err != nil, "failed to count:", err)
	defer cl(rows)

	for rows.Next() {
		var res int
		err = rows.Scan(&res)
		assert(err != nil, "failed to extract count() result:", err)

		return res
	}
	return 0
}

// 获取单个值
func (ds *DataSource) GetValue(args []any) any {
	sql, vals := ds.parseArgs("getValue", args)

	rows, err := ds.db.Query(sql, vals...)
	assert(err != nil, "failed to query data:", err)
	defer cl(rows)

	colTypes, _ := rows.ColumnTypes()
	colCount := len(colTypes)
	assert(colCount > 1, "sql query return over more field value")

	var value any
	rowCount := 1
	for rows.Next() {
		assert(rowCount > 1, "sql query return over more row data")

		err = rows.Scan(&value)
		assert(err != nil, "failed to extract field value:", err)

		value = ds.getFieldValue(value, colTypes[0])

		rowCount++
	}
	return value
}
func (ds *DataSource) Val(args []any) any {
	return ds.GetValue(args)
}

func (ds *DataSource) GetRow(args []any) map[string]any {
	sql, vals := ds.parseArgs("getRow", args)

	rows, err := ds.db.Query(sql, vals...)
	assert(err != nil, "failed to query data:", err)
	defer cl(rows)

	colTypes, _ := rows.ColumnTypes()
	colCount := len(colTypes)
	var list []map[string]any
	rowCount := 1
	for rows.Next() {
		assert(rowCount > 1, "sql query return over more row data")

		valueContainers := getValueContainers(colCount)
		err = rows.Scan(valueContainers...)
		assert(err != nil, "failed to extract field value:", err)

		row := make(map[string]any)
		for i, colType := range colTypes {
			colName := colType.Name()
			tmp := *valueContainers[i].(*any)

			row[colName] = ds.getFieldValue(tmp, colType)
		}
		list = append(list, row)
		rowCount++
	}
	if len(list) > 0 {
		return list[0]
	} else {
		return nil
	}
}
func (ds *DataSource) Row(args []any) map[string]any {
	return ds.GetRow(args)
}

func (ds *DataSource) GetRows(args ...any) []any {
	sql, vals := ds.parseArgs("getRows", args)

	rows, err := ds.db.Query(sql, vals...)
	assert(err != nil, "failed to query data:", err)
	defer cl(rows)

	colTypes, _ := rows.ColumnTypes()
	colCount := len(colTypes)
	var list []any
	for rows.Next() {
		valueContainers := getValueContainers(colCount)
		err = rows.Scan(valueContainers...)
		assert(err != nil, "failed to extract field value:", err)

		row := make(map[string]any)
		for i, colType := range colTypes {
			colName := colType.Name()
			tmp := *valueContainers[i].(*any)

			row[colName] = ds.getFieldValue(tmp, colType)
		}
		list = append(list, row)
	}
	return list
}
func (ds *DataSource) Rows(args []any) []any {
	return ds.GetRows(args)
}

func (ds *DataSource) parseArgs(methodName string, args []any) (sql string, values []any) {
	assert(len(args) < 1, fmt.Sprintf("method db.%v() must has one parameters.", methodName))
	sql, ok := args[0].(string)
	sql = ds.parseSqlServerSql(sql)
	sql = ds.parsePostgresSql(sql)
	sql = ds.parseOracleSql(sql)
	assert(!ok, fmt.Sprintf("method db.%v() the first parameter must be string type.", methodName))
	return sql, args[1:]
}

func (ds *DataSource) parseOracleSql(sql string) string {
	if ds.driver != dbOracle {
		return sql
	}

	argsCount := 1
	var chs []rune
	for _, ch := range sql {
		if ch == '?' {
			chs = append(chs, ':')
			chs = append(chs, intToRunes(argsCount)...)
			argsCount++
		} else {
			chs = append(chs, ch)
		}
	}
	return string(chs)
}

func (ds *DataSource) parseSqlServerSql(sql string) string {
	if ds.driver != dbSqlServer {
		return sql
	}

	argsCount := 1
	var chs []rune
	for _, ch := range sql {
		if ch == '?' {
			chs = append(chs, '@')
			chs = append(chs, 'p')
			chs = append(chs, intToRunes(argsCount)...)
			argsCount++
		} else {
			chs = append(chs, ch)
		}
	}
	return string(chs)
}

func (ds *DataSource) parsePostgresSql(sql string) string {
	if ds.driver != dbPostgres {
		return sql
	}

	argsCount := 1
	var chs []rune
	for _, ch := range sql {
		if ch == '?' {
			chs = append(chs, '$')
			chs = append(chs, intToRunes(argsCount)...)
			argsCount++
		} else {
			chs = append(chs, ch)
		}
	}
	return string(chs)
}

func getValueContainers(size int) []any {
	var valueContainers []any
	for i := 0; i < size; i++ {
		var container any
		valueContainers = append(valueContainers, &container)
	}
	return valueContainers
}

// 释放资源
func cl(obj io.Closer) {
	err := obj.Close()
	if err != nil {
		runtimeExcption(err)
	}
}
