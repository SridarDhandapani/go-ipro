package enum

type OsdPosition string

const (
	OsdPositionUpperLeft    OsdPosition = "ul"
	OsdPositionBottomLeft   OsdPosition = "bl"
	OsdPositionUpperRight   OsdPosition = "ur"
	OsdPositionBottomRight  OsdPosition = "br"
	OsdPositionUpperCenter  OsdPosition = "uc"
	OsdPositionBottomCenter OsdPosition = "bc"
)

type OsdSize string

const (
	OsdSizeLarge  OsdSize = "large"
	OsdSizeMedium OsdSize = "middle"
	OsdSizeSmall  OsdSize = "small"
	OsdSizeXSmall OsdSize = "small2"
)

type TimeDisplay string

const (
	TimeDisplay12H TimeDisplay = "12"
	TimeDisplay24H TimeDisplay = "24"
	TimeDisplayOff TimeDisplay = "off"
)
