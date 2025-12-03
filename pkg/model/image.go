package model

import "github.com/SridarDhandapani/go-ipro/pkg/enum"

type CaptureMode struct {
	Mode  string `url:"img_mode,omitempty"`
	Ratio string `url:"imgratio,omitempty"`
	Fps   string `url:"img_fps,omitempty"`
}

type DayNightMode struct {
	Mode enum.DayNightModeType `url:"black_white,omitempty"`
}
