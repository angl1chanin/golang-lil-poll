package entity

type Option struct {
	ID     uint   `json:"id"`
	PollID uint   `json:"poll_id"`
	Name   string `json:"name"`
	Votes  uint   `json:"votes"`
}
