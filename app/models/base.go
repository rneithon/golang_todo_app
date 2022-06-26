package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"../../config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)


var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {

	cmd := fmt.Sprintf(
		"%s:%s@(%s)/%s",config.Config.UserName, config.Config.Password, config.Config.Host, config.Config.DbName,
		) 
	
	Db, err = sql.Open(config.Config.SQLDriver, cmd)
	if err != nil{
		log.Fatal(err)
	}
	defer fmt.Println("close")
	
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(255) NOT NULL UNIQUE,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL);`, tableNameUser)
		
	_,err = Db.Exec(cmdU)
	if err != nil {
		log.Fatal("CREATE failed: ",err)
	}
}
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}