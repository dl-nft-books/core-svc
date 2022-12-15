package resources

type PromocodeState int64

const (
	PromocodeActive PromocodeState = iota + 1
	PromocodeExpired
	PromocodeFullyUsed
	PromocodeNotFound
)

var promocodeStates = map[PromocodeState]string{
	PromocodeActive:    "active",
	PromocodeExpired:   "expired",
	PromocodeFullyUsed: "fully_used",
	PromocodeNotFound:  "not_found",
}

func (s PromocodeState) String() string {
	return promocodeStates[s]
}
