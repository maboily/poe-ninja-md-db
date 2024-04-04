package poeninja

type ItemType string

const (
	ItemTypeTattoo ItemType = "Tattoo"
)

type ItemOverviewLine struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	BaseType    string  `json:"baseType"`
	ChaosValue  float64 `json:"chaosValue"`
	DivineValue float64 `json:"divineValue"`

	GemLevel   int `json:"gemLevel"`
	GemQuality int `json:"gemQuality"`

	Count int `json:"count"`
}

type ItemOverviewResponse struct {
	Lines []ItemOverviewLine `json:"lines"`
}
