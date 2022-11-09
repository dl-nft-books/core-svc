package resources

type TokenStatus int64

const (
	TokenPending TokenStatus = iota + 1
	TokenUploading
	TokenFinishedUploading
)

var tokenStates = map[TokenStatus]string{
	TokenPending:           "pending",
	TokenUploading:         "uploading",
	TokenFinishedUploading: "finished_uploading",
}

func (s TokenStatus) String() string {
	return tokenStates[s]
}
