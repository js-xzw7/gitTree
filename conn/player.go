package conn

type Pet struct {
	PetName     string // 宠物名
	PetNo       string // 宠物编号
	PetId       string // 宠物id
	PetElement  string // 宠物穿戴装备
	PetStatus   int64  //宠物状态 1:学习 2:打工
	ToDoEndTime int64  //干活结束时间
}
type Player struct {
	Id        string
	Name      string
	AvatarUrl string
	Score     int64
	Sex       int32
	PetInfo   Pet
}
