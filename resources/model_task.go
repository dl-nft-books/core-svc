/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Task struct {
	Key
	Attributes TaskAttributes `json:"attributes"`
}
type TaskResponse struct {
	Data     Task     `json:"data"`
	Included Included `json:"included"`
}

type TaskListResponse struct {
	Data     []Task   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustTask - returns Task from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustTask(key Key) *Task {
	var task Task
	if c.tryFindEntry(key, &task) {
		return &task
	}
	return nil
}
