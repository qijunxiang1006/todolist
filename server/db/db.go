package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iris-contrib/go.uuid"
	
	`todolist/model`
)

const (
	userName = "root"
	passWord = "123456"
	host     = "172.10.22.2"
	port     = "3306"
	network  = "tcp"
	database = "todolist"
	table    = "todolist"
	id       = "id"
	uid      = "uuid"
	value    = "value"
	status   = "status"
)
const (
	insert       = "INSERT INTO %v (%v,%v,%v) VALUES ('%v','%v',%v)"
	update       = "UPDATE %v SET %v=%v WHERE %v='%v' && %v=%v"
	delete       = "DELETE FROM %v WHERE %v='%v' && %v=%v"
	search       = "SELECT * FROM %v WHERE %v='%v'"
	lastInsertID = "SELECT LAST_INSERT_ID()"
)

type db struct {
	db *sql.DB
}

func (d *db) getLastInsertID() (int, error) {
	r := d.db.QueryRow(lastInsertID)
	id := 0
	err := r.Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (d *db) Insert(uuid uuid.UUID, data *model.DBListItem) (int, error) {
	fmt.Println(fmt.Sprintf(insert, table, uid, value, status, uuid, data.Value, data.Status))
	_, err := d.db.Query(fmt.Sprintf(insert, table, uid, value, status, uuid.String(), data.Value, data.Status))
	if err != nil {
		return 0, err
	}
	id, err := d.getLastInsertID()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *db) Delete(uuid uuid.UUID, ID int) error {
	fmt.Println(fmt.Sprintf(delete, table, uid, uuid.String(), id, ID))
	_, err := d.db.Query(fmt.Sprintf(delete, table, uid, uuid.String(), id, ID))
	if err != nil {
		return err
	}
	return nil
}

func (d *db) Update(uuid uuid.UUID, data *model.DBListItem) error {
	fmt.Println(fmt.Sprintf(update, table, status, data.Status, uid, uuid, id, data.ID))
	_, err := d.db.Query(fmt.Sprintf(update, table, status, data.Status, uid, uuid.String(), id, data.ID))
	if err != nil {
		return err
	}
	return nil
}

func (d *db) Get(uuid uuid.UUID) (*model.DBData, error) {
	fmt.Println(fmt.Sprintf(search, table, uid, uuid))
	r, err := d.db.Query(fmt.Sprintf(search, table, uid, uuid.String()))
	if err != nil {
		return nil, err
	}
	data := new(model.DBData)
	data.List = make([]*model.DBListItem, 0, 10)
	for r.Next() {
		item := new(model.DBListItem)
		err := r.Scan(&item.ID, &data.UUID, &item.Value, &item.Status)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		data.List = append(data.List, item)
	}
	fmt.Println(*data)
	return data, err
}

func NewDB() (DB, error) {
	db := new(db)
	newDB, err := sql.Open("mysql", fmt.Sprintf("%v:%v@%v(%v:%v)/%v", userName, passWord, network, host, port, database))
	if err != nil {
		return nil, err
	}
	db.db = newDB
	err = db.db.Ping()
	if err != nil {
		return nil, err
	}
	err = db.checkDB()
	if err != nil {
		return nil, err
	}
	return db, nil
}
func (d *db) checkDB() error {
	_, err := d.db.Query(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", table))
	if err != nil {
		return err
	}
	_, err = d.db.Query(fmt.Sprintf("USE %v", table))
	if err != nil {
		return err
	}
	_, err = d.db.Query("CREATE TABLE IF NOT EXISTS todolist (" +
		"id INT UNSIGNED NOT NULL AUTO_INCREMENT," +
		"uuid CHAR(36) NOT NULL," +
		"value VARCHAR(255) NOT NULL," +
		"status BOOLEAN NOT NULL," +
		"PRIMARY KEY (id)" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8")
	if err != nil {
		return err
	}
	return nil
}

type DB interface {
	Insert(uuid.UUID, *model.DBListItem) (int, error)
	Delete(uuid.UUID, int) error
	Update(uuid.UUID, *model.DBListItem) error
	Get(uuid.UUID) (*model.DBData, error)
}
