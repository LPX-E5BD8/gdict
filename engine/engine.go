package engine

type Engine interface {
	Query() string
	// output for alfred workflow
	WFOutput() string
}

// EngAlfResult describe a workflow item.
type EngAlfResult struct {
	Items []*WFItem `json:"items"`
}

// workflow item for alfred 3
// https://www.alfredapp.com/help/workflows/inputs/script-filter/json/
// TODO: https://www.alfredapp.com/help/workflows/inputs/script-filter/xml/
type WFItem struct {
	Valid    bool   `json:"valid"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
	Icon     struct {
		Type string `json:"type"`
		Path string `json:"path"`
	} `json:"icon"`
}
