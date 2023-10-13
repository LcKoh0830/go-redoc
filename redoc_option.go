package redoc

type RedocOption struct {
	HideTryItPanel bool  `json:"hideTryItPanel"`
	Theme          Theme `json:"theme"`
}

type Theme struct {
	Sidebar    Sidebar    `json:"sidebar"`
	Typography Typography `json:"typography"`
}

type Sidebar struct {
	Level1Items Level1Items `json:"level1Items"`
}

type Typography struct {
	FontFamily string `json:"fontFamily"`
}

type Level1Items struct {
	TextTransform string `json:"textTransform"`
}
