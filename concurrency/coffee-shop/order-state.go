package coffeeshop

type OrderState int

const (
	StateWaiting = iota
	StateReceived
	StateGrounding
	StateGrounded
	StateMaking
	StateReady
	StateServing
	StateServed
)

var orderStates = map[OrderState]string{
	StateWaiting:   "waiting",
	StateReceived:  "received",
	StateGrounding: "grounding",
	StateGrounded:  "grounded",
	StateMaking:    "making",
	StateReady:     "ready",
	StateServing:   "serving",
	StateServed:    "served",
}

func (o OrderState) String() string {
	return orderStates[o]
}
