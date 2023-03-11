package data

type ImageWorker struct {
	Worker
	// image data
	MaxPixels               int     `json:"max_pixels"`
	MegapixelstepsGenerated float64 `json:"megapixelsteps_generated"`
	Img2Img                 bool    `json:"img2img"`
	Painting                bool    `json:"painting"`
	PostProcessing          bool    `json:"post-processing"`
}
