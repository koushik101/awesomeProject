package utils

type RiskState string

var ValidStates = map[RiskState]bool{
	"open":          true,
	"closed":        true,
	"accepted":      true,
	"investigating": true,
}

type Risk struct {
	ID          string    `json:"id"`
	State       RiskState `json:"RiskState"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
