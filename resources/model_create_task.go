/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateTask struct {
	Key
	Attributes CreateTaskAttributes `json:"attributes"`
}
type CreateTaskRequest struct {
	Data     CreateTask `json:"data"`
	Included Included   `json:"included"`
}

type CreateTaskListRequest struct {
	Data     []CreateTask `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustCreateTask - returns CreateTask from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateTask(key Key) *CreateTask {
	var createTask CreateTask
	if c.tryFindEntry(key, &createTask) {
		return &createTask
	}
	return nil
}
