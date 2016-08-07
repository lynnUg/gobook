package xkcd

const db       = "./pages.json"
const  XkcdURL = "http://xkcd.com/%s/info.0.json"

type  Page struct {
	Num       int   `json:"num"`
	Transcript string  `json:"trascript"`
}
