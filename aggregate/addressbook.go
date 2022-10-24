package aggregate

//联系人
type Contact struct {
	Id          int64
	Name        string
	PhoneNumber string
	GroupId     int64
}

//分组，比如 家人、朋友
type Group struct {
	Id           int64
	Name         string
	ContactCount int
	//0启用，1删除
	State int
}

func (group *Group) AddTo() {
	group.ContactCount++
}

func (group *Group) RemoveFrom() {
	group.ContactCount--
}

func (group *Group) SetAsRemoved() {
	group.State = 1
}