package controller

import (
	`fmt`
	
	uuid "github.com/iris-contrib/go.uuid"
	
	`todolist/db`
	`todolist/model`
)

type service struct {
	db db.DB
}

func (s *service) GetList(uid uuid.UUID) (*model.DBData, error) {
	if s.db==nil{
		fmt.Println("db is nil,try to init db")
		var err error
		s.db,err = db.NewDB()
		if err!=nil{
			fmt.Println("init db failed,",err.Error())
			return nil,err
		}
	}
	return s.db.Get(uid)
}

func (s *service) PutList(uuid uuid.UUID, data *model.DBListItem) error {
	if s.db==nil{
		fmt.Println("db is nil,try to init db")
		
		var err error
		s.db,err = db.NewDB()
		if err!=nil{
			fmt.Println("init db failed,",err.Error())
			return err
		}
	}
	return s.db.Update(uuid, data)
}

func (s *service) PostList(uuid uuid.UUID, data *model.DBListItem) (int, error) {
	if s.db==nil{
		fmt.Println("db is nil,try to init db")
		
		var err error
		s.db,err = db.NewDB()
		if err!=nil{
			fmt.Println("init db failed,",err.Error())
			return 0,err
		}
	}
	return s.db.Insert(uuid, data)
}

func (s *service) DeleteList(uuid uuid.UUID, ID int) error {
	if s.db==nil{
		fmt.Println("db is nil,try to init db")
		var err error
		s.db,err = db.NewDB()
		if err!=nil{
			fmt.Println("init db failed,",err.Error())
			return err
		}
	}
	return s.db.Delete(uuid, ID)
}

type Service interface {
	GetList(uid uuid.UUID) (*model.DBData, error)
	PutList(uuid uuid.UUID, data *model.DBListItem) error
	PostList(uuid uuid.UUID, data *model.DBListItem) (int, error)
	DeleteList(uuid uuid.UUID, ID int) error
}

func NewService() Service {
	ser := new(service)
	db,err:=db.NewDB()
	if err!=nil{
		fmt.Println("init DB failed,",err.Error())
	}
	ser.db = db
	return ser
}
