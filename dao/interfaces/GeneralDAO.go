package interfaces


type GeneralDAO interface{
	Create( arg *struct{} ) error
	Delete ( arg *struct{} ) error
	Update ( arg *struct{} ) error
	GetById ( id int) ( struct{} , error)
	GetAll() ( []struct{}, error)
}
