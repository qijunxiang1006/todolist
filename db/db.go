package db

import (
	`database/sql`
	`fmt`
	
	_ `github.com/go-sql-driver/mysql`
	
	`todolist/model`
)

const(
	userName = "root"
	passWord = "123456"
	host="localhost"
	port="3306"
	network="tcp"
	database="todolist"
	table="todolist"
	id="id"
	createTime="createtime"
	updateTime="updatetime"

)
const(
	INSERT="INSERT INTO %v (%v,%v,%v) VALUES (%v,CURRENT_TIMESTAMP(),CURRENT_TIMESTAMP())"
	UPDATE="UPDATE %v SET %v=CURRENT_TIMESTAMP() WHERE %v=%v"
	DELETE="DELETE FROM %v WHERE %v=%v"
	GET="SELECT %v, %v, %v FROM %v WHERE %v=%v"
)
type db struct {
	db *sql.DB
}

func (d *db) Insert(data *model.DBData)error {
	fmt.Println(fmt.Sprintf(INSERT,table,id,createTime,updateTime,data.ID))
	_,err:=d.db.Query(fmt.Sprintf(INSERT,table,id,createTime,updateTime,data.ID))
	if err!=nil{
		return err
	}
	return nil
}

func (d *db) Delete(mac int64)error {
	_,err:=d.db.Query(fmt.Sprintf(DELETE,table,id,mac))
	if err!=nil{
		return err
	}
	return nil
}

func (d *db) Update(data *model.DBData)error {
	fmt.Println(fmt.Sprintf(UPDATE,table,updateTime,id,data.ID))
	_,err:=d.db.Query(fmt.Sprintf(UPDATE,table,updateTime,id,data.ID))
	if err!=nil{
		return err
	}
	return nil
}

func (d *db) Get(mac int64)(*model.DBData,error) {
	fmt.Println(fmt.Sprintf(GET,id,createTime,updateTime,table,id,mac))
	r:=d.db.QueryRow(fmt.Sprintf(GET,id,createTime,updateTime,table,id,mac))
	data:=new(model.DBData)
	err:=r.Scan(&data.ID,&data.CreateTime,&data.UpdateTime)
	fmt.Println(*data)
	return data,err
}

func NewDB() (DB,error){
	db:=new(db)
	newDB,err := sql.Open("mysql",fmt.Sprintf("%v:%v@%v(%v:%v)/%v",userName,passWord,network,host,port,database))
	if err!=nil{
		return nil,err
	}
	db.db=newDB
	err=db.db.Ping();if err!=nil{
		return nil,err
	}
	err=db.checkDB()
	if err!=nil{
		return nil,err
	}
	return db,nil
}
func (d *db)checkDB()error{
	_,err:=d.db.Query("CREATE DATABASE IF NOT EXISTS todolist")
	if err!=nil{
		return err
	}
	_,err=d.db.Query("USE todolist")
	if err!=nil{
		return err
	}
	_,err=d.db.Query("CREATE TABLE IF NOT EXISTS todolist (id INT UNSIGNED NOT NULL," +
		"createtime TIMESTAMP NOT NULL," +
		"updatetime TIMESTAMP NOT NULL," +
		"PRIMARY KEY (id)" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8")
	if err!=nil{
		return err
	}
	return nil
}
type DB interface {
	Insert(data *model.DBData)error
	Delete(mac int64)error
	Update(data *model.DBData)error
	Get(mac int64)(*model.DBData,error)
}