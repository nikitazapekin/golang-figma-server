package models

type FigureObject struct {
	CoordX      float64 `json:"coordX"`
	CoordY      float64 `json:"coordY"`
	Type        string  `json:"type"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	Opacity     float64 `json:"opacity"`
	Border      float64 `json:"border"`
	Background  string  `json:"background"`
	Stroke      float64 `json:"stroke"`
	StrokeColor string  `json:"strokeColor"`
	ShadowColor string  `json:"shadowColor"`
	ShadowX     float64 `json:"shadowX"`
	ShadowY     float64 `json:"shadowY"`
	Layout      int     `json:"layout"`
}

type LineObject struct {
	CoordX      float64 `json:"coordX"`
	CoordY      float64 `json:"coordY"`
	Type        string  `json:"type"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	Path        []map[string]float64 `json:"path"` 
	StrokeWidth float64 `json:"strokeWidth"`
	Color       string  `json:"color"`
	Layout      int     `json:"layout"`
}

type CanvasRequest struct {
	ID            string         `json:"id"`
	ArrayOfFigures []FigureObject `json:"arrayOfFigures"`
	ArrayOfLines   []LineObject   `json:"arrayOfLines"`
}
