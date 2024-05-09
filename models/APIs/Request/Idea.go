package Request

type UpdateIdeaName struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

type UpdateIdeaDesc struct {
	ID          int    `json:"ID"`
	Description string `json:"Description"`
}

type IdeaItem struct {
	ID          int    `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Author      int    `json:"Author"`
}
