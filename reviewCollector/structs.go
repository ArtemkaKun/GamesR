package reviewCollector

type PARAMS struct {
	Filter        string
	Language      string
	Day_range     string
	Cursor        string
	Review_type   string
	Purchase_type string
	Num_per_page  string
	AppID         string
}

type Info struct {
	Reviews []Review
	Cursor string
}
type Review struct {
	Review string
	Language string
}
