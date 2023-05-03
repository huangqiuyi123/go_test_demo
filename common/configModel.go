package common

type AddActivity struct {
	ID      int                    `json:"id"`
	Name    string                 `json:"name"`
	Url     string                 `json:"url"`
	Method  string                 `json:"method"`
	ResBody map[string]interface{} `json:"request"`
}
