/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateTask struct {
	Key
	Attributes UpdateTaskAttributes `json:"attributes"`
}
type UpdateTaskRequest struct {
	Data     UpdateTask `json:"data"`
	Included Included   `json:"included"`
}

type UpdateTaskListRequest struct {
	Data     []UpdateTask `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustUpdateTask - returns UpdateTask from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateTask(key Key) *UpdateTask {
	var updateTask UpdateTask
	if c.tryFindEntry(key, &updateTask) {
		return &updateTask
	}
	return nil
}
