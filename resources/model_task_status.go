package resources

type TaskStatus int64

const (
	TaskPending TaskStatus = iota + 1
	TaskGenerating
	TaskFinishedGeneration
	TaskUploading
	TaskFinishedUploading
)

var taskStates = map[TaskStatus]string{
	TaskPending:            "pending",
	TaskGenerating:         "generating",
	TaskFinishedGeneration: "finished_generation",
	TaskUploading:          "uploading",
	TaskFinishedUploading:  "finished_uploading",
}

func (s TaskStatus) String() string {
	return taskStates[s]
}
