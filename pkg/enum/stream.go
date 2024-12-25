package enum

type StreamEncodeType string

const (
	StreamEncodeTypeH264 StreamEncodeType = "1"
	StreamEncodeTypeH265 StreamEncodeType = "2"
)

type StreamResolution string

const (
	StreamResolutionQVGA      StreamResolution = "320"
	StreamResolution400x300   StreamResolution = "400"
	StreamResolutionVGA       StreamResolution = "640"
	StreamResolution800x600   StreamResolution = "800"
	StreamResolution1280x960  StreamResolution = "1280"
	StreamResolution1600x1200 StreamResolution = "1600"
	StreamResolution2048x1536 StreamResolution = "2048"
	StreamResolution2560x1920 StreamResolution = "2560"
	StreamResolution3072x2304 StreamResolution = "3072"
	StreamResolution320x180   StreamResolution = "320"
	StreamResolution640x360   StreamResolution = "640"
	StreamResolution1280x720  StreamResolution = "1280"
	StreamResolution1920x1080 StreamResolution = "1920"
	StreamResolution2560x1440 StreamResolution = "2560"
	StreamResolution3072x1728 StreamResolution = "3072"
	StreamResolution3840x2160 StreamResolution = "3840"
	StreamResolution640x640   StreamResolution = "640"
	StreamResolution320x320   StreamResolution = "320"
	StreamResolution1280x1280 StreamResolution = "1280"
	StreamResolution2192x2192 StreamResolution = "2192"
	StreamResolution2992x2992 StreamResolution = "2992"
)

type StreamPriority string

const (
	StreamPriorityBitRate    StreamPriority = "0"
	StreamPriorityFrameRate  StreamPriority = "1"
	StreamPriorityBestEffort StreamPriority = "2"
	StreamPriorityVBR        StreamPriority = "4"
)

type StreamFrameRate string

const (
	StreamFrameRate1fps  StreamFrameRate = "1"
	StreamFrameRate3fps  StreamFrameRate = "3"
	StreamFrameRate5fps  StreamFrameRate = "5"
	StreamFrameRAte7fps  StreamFrameRate = "7.5"
	StreamFrameRate10fps StreamFrameRate = "10"
	StreamFrameRate12fps StreamFrameRate = "12"
	StreamFrameRate15fps StreamFrameRate = "15"
	StreamFrameRate20fps StreamFrameRate = "20"
	StreamFrameRate30fps StreamFrameRate = "30"
	StreamFrameRate60fps StreamFrameRate = "60"
)

type StreamBandwidth string

const (
	StreamBandwidth64kbps    StreamBandwidth = "64"
	StreamBandwidth128kbps   StreamBandwidth = "128"
	StreamBandwidth256kbps   StreamBandwidth = "256"
	StreamBandwidth384kbps   StreamBandwidth = "384"
	StreamBandwidth512kbps   StreamBandwidth = "512"
	StreamBandwidth768kbps   StreamBandwidth = "768"
	StreamBandwidth1024kbps  StreamBandwidth = "1024"
	StreamBandwidth1536kbps  StreamBandwidth = "1536"
	StreamBandwidth2048kbps  StreamBandwidth = "2048"
	StreamBandwidth3072kbps  StreamBandwidth = "3072"
	StreamBandwidth4096kbps  StreamBandwidth = "4096"
	StreamBandwidth6144kbps  StreamBandwidth = "6144"
	StreamBandwidth8192kbps  StreamBandwidth = "8192"
	StreamBandwidth10240kbps StreamBandwidth = "10240"
	StreamBandwidth12288kbps StreamBandwidth = "12288"
	StreamBandwidth14336kbps StreamBandwidth = "14336"
	StreamBandwidth16384kbps StreamBandwidth = "16384"
	StreamBandwidth20480kbps StreamBandwidth = "20480"
	StreamBandwidth24576kbps StreamBandwidth = "24576"
)

type StreamQuality string

const (
	StreamQualityFine   StreamQuality = "fine"
	StreamQualityNormal StreamQuality = "normal"
	StreamQualityLow    StreamQuality = "low"
	StreamQualityVBR0   StreamQuality = "0"
	StreamQualityVBR1   StreamQuality = "1"
	StreamQualityVBR2   StreamQuality = "2"
	StreamQualityVBR3   StreamQuality = "3"
	StreamQualityVBR4   StreamQuality = "4"
	StreamQualityVBR5   StreamQuality = "5"
	StreamQualityVBR6   StreamQuality = "6"
	StreamQualityVBR7   StreamQuality = "7"
	StreamQualityVBR8   StreamQuality = "8"
	StreamQualityVBR9   StreamQuality = "9"
)

type StreamTransmissionType string

const (
	StreamTransmissionTypeUnicastAuto   StreamTransmissionType = "uni"
	StreamTransmissionTypeMulticast     StreamTransmissionType = "multi"
	StreamTransmissionTypeUnicastManual StreamTransmissionType = "uni_manual"
)
