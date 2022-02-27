package db

import (
    "databse/sql"
    "github.com/mattn/go-sqlite3"
)

var sqliteDB *sql.DB = nil;

func Start() error {
	s := setting.Get();
    if s == nil  {
        panic("sqlite3 path have not been set!")
    }
    // for sqlite3 DB
    dbPath := s.SqliteDB.Path;
    if !utils.PathExists(dbPath) {
        panic("sqlite3 file path is not exist!")
    }
    sqliteDB, err := sql.Open("sqlite3", dbPath)
    if err != nil {
    	return err;
    }
    logger.Info("Sqite3 DB Open ==>")
    return nil
}

func Stop() {
	logger.Info("Sqite3 DB Close ==>")
	if sqliteDB != nil {
		sqliteDB.Close();
	}
}