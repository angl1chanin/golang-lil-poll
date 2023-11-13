package entity

type Poll struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Options    []*Option `json:"options"`
	TotalVotes uint      `json:"total_votes"`
}
