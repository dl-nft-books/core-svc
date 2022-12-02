package resources

type DeployStatus int64

const (
	DeployPending DeployStatus = iota + 1
	DeploySuccessful
	DeployFailed
)

var deployStates = map[DeployStatus]string{
	DeployPending:    "pending",
	DeploySuccessful: "successful",
	DeployFailed:     "failed",
}

func (s DeployStatus) String() string {
	return deployStates[s]
}
