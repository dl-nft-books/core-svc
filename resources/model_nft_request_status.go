package resources

type NftRequestStatus int8

const (
	RequestPending NftRequestStatus = iota + 1
	RequestCanceled
	RequestRejected
	RequestAccepted
	RequestExchanged
)

var nftRequestStatuses = map[NftRequestStatus]string{
	RequestPending:   "pending",
	RequestCanceled:  "canceled",
	RequestRejected:  "rejected",
	RequestAccepted:  "accepted",
	RequestExchanged: "exchanged",
}

func (s NftRequestStatus) String() string {
	return nftRequestStatuses[s]
}
