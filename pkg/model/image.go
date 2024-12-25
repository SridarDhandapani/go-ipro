package model

type CaptureMode struct {
	Mode  string `url:"img_mode,omitempty"`
	Ratio string `url:"imgratio,omitempty"`
	Fps   string `url:"img_fps,omitempty"`
}
