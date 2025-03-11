package temporal

type Root struct {
	Items    []Item          `json:"items"`
	Links    Links `json:"links"` 
	PageSize int             `json:"pageSize"`
	Self     string          `json:"self"`
	Total    int             `json:"total"`
}

type Item struct {
	Did  string `json:"did"`
	Href string `json:"href"`
}

type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

type FetchIssuersActivityResponse struct { Issuers []string }

const TaskQueue = "schedular-queue"