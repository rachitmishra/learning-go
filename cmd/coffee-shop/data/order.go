package data

type order struct {
	id    int32
	state OrderState
}

func newOrder(cId int32) *order {
	order := order{
		id:    cId,
		state: StateWaiting,
	}
	return &order
}
