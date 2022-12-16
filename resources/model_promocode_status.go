package resources

type PromocodeState int64

const (
	PromocodeActive PromocodeState = iota + 1
	PromocodeExpired
	PromocodeFullyUsed
)

var promocodeStates = map[PromocodeState]string{
	PromocodeActive:    "active",
	PromocodeExpired:   "expired",
	PromocodeFullyUsed: "fully_used",
}

func (s PromocodeState) String() string {
	return promocodeStates[s]
}
