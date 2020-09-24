package controller

type service struct {

}

type Service interface {
}

func NewSerice()(Service,error){
	ser:=new(service)
	return ser,nil
}