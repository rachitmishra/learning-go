package data

type Order struct {
	Id    int32
	State OrderState
}

func NewOrder(cId int32) *Order {
	order := Order{
		Id:    cId,
		State: StateWaiting,
	}
	return &order
}
