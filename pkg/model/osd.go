package model

import (
	"github.com/SridarDhandapani/go-ipro/pkg/enum"
)

type Title struct {
	CamTitle          string           `url:"CAMTITLE,omitempty"`
	CamIDDisplay      enum.OnOff       `url:"OSDNAMEDISP,omitempty"`
	CamID             string           `url:"OSDNAME,omitempty"`
	CamIDDisplayPlace enum.OsdPosition `url:"CAMIDPOSI,omitempty"`
	DisplaySize       enum.OsdSize     `url:"OSDSIZE,omitempty"`
}

type Clock struct {
	TimeDate      string           `url:"TIMEDATE,omitempty" validate:"regexp=^[0-9]{4}\\,[0-9]{2}\\,[0-9]{2}\\,[0-9]{2}\\,[0-9]{2}\\,[0-9]{2}$"`
	TimeDisplay   enum.TimeDisplay `url:"TIMEDISP,omitempty"`
	DisplayFormat int              `url:"TIMEFORMAT,omitempty" validate:"min=1,max=5"`
	DisplayPlace  enum.OsdPosition `url:"OSDPOSI,omitempty"`
	DisplaySize   enum.OsdSize     `url:"OSDSIZE,omitempty"`
	Timezone      int              `url:"TIMEZONE,omitempty" validate:"min=1,max=76"`
	SummerTime    enum.Toggle      `url:"STIME,omitempty"`
}
