package redoc

type RedocOption struct {
	Theme Theme `json:"theme"`
}

type Theme struct {
	Sidebar Sidebar `json:"sidebar"`
}

type Sidebar struct {
	Level1Items Level1Items `json:"level1Items"`
}

type Level1Items struct {
	TextTransform string `json:"textTransform"`
}
