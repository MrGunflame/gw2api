package gw2api

// Color is a game item color
type Color struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	BaseRGB []int  `json:"base_rgb"`
	Cloth   struct {
		Brightness int     `json:"brightness"`
		Contrast   float64 `json:"contrast"`
		Hue        int     `json:"hue"`
		Saturation float64 `json:"saturation"`
		Lightness  float64 `json:"lightness"`
		RGB        []int   `json:"rgb"`
	} `json:"cloth"`
	Leather struct {
		Brightness int     `json:"brightness"`
		Contrast   float64 `json:"contrast"`
		Hue        int     `json:"hue"`
		Saturation float64 `json:"saturation"`
		Lightness  float64 `json:"lightness"`
		RGB        []int   `json:"rgb"`
	} `json:"leather"`
	Metal struct {
		Brightness int     `json:"brightness"`
		Contrast   float64 `json:"contrast"`
		Hue        int     `json:"hue"`
		Saturation float64 `json:"saturation"`
		Lightness  float64 `json:"lightness"`
		RGB        []int   `json:"rgb"`
	} `json:"metal"`
	Fur struct {
		Brightness int     `json:"brightness"`
		Contrast   float64 `json:"contrast"`
		Hue        int     `json:"hue"`
		Saturation float64 `json:"saturation"`
		Lightness  float64 `json:"lightness"`
		RGB        []int   `json:"rgb"`
	} `json:"fur"`
	Item       int      `json:"item"`
	Categories []string `json:"categories"`
}

// Colors returns the colors with the given ids
func (s *Session) Colors(ids ...int) (resp []*Color, err error) {
	err = s.getWithLang(concatStrings("/v2/colors", genArgs(ids...)), &resp)
	return
}
