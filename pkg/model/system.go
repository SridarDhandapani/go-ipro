package model

type Initialize struct {
	Cmd string `url:"cmd" validate:"nonzero,regexp=^(data|html|all|reset|restart)$"`
}

//nolint:stylecheck
type CameraInfo struct {
	Browser           string `json:"BROWSER,omitempty"`
	Destination       string `json:"DESTINATION,omitempty"`
	ImageCaptureMode  string `json:"ImageCaptureMode,omitempty"`
	MacAddr           string `json:"MAC,omitempty"`
	MaxFPS            int    `json:"Maxfps,string,omitempty"`
	Name              string `json:"NAME,omitempty"`
	ProductName       string `json:"PRODUCT_NAME,omitempty"`
	ProductNameAlt    string `json:"PRODUCT_NAME_ALT,omitempty"`
	Rotation          int    `json:"Rotation,string,omitempty"`
	SDrec             string `json:"SDrec,omitempty"`
	Serial            string `json:"SERIAL,omitempty"`
	StreamEncode      int    `json:"StreamEncode,string,omitempty"`
	StreamEncode_2    int    `json:"StreamEncode_2,string,omitempty"`
	StreamEncode_3    int    `json:"StreamEncode_3,string,omitempty"`
	StreamEncode_4    int    `json:"StreamEncode_4,string,omitempty"`
	StreamMode        int    `json:"StreamMode,string,omitempty"`
	Version           string `json:"VERSION,omitempty"`
	ABitrate          int    `json:"aBitrate,string,omitempty"`
	ABitrate2         int    `json:"aBitrate2,string,omitempty"`
	ABitrate3         int    `json:"aBitrate3,string,omitempty"`
	AEnable           string `json:"aEnable,omitempty"`
	AEnc              int    `json:"aEnc,string,omitempty"`
	AInInterval       int    `json:"aInInterval,string,omitempty"`
	AInPortH264       int    `json:"aInPort_h264,string,omitempty"`
	AInPortH264_2     int    `json:"aInPort_h264_2,string,omitempty"`
	AInPortH264_3     int    `json:"aInPort_h264_3,string,omitempty"`
	AInPort_H264_4    int    `json:"aInPort_h264_4,string,omitempty"`
	AOutInterval      int    `json:"aOutInterval,string,omitempty"`
	AoutPort          int    `json:"aOutPort,string,omitempty"`
	AOoutStatus       string `json:"aOutStatus,omitempty"`
	AOutUID           int    `json:"aOutUID,string,omitempty"`
	EPort             int    `json:"ePort,string,omitempty"`
	IBitrateH264      int    `json:"iBitrate_h264,string,omitempty"`
	IBitrateH264_2    int    `json:"iBitrate_h264_2,string,omitempty"`
	IBitrateH264_3    int    `json:"iBitrate_h264_3,string,omitempty"`
	IMultiAutoH264    int    `json:"iMultiAuto_h264,string,omitempty"`
	IMultiAutoH264_2  int    `json:"iMultiAuto_h264_2,string,omitempty"`
	IMultiAutoH264_3  int    `json:"iMultiAuto_h264_3,string,omitempty"`
	IQualityH264      int    `json:"iQuality_h264,string,omitempty"`
	IQualityH264_2    int    `json:"iQuality_h264_2,string,omitempty"`
	IQualityH264_3    int    `json:"iQuality_h264_3,string,omitempty"`
	IResolutionH264   int    `json:"iResolution_h264,string,omitempty"`
	IResolutionH264_2 int    `json:"iResolution_h264_2,string,omitempty"`
	IResolutionH264_3 int    `json:"iResolution_h264_3,string,omitempty"`
	ISmartCoding      int    `json:"iSmartCoding,string,omitempty"`
	ISmartCoding_2    int    `json:"iSmartCoding_2,string,omitempty"`
	ISmartCoding_3    int    `json:"iSmartCoding_3,string,omitempty"`
	ITransmitH264     int    `json:"iTransmit_h264,string,omitempty"`
	ITransmitH264_2   int    `json:"iTransmit_h264_2,string,omitempty"`
	ITransmitH264_3   int    `json:"iTransmit_h264_3,string,omitempty"`
	ITransmitMode     int    `json:"iTransmit_mode,string,omitempty"`
	ITransmitMode_2   int    `json:"iTransmit_mode_2,string,omitempty"`
	ITransmitMode_3   int    `json:"iTransmit_mode_3,string,omitempty"`
	Ration            string `json:"ratio,omitempty"`
	SAux              string `json:"sAUX,omitempty"`
	SAlarm            string `json:"sAlarm,omitempty"`
	SDeliveryH264     string `json:"sDelivery_h264,omitempty"`
	SDeliveryH264_2   string `json:"sDelivery_h264_2,omitempty"`
	SDeliveryH264_3   string `json:"sDelivery_h264_3,omitempty"`
	SRtspModeH264     int    `json:"sRtspMode_h264,string,omitempty"`
	SRtspModeH264_2   int    `json:"sRtspMode_h264_2,string,omitempty"`
	SRtspModeH264_3   int    `json:"sRtspMode_h264_3,string,omitempty"`
	SRtspModeH264_4   int    `json:"sRtspMode_h264_4,string,omitempty"`
}

//nolint:stylecheck
type CameraData struct {
	CAMTITLE                  string `json:"CAMTITLE,omitempty" url:"CAMTITLE,omitempty"`
	TIMEDATE                  string `json:"TIMEDATE,omitempty" url:"TIMEDATE,omitempty"`
	TIMEFORMAT                string `json:"TIMEFORMAT,omitempty" url:"TIMEFORMAT,omitempty"`
	TIMEDISP                  string `json:"TIMEDISP,omitempty" url:"TIMEDISP,omitempty"`
	TIMEZONE                  string `json:"TIMEZONE,omitempty" url:"TIMEZONE,omitempty"`
	STIME                     string `json:"STIME,omitempty" url:"STIME,omitempty"`
	STIMES_MON                string `json:"STIMES_MON,omitempty" url:"STIMES_MON,omitempty"`
	STIMES_WEEK               string `json:"STIMES_WEEK,omitempty" url:"STIMES_WEEK,omitempty"`
	STIMES_DOTW               string `json:"STIMES_DOTW,omitempty" url:"STIMES_DOTW,omitempty"`
	STIMES_HOUR               string `json:"STIMES_HOUR,omitempty" url:"STIMES_HOUR,omitempty"`
	STIMES_AMPM               string `json:"STIMES_AMPM,omitempty" url:"STIMES_AMPM,omitempty"`
	STIMEE_MON                string `json:"STIMEE_MON,omitempty" url:"STIMEE_MON,omitempty"`
	STIMEE_WEEK               string `json:"STIMEE_WEEK,omitempty" url:"STIMEE_WEEK,omitempty"`
	STIMEE_DOTW               string `json:"STIMEE_DOTW,omitempty" url:"STIMEE_DOTW,omitempty"`
	STIMEE_HOUR               string `json:"STIMEE_HOUR,omitempty" url:"STIMEE_HOUR,omitempty"`
	STIMEE_AMPM               string `json:"STIMEE_AMPM,omitempty" url:"STIMEE_AMPM,omitempty"`
	OSDNAMEDISP               string `json:"OSDNAMEDISP,omitempty" url:"OSDNAMEDISP,omitempty"`
	OSDNAME                   string `json:"OSDNAME,omitempty" url:"OSDNAME,omitempty"`
	OSDPOSI                   string `json:"OSDPOSI,omitempty" url:"OSDPOSI,omitempty"`
	CAMIDPOSI                 string `json:"CAMIDPOSI,omitempty" url:"CAMIDPOSI,omitempty"`
	OSDSIZE                   string `json:"OSDSIZE,omitempty" url:"OSDSIZE,omitempty"`
	BRIGHTNESSDISP            string `json:"BRIGHTNESSDISP,omitempty" url:"BRIGHTNESSDISP,omitempty"`
	UPSIDEDOWN                string `json:"UPSIDEDOWN,omitempty" url:"UPSIDEDOWN,omitempty"`
	IMAGE_ROTATION            string `json:"IMAGE_ROTATION,omitempty" url:"IMAGE_ROTATION,omitempty"`
	LED                       string `json:"LED,omitempty" url:"LED,omitempty"`
	IMAGESELECT               string `json:"IMAGESELECT,omitempty" url:"IMAGESELECT,omitempty"`
	IMAGERATIO                string `json:"IMAGERATIO,omitempty" url:"IMAGERATIO,omitempty"`
	IMAGERATIO_SUB            string `json:"IMAGERATIO_SUB,omitempty" url:"IMAGERATIO_SUB,omitempty"`
	IMAGEFPS                  string `json:"IMAGEFPS,omitempty" url:"IMAGEFPS,omitempty"`
	LIVESTREAM                string `json:"LIVESTREAM,omitempty" url:"LIVESTREAM,omitempty"`
	LIVEINT                   string `json:"LIVEINT,omitempty" url:"LIVEINT,omitempty"`
	LIVEQUALBASE              string `json:"LIVEQUALBASE,omitempty" url:"LIVEQUALBASE,omitempty"`
	LIVEREQMODE               string `json:"LIVEREQMODE,omitempty" url:"LIVEREQMODE,omitempty"`
	LIVESIZE                  string `json:"LIVESIZE,omitempty" url:"LIVESIZE,omitempty"`
	LIVEQUAL                  string `json:"LIVEQUAL,omitempty" url:"LIVEQUAL,omitempty"`
	LIVESIZE2                 string `json:"LIVESIZE2,omitempty" url:"LIVESIZE2,omitempty"`
	LIVEQUAL2                 string `json:"LIVEQUAL2,omitempty" url:"LIVEQUAL2,omitempty"`
	LIVESIZE3                 string `json:"LIVESIZE3,omitempty" url:"LIVESIZE3,omitempty"`
	LIVEQUAL3                 string `json:"LIVEQUAL3,omitempty" url:"LIVEQUAL3,omitempty"`
	STREAMMODE                string `json:"STREAMMODE,omitempty" url:"STREAMMODE,omitempty"`
	H264                      string `json:"H264,omitempty" url:"H264,omitempty"`
	STREAMENCMODE             string `json:"STREAMENCMODE,omitempty" url:"STREAMENCMODE,omitempty"`
	H264RTSPMODE              string `json:"H264RTSPMODE,omitempty" url:"H264RTSPMODE,omitempty"`
	H264BWC                   string `json:"H264BWC,omitempty" url:"H264BWC,omitempty"`
	NRH264BWC                 string `json:"NRH264BWC,omitempty" url:"NRH264BWC,omitempty"`
	H264BWCMIN                string `json:"H264BWCMIN,omitempty" url:"H264BWCMIN,omitempty"`
	H264SIZE                  string `json:"H264SIZE,omitempty" url:"H264SIZE,omitempty"`
	NRH264SIZE                string `json:"NRH264SIZE,omitempty" url:"NRH264SIZE,omitempty"`
	H264FPRIORITY             string `json:"H264FPRIORITY,omitempty" url:"H264FPRIORITY,omitempty"`
	H264NRFRAMERATE           string `json:"H264NRFRAMERATE,omitempty" url:"H264NRFRAMERATE,omitempty"`
	H264QUAL                  string `json:"H264QUAL,omitempty" url:"H264QUAL,omitempty"`
	NRH264QUAL                string `json:"NRH264QUAL,omitempty" url:"NRH264QUAL,omitempty"`
	H264RINT                  string `json:"H264RINT,omitempty" url:"H264RINT,omitempty"`
	H264MTD                   string `json:"H264MTD,omitempty" url:"H264MTD,omitempty"`
	UNIANDMULTI               string `json:"UNIANDMULTI,omitempty" url:"UNIANDMULTI,omitempty"`
	H264MLADD1                string `json:"H264MLADD1,omitempty" url:"H264MLADD1,omitempty"`
	H264MLADD2                string `json:"H264MLADD2,omitempty" url:"H264MLADD2,omitempty"`
	H264MLADD3                string `json:"H264MLADD3,omitempty" url:"H264MLADD3,omitempty"`
	H264MLADD4                string `json:"H264MLADD4,omitempty" url:"H264MLADD4,omitempty"`
	H264MLADD                 string `json:"H264MLADD,omitempty" url:"H264MLADD,omitempty"`
	H264MLPORT                string `json:"H264MLPORT,omitempty" url:"H264MLPORT,omitempty"`
	H264MLTTL                 string `json:"H264MLTTL,omitempty" url:"H264MLTTL,omitempty"`
	H264UNIPORT               string `json:"H264UNIPORT,omitempty" url:"H264UNIPORT,omitempty"`
	H264UNIPORT2              string `json:"H264UNIPORT2,omitempty" url:"H264UNIPORT2,omitempty"`
	H264ENCTYPE               string `json:"H264ENCTYPE,omitempty" url:"H264ENCTYPE,omitempty"`
	H264_2                    string `json:"H264_2,omitempty" url:"H264_2,omitempty"`
	STREAMENCMODE_2           string `json:"STREAMENCMODE_2,omitempty" url:"STREAMENCMODE_2,omitempty"`
	H264RTSPMODE_2            string `json:"H264RTSPMODE_2,omitempty" url:"H264RTSPMODE_2,omitempty"`
	H264BWC_2                 string `json:"H264BWC_2,omitempty" url:"H264BWC_2,omitempty"`
	NRH264BWC_2               string `json:"NRH264BWC_2,omitempty" url:"NRH264BWC_2,omitempty"`
	H264BWCMIN_2              string `json:"H264BWCMIN_2,omitempty" url:"H264BWCMIN_2,omitempty"`
	H264SIZE_2                string `json:"H264SIZE_2,omitempty" url:"H264SIZE_2,omitempty"`
	NRH264SIZE_2              string `json:"NRH264SIZE_2,omitempty" url:"NRH264SIZE_2,omitempty"`
	H264FPRIORITY_2           string `json:"H264FPRIORITY_2,omitempty" url:"H264FPRIORITY_2,omitempty"`
	H264NRFRAMERATE_2         string `json:"H264NRFRAMERATE_2,omitempty" url:"H264NRFRAMERATE_2,omitempty"`
	H264QUAL_2                string `json:"H264QUAL_2,omitempty" url:"H264QUAL_2,omitempty"`
	NRH264QUAL_2              string `json:"NRH264QUAL_2,omitempty" url:"NRH264QUAL_2,omitempty"`
	H264RINT_2                string `json:"H264RINT_2,omitempty" url:"H264RINT_2,omitempty"`
	H264MTD_2                 string `json:"H264MTD_2,omitempty" url:"H264MTD_2,omitempty"`
	UNIANDMULTI2              string `json:"UNIANDMULTI2,omitempty" url:"UNIANDMULTI2,omitempty"`
	H264MLADD1_2              string `json:"H264MLADD1_2,omitempty" url:"H264MLADD1_2,omitempty"`
	H264MLADD2_2              string `json:"H264MLADD2_2,omitempty" url:"H264MLADD2_2,omitempty"`
	H264MLADD3_2              string `json:"H264MLADD3_2,omitempty" url:"H264MLADD3_2,omitempty"`
	H264MLADD4_2              string `json:"H264MLADD4_2,omitempty" url:"H264MLADD4_2,omitempty"`
	H264MLADD_2               string `json:"H264MLADD_2,omitempty" url:"H264MLADD_2,omitempty"`
	H264MLPORT_2              string `json:"H264MLPORT_2,omitempty" url:"H264MLPORT_2,omitempty"`
	H264MLTTL_2               string `json:"H264MLTTL_2,omitempty" url:"H264MLTTL_2,omitempty"`
	H264UNIPORT_2             string `json:"H264UNIPORT_2,omitempty" url:"H264UNIPORT_2,omitempty"`
	H264UNIPORT2_2            string `json:"H264UNIPORT2_2,omitempty" url:"H264UNIPORT2_2,omitempty"`
	H264ENCTYPE_2             string `json:"H264ENCTYPE_2,omitempty" url:"H264ENCTYPE_2,omitempty"`
	H264_3                    string `json:"H264_3,omitempty" url:"H264_3,omitempty"`
	STREAMENCMODE_3           string `json:"STREAMENCMODE_3,omitempty" url:"STREAMENCMODE_3,omitempty"`
	H264RTSPMODE_3            string `json:"H264RTSPMODE_3,omitempty" url:"H264RTSPMODE_3,omitempty"`
	H264BWC_3                 string `json:"H264BWC_3,omitempty" url:"H264BWC_3,omitempty"`
	NRH264BWC_3               string `json:"NRH264BWC_3,omitempty" url:"NRH264BWC_3,omitempty"`
	H264BWCMIN_3              string `json:"H264BWCMIN_3,omitempty" url:"H264BWCMIN_3,omitempty"`
	H264SIZE_3                string `json:"H264SIZE_3,omitempty" url:"H264SIZE_3,omitempty"`
	NRH264SIZE_3              string `json:"NRH264SIZE_3,omitempty" url:"NRH264SIZE_3,omitempty"`
	H264FPRIORITY_3           string `json:"H264FPRIORITY_3,omitempty" url:"H264FPRIORITY_3,omitempty"`
	H264NRFRAMERATE_3         string `json:"H264NRFRAMERATE_3,omitempty" url:"H264NRFRAMERATE_3,omitempty"`
	H264QUAL_3                string `json:"H264QUAL_3,omitempty" url:"H264QUAL_3,omitempty"`
	NRH264QUAL_3              string `json:"NRH264QUAL_3,omitempty" url:"NRH264QUAL_3,omitempty"`
	H264RINT_3                string `json:"H264RINT_3,omitempty" url:"H264RINT_3,omitempty"`
	H264MTD_3                 string `json:"H264MTD_3,omitempty" url:"H264MTD_3,omitempty"`
	UNIANDMULTI3              string `json:"UNIANDMULTI3,omitempty" url:"UNIANDMULTI3,omitempty"`
	H264MLADD1_3              string `json:"H264MLADD1_3,omitempty" url:"H264MLADD1_3,omitempty"`
	H264MLADD2_3              string `json:"H264MLADD2_3,omitempty" url:"H264MLADD2_3,omitempty"`
	H264MLADD3_3              string `json:"H264MLADD3_3,omitempty" url:"H264MLADD3_3,omitempty"`
	H264MLADD4_3              string `json:"H264MLADD4_3,omitempty" url:"H264MLADD4_3,omitempty"`
	H264MLADD_3               string `json:"H264MLADD_3,omitempty" url:"H264MLADD_3,omitempty"`
	H264MLPORT_3              string `json:"H264MLPORT_3,omitempty" url:"H264MLPORT_3,omitempty"`
	H264MLTTL_3               string `json:"H264MLTTL_3,omitempty" url:"H264MLTTL_3,omitempty"`
	H264UNIPORT_3             string `json:"H264UNIPORT_3,omitempty" url:"H264UNIPORT_3,omitempty"`
	H264UNIPORT2_3            string `json:"H264UNIPORT2_3,omitempty" url:"H264UNIPORT2_3,omitempty"`
	H264ENCTYPE_3             string `json:"H264ENCTYPE_3,omitempty" url:"H264ENCTYPE_3,omitempty"`
	RTSPPORT                  string `json:"RTSPPORT,omitempty" url:"RTSPPORT,omitempty"`
	H264MLAUTO                string `json:"H264MLAUTO,omitempty" url:"H264MLAUTO,omitempty"`
	H264MLAUTO_2              string `json:"H264MLAUTO_2,omitempty" url:"H264MLAUTO_2,omitempty"`
	H264MLAUTO_3              string `json:"H264MLAUTO_3,omitempty" url:"H264MLAUTO_3,omitempty"`
	MEGASD                    string `json:"MEGASD,omitempty" url:"MEGASD,omitempty"`
	SD_LEVEL                  string `json:"SD_LEVEL,omitempty" url:"SD_LEVEL,omitempty"`
	GAIN_SD_OFF               string `json:"GAIN_SD_OFF,omitempty" url:"GAIN_SD_OFF,omitempty"`
	BLC                       string `json:"BLC,omitempty" url:"BLC,omitempty"`
	BLCMASK_16                string `json:"BLCMASK_16,omitempty" url:"BLCMASK_16,omitempty"`
	ALCELC                    string `json:"ALCELC,omitempty" url:"ALCELC,omitempty"`
	BRIGHTNESS                string `json:"BRIGHTNESS,omitempty" url:"BRIGHTNESS,omitempty"`
	AGC                       string `json:"AGC,omitempty" url:"AGC,omitempty"`
	SHUTTER_IRIS_PRIORITY     string `json:"SHUTTER_IRIS_PRIORITY,omitempty" url:"SHUTTER_IRIS_PRIORITY,omitempty"`
	SLOWSHUTTER_GAIN          string `json:"SLOWSHUTTER_GAIN,omitempty" url:"SLOWSHUTTER_GAIN,omitempty"`
	SLOWSHUTTER_BRIGHT        string `json:"SLOWSHUTTER_BRIGHT,omitempty" url:"SLOWSHUTTER_BRIGHT,omitempty"`
	SENSITIVITY               string `json:"SENSITIVITY,omitempty" url:"SENSITIVITY,omitempty"`
	SHUTTER                   string `json:"SHUTTER,omitempty" url:"SHUTTER,omitempty"`
	MINIMUMSHUTTER            string `json:"MINIMUMSHUTTER,omitempty" url:"MINIMUMSHUTTER,omitempty"`
	OPEN_LIMIT                string `json:"OPEN_LIMIT,omitempty" url:"OPEN_LIMIT,omitempty"`
	CLOSE_LIMIT               string `json:"CLOSE_LIMIT,omitempty" url:"CLOSE_LIMIT,omitempty"`
	BW                        string `json:"BW,omitempty" url:"BW,omitempty"`
	BWLEVEL                   string `json:"BWLEVEL,omitempty" url:"BWLEVEL,omitempty"`
	BWTIME                    string `json:"BWTIME,omitempty" url:"BWTIME,omitempty"`
	IRLED                     string `json:"IRLED,omitempty" url:"IRLED,omitempty"`
	IRLEDCTRL                 string `json:"IRLEDCTRL,omitempty" url:"IRLEDCTRL,omitempty"`
	COLOR_RETURN_LEVEL        string `json:"COLOR_RETURN_LEVEL,omitempty" url:"COLOR_RETURN_LEVEL,omitempty"`
	WHITEBAL                  string `json:"WHITEBAL,omitempty" url:"WHITEBAL,omitempty"`
	RVOL                      string `json:"RVOL,omitempty" url:"RVOL,omitempty"`
	BVOL                      string `json:"BVOL,omitempty" url:"BVOL,omitempty"`
	AUTO_ADJUST               string `json:"AUTO_ADJUST,omitempty" url:"AUTO_ADJUST,omitempty"`
	MOTION_DETECT_LEVEL       string `json:"MOTION_DETECT_LEVEL,omitempty" url:"MOTION_DETECT_LEVEL,omitempty"`
	BLC_DETECT_LEVEL          string `json:"BLC_DETECT_LEVEL,omitempty" url:"BLC_DETECT_LEVEL,omitempty"`
	HLC_DETECT_LEVEL          string `json:"HLC_DETECT_LEVEL,omitempty" url:"HLC_DETECT_LEVEL,omitempty"`
	CONTRAST_AUTO             string `json:"CONTRAST_AUTO,omitempty" url:"CONTRAST_AUTO,omitempty"`
	DNR                       string `json:"DNR,omitempty" url:"DNR,omitempty"`
	CHROMA                    string `json:"CHROMA,omitempty" url:"CHROMA,omitempty"`
	APERTURE                  string `json:"APERTURE,omitempty" url:"APERTURE,omitempty"`
	PEDESTAL                  string `json:"PEDESTAL,omitempty" url:"PEDESTAL,omitempty"`
	DNR_DETAIL                string `json:"DNR_DETAIL,omitempty" url:"DNR_DETAIL,omitempty"`
	THREEDNR_LEVEL            string `json:"3DNR_LEVEL,omitempty" url:"3DNR_LEVEL,omitempty"`
	TWODNR_LEVEL              string `json:"2DNR_LEVEL,omitempty" url:"2DNR_LEVEL,omitempty"`
	HLC                       string `json:"HLC,omitempty" url:"HLC,omitempty"`
	FOG                       string `json:"FOG,omitempty" url:"FOG,omitempty"`
	FOGLEVEL                  string `json:"FOGLEVEL,omitempty" url:"FOGLEVEL,omitempty"`
	CLBW                      string `json:"CLBW,omitempty" url:"CLBW,omitempty"`
	AFULX1                    string `json:"AFULX1,omitempty" url:"AFULX1,omitempty"`
	AFULY1                    string `json:"AFULY1,omitempty" url:"AFULY1,omitempty"`
	AFBRX1                    string `json:"AFBRX1,omitempty" url:"AFBRX1,omitempty"`
	AFBRX2                    string `json:"AFBRX2,omitempty" url:"AFBRX2,omitempty"`
	AFBRY1                    string `json:"AFBRY1,omitempty" url:"AFBRY1,omitempty"`
	ELZOOM                    string `json:"ELZOOM,omitempty" url:"ELZOOM,omitempty"`
	PRVENT                    string `json:"PRVENT,omitempty" url:"PRVENT,omitempty"`
	PRVULX1                   string `json:"PRVULX1,omitempty" url:"PRVULX1,omitempty"`
	PRVULY1                   string `json:"PRVULY1,omitempty" url:"PRVULY1,omitempty"`
	PRVBRX1                   string `json:"PRVBRX1,omitempty" url:"PRVBRX1,omitempty"`
	PRVBRY1                   string `json:"PRVBRY1,omitempty" url:"PRVBRY1,omitempty"`
	PRVULX2                   string `json:"PRVULX2,omitempty" url:"PRVULX2,omitempty"`
	PRVULY2                   string `json:"PRVULY2,omitempty" url:"PRVULY2,omitempty"`
	PRVBRX2                   string `json:"PRVBRX2,omitempty" url:"PRVBRX2,omitempty"`
	PRVBRY2                   string `json:"PRVBRY2,omitempty" url:"PRVBRY2,omitempty"`
	PRVULX3                   string `json:"PRVULX3,omitempty" url:"PRVULX3,omitempty"`
	PRVULY3                   string `json:"PRVULY3,omitempty" url:"PRVULY3,omitempty"`
	PRVBRX3                   string `json:"PRVBRX3,omitempty" url:"PRVBRX3,omitempty"`
	PRVBRY3                   string `json:"PRVBRY3,omitempty" url:"PRVBRY3,omitempty"`
	PRVULX4                   string `json:"PRVULX4,omitempty" url:"PRVULX4,omitempty"`
	PRVULY4                   string `json:"PRVULY4,omitempty" url:"PRVULY4,omitempty"`
	PRVBRX4                   string `json:"PRVBRX4,omitempty" url:"PRVBRX4,omitempty"`
	PRVBRY4                   string `json:"PRVBRY4,omitempty" url:"PRVBRY4,omitempty"`
	PRVULX5                   string `json:"PRVULX5,omitempty" url:"PRVULX5,omitempty"`
	PRVULY5                   string `json:"PRVULY5,omitempty" url:"PRVULY5,omitempty"`
	PRVBRX5                   string `json:"PRVBRX5,omitempty" url:"PRVBRX5,omitempty"`
	PRVBRY5                   string `json:"PRVBRY5,omitempty" url:"PRVBRY5,omitempty"`
	PRVULX6                   string `json:"PRVULX6,omitempty" url:"PRVULX6,omitempty"`
	PRVULY6                   string `json:"PRVULY6,omitempty" url:"PRVULY6,omitempty"`
	PRVBRX6                   string `json:"PRVBRX6,omitempty" url:"PRVBRX6,omitempty"`
	PRVBRY6                   string `json:"PRVBRY6,omitempty" url:"PRVBRY6,omitempty"`
	PRVULX7                   string `json:"PRVULX7,omitempty" url:"PRVULX7,omitempty"`
	PRVULY7                   string `json:"PRVULY7,omitempty" url:"PRVULY7,omitempty"`
	PRVBRX7                   string `json:"PRVBRX7,omitempty" url:"PRVBRX7,omitempty"`
	PRVBRY7                   string `json:"PRVBRY7,omitempty" url:"PRVBRY7,omitempty"`
	PRVULX8                   string `json:"PRVULX8,omitempty" url:"PRVULX8,omitempty"`
	PRVULY8                   string `json:"PRVULY8,omitempty" url:"PRVULY8,omitempty"`
	PRVBRX8                   string `json:"PRVBRX8,omitempty" url:"PRVBRX8,omitempty"`
	PRVBRY8                   string `json:"PRVBRY8,omitempty" url:"PRVBRY8,omitempty"`
	VIQSSTREAM1               string `json:"VIQSSTREAM1,omitempty" url:"VIQSSTREAM1,omitempty"`
	VIQSSTREAM2               string `json:"VIQSSTREAM2,omitempty" url:"VIQSSTREAM2,omitempty"`
	VIQSSTREAM3               string `json:"VIQSSTREAM3,omitempty" url:"VIQSSTREAM3,omitempty"`
	VIQSULX                   string `json:"VIQSULX,omitempty" url:"VIQSULX,omitempty"`
	VIQSULY                   string `json:"VIQSULY,omitempty" url:"VIQSULY,omitempty"`
	VIQSBRX                   string `json:"VIQSBRX,omitempty" url:"VIQSBRX,omitempty"`
	VIQSBRY                   string `json:"VIQSBRY,omitempty" url:"VIQSBRY,omitempty"`
	VIQSULX2                  string `json:"VIQSULX2,omitempty" url:"VIQSULX2,omitempty"`
	VIQSULY2                  string `json:"VIQSULY2,omitempty" url:"VIQSULY2,omitempty"`
	VIQSBRX2                  string `json:"VIQSBRX2,omitempty" url:"VIQSBRX2,omitempty"`
	VIQSBRY2                  string `json:"VIQSBRY2,omitempty" url:"VIQSBRY2,omitempty"`
	VIQSULX3                  string `json:"VIQSULX3,omitempty" url:"VIQSULX3,omitempty"`
	VIQSULY3                  string `json:"VIQSULY3,omitempty" url:"VIQSULY3,omitempty"`
	VIQSBRX3                  string `json:"VIQSBRX3,omitempty" url:"VIQSBRX3,omitempty"`
	VIQSBRY3                  string `json:"VIQSBRY3,omitempty" url:"VIQSBRY3,omitempty"`
	VIQSULX4                  string `json:"VIQSULX4,omitempty" url:"VIQSULX4,omitempty"`
	VIQSULY4                  string `json:"VIQSULY4,omitempty" url:"VIQSULY4,omitempty"`
	VIQSBRX4                  string `json:"VIQSBRX4,omitempty" url:"VIQSBRX4,omitempty"`
	VIQSBRY4                  string `json:"VIQSBRY4,omitempty" url:"VIQSBRY4,omitempty"`
	VIQSULX5                  string `json:"VIQSULX5,omitempty" url:"VIQSULX5,omitempty"`
	VIQSULY5                  string `json:"VIQSULY5,omitempty" url:"VIQSULY5,omitempty"`
	VIQSBRX5                  string `json:"VIQSBRX5,omitempty" url:"VIQSBRX5,omitempty"`
	VIQSBRY5                  string `json:"VIQSBRY5,omitempty" url:"VIQSBRY5,omitempty"`
	VIQSULX6                  string `json:"VIQSULX6,omitempty" url:"VIQSULX6,omitempty"`
	VIQSULY6                  string `json:"VIQSULY6,omitempty" url:"VIQSULY6,omitempty"`
	VIQSBRX6                  string `json:"VIQSBRX6,omitempty" url:"VIQSBRX6,omitempty"`
	VIQSBRY6                  string `json:"VIQSBRY6,omitempty" url:"VIQSBRY6,omitempty"`
	VIQSULX7                  string `json:"VIQSULX7,omitempty" url:"VIQSULX7,omitempty"`
	VIQSULY7                  string `json:"VIQSULY7,omitempty" url:"VIQSULY7,omitempty"`
	VIQSBRX7                  string `json:"VIQSBRX7,omitempty" url:"VIQSBRX7,omitempty"`
	VIQSBRY7                  string `json:"VIQSBRY7,omitempty" url:"VIQSBRY7,omitempty"`
	VIQSULX8                  string `json:"VIQSULX8,omitempty" url:"VIQSULX8,omitempty"`
	VIQSULY8                  string `json:"VIQSULY8,omitempty" url:"VIQSULY8,omitempty"`
	VIQSBRX8                  string `json:"VIQSBRX8,omitempty" url:"VIQSBRX8,omitempty"`
	VIQSBRY8                  string `json:"VIQSBRY8,omitempty" url:"VIQSBRY8,omitempty"`
	AUDIO                     string `json:"AUDIO,omitempty" url:"AUDIO,omitempty"`
	AUDIOSENS                 string `json:"AUDIOSENS,omitempty" url:"AUDIOSENS,omitempty"`
	AUDIOENC                  string `json:"AUDIOENC,omitempty" url:"AUDIOENC,omitempty"`
	AUDIOBITRATE              string `json:"AUDIOBITRATE,omitempty" url:"AUDIOBITRATE,omitempty"`
	AUDIOBITRATE_AAC          string `json:"AUDIOBITRATE_AAC,omitempty" url:"AUDIOBITRATE_AAC,omitempty"`
	AUDIOINT                  string `json:"AUDIOINT,omitempty" url:"AUDIOINT,omitempty"`
	AUDIOOUTINT               string `json:"AUDIOOUTINT,omitempty" url:"AUDIOOUTINT,omitempty"`
	AUDIOOUTSENS              string `json:"AUDIOOUTSENS,omitempty" url:"AUDIOOUTSENS,omitempty"`
	AUDIOOUTTLIMIT            string `json:"AUDIOOUTTLIMIT,omitempty" url:"AUDIOOUTTLIMIT,omitempty"`
	AUDIOOUTPORT              string `json:"AUDIOOUTPORT,omitempty" url:"AUDIOOUTPORT,omitempty"`
	AUDIOSTATUS               string `json:"AUDIOSTATUS,omitempty" url:"AUDIOSTATUS,omitempty"`
	AUDIOTLIMIT               string `json:"AUDIOTLIMIT,omitempty" url:"AUDIOTLIMIT,omitempty"`
	AUDIOAGC                  string `json:"AUDIOAGC,omitempty" url:"AUDIOAGC,omitempty"`
	G726_ENDIAN               string `json:"G726_ENDIAN,omitempty" url:"G726_ENDIAN,omitempty"`
	SDCARD                    string `json:"SDCARD,omitempty" url:"SDCARD,omitempty"`
	SDREC_AUDIO               string `json:"SDREC_AUDIO,omitempty" url:"SDREC_AUDIO,omitempty"`
	SDREC_AUDIO2              string `json:"SDREC_AUDIO2,omitempty" url:"SDREC_AUDIO2,omitempty"`
	SDREMNOTICE               string `json:"SDREMNOTICE,omitempty" url:"SDREMNOTICE,omitempty"`
	SDRECPTN                  string `json:"SDRECPTN,omitempty" url:"SDRECPTN,omitempty"`
	SDRECCODEC                string `json:"SDRECCODEC,omitempty" url:"SDRECCODEC,omitempty"`
	SDREC                     string `json:"SDREC,omitempty" url:"SDREC,omitempty"`
	SDREC_VMD                 string `json:"SDREC_VMD,omitempty" url:"SDREC_VMD,omitempty"`
	SDREC_SCD                 string `json:"SDREC_SCD,omitempty" url:"SDREC_SCD,omitempty"`
	SDREC_COM                 string `json:"SDREC_COM,omitempty" url:"SDREC_COM,omitempty"`
	SDREC_EXTMODE             string `json:"SDREC_EXTMODE,omitempty" url:"SDREC_EXTMODE,omitempty"`
	SDRECNAME                 string `json:"SDRECNAME,omitempty" url:"SDRECNAME,omitempty"`
	SDRECINT                  string `json:"SDRECINT,omitempty" url:"SDRECINT,omitempty"`
	SDRECRESOLUTION           string `json:"SDRECRESOLUTION,omitempty" url:"SDRECRESOLUTION,omitempty"`
	SDRECNUM                  string `json:"SDRECNUM,omitempty" url:"SDRECNUM,omitempty"`
	SDRECINT_PRE              string `json:"SDRECINT_PRE,omitempty" url:"SDRECINT_PRE,omitempty"`
	SDRECNUM_PRE              string `json:"SDRECNUM_PRE,omitempty" url:"SDRECNUM_PRE,omitempty"`
	PREALM_TIME               string `json:"PREALM_TIME,omitempty" url:"PREALM_TIME,omitempty"`
	ALMRECTIME                string `json:"ALMRECTIME,omitempty" url:"ALMRECTIME,omitempty"`
	SDREM                     string `json:"SDREM,omitempty" url:"SDREM,omitempty"`
	SDTOTAL                   string `json:"SDTOTAL,omitempty" url:"SDTOTAL,omitempty"`
	SDREC_NWLOST_AUTO_INT     string `json:"SDREC_NWLOST_AUTO_INT,omitempty" url:"SDREC_NWLOST_AUTO_INT,omitempty"`
	AVMD                      string `json:"AVMD,omitempty" url:"AVMD,omitempty"`
	ACMD                      string `json:"ACMD,omitempty" url:"ACMD,omitempty"`
	ACMDPORT                  string `json:"ACMDPORT,omitempty" url:"ACMDPORT,omitempty"`
	MASKTIME                  string `json:"MASKTIME,omitempty" url:"MASKTIME,omitempty"`
	ARETRY                    string `json:"ARETRY,omitempty" url:"ARETRY,omitempty"`
	APREINT                   string `json:"APREINT,omitempty" url:"APREINT,omitempty"`
	APRENUM                   string `json:"APRENUM,omitempty" url:"APRENUM,omitempty"`
	VMDULX1                   string `json:"VMDULX1,omitempty" url:"VMDULX1,omitempty"`
	VMDULY1                   string `json:"VMDULY1,omitempty" url:"VMDULY1,omitempty"`
	VMDBRX1                   string `json:"VMDBRX1,omitempty" url:"VMDBRX1,omitempty"`
	VMDBRY1                   string `json:"VMDBRY1,omitempty" url:"VMDBRY1,omitempty"`
	VMDULX2                   string `json:"VMDULX2,omitempty" url:"VMDULX2,omitempty"`
	VMDULY2                   string `json:"VMDULY2,omitempty" url:"VMDULY2,omitempty"`
	VMDBRX2                   string `json:"VMDBRX2,omitempty" url:"VMDBRX2,omitempty"`
	VMDBRY2                   string `json:"VMDBRY2,omitempty" url:"VMDBRY2,omitempty"`
	VMDULX3                   string `json:"VMDULX3,omitempty" url:"VMDULX3,omitempty"`
	VMDULY3                   string `json:"VMDULY3,omitempty" url:"VMDULY3,omitempty"`
	VMDBRX3                   string `json:"VMDBRX3,omitempty" url:"VMDBRX3,omitempty"`
	VMDBRY3                   string `json:"VMDBRY3,omitempty" url:"VMDBRY3,omitempty"`
	VMDULX4                   string `json:"VMDULX4,omitempty" url:"VMDULX4,omitempty"`
	VMDULY4                   string `json:"VMDULY4,omitempty" url:"VMDULY4,omitempty"`
	VMDBRX4                   string `json:"VMDBRX4,omitempty" url:"VMDBRX4,omitempty"`
	VMDBRY4                   string `json:"VMDBRY4,omitempty" url:"VMDBRY4,omitempty"`
	VMDSTATUS1                string `json:"VMDSTATUS1,omitempty" url:"VMDSTATUS1,omitempty"`
	VMDSTATUS2                string `json:"VMDSTATUS2,omitempty" url:"VMDSTATUS2,omitempty"`
	VMDSTATUS3                string `json:"VMDSTATUS3,omitempty" url:"VMDSTATUS3,omitempty"`
	VMDSTATUS4                string `json:"VMDSTATUS4,omitempty" url:"VMDSTATUS4,omitempty"`
	VMDAREA                   string `json:"VMDAREA,omitempty" url:"VMDAREA,omitempty"`
	VMDAREA2                  string `json:"VMDAREA2,omitempty" url:"VMDAREA2,omitempty"`
	VMDAREA3                  string `json:"VMDAREA3,omitempty" url:"VMDAREA3,omitempty"`
	VMDAREA4                  string `json:"VMDAREA4,omitempty" url:"VMDAREA4,omitempty"`
	VMDSENSE                  string `json:"VMDSENSE,omitempty" url:"VMDSENSE,omitempty"`
	VMDSENSE2                 string `json:"VMDSENSE2,omitempty" url:"VMDSENSE2,omitempty"`
	VMDSENSE3                 string `json:"VMDSENSE3,omitempty" url:"VMDSENSE3,omitempty"`
	VMDSENSE4                 string `json:"VMDSENSE4,omitempty" url:"VMDSENSE4,omitempty"`
	VMDLIGHTC                 string `json:"VMDLIGHTC,omitempty" url:"VMDLIGHTC,omitempty"`
	VMDINFO                   string `json:"VMDINFO,omitempty" url:"VMDINFO,omitempty"`
	ORGUSE                    string `json:"ORGUSE,omitempty" url:"ORGUSE,omitempty"`
	ORGEXT                    string `json:"ORGEXT,omitempty" url:"ORGEXT,omitempty"`
	ORGPORT                   string `json:"ORGPORT,omitempty" url:"ORGPORT,omitempty"`
	ORGRTRY                   string `json:"ORGRTRY,omitempty" url:"ORGRTRY,omitempty"`
	ORGALM1                   string `json:"ORGALM1,omitempty" url:"ORGALM1,omitempty"`
	ORGALM2                   string `json:"ORGALM2,omitempty" url:"ORGALM2,omitempty"`
	ORGALM3                   string `json:"ORGALM3,omitempty" url:"ORGALM3,omitempty"`
	ORGALM4                   string `json:"ORGALM4,omitempty" url:"ORGALM4,omitempty"`
	ORGALM5                   string `json:"ORGALM5,omitempty" url:"ORGALM5,omitempty"`
	ORGALM6                   string `json:"ORGALM6,omitempty" url:"ORGALM6,omitempty"`
	ORGALM7                   string `json:"ORGALM7,omitempty" url:"ORGALM7,omitempty"`
	ORGALM8                   string `json:"ORGALM8,omitempty" url:"ORGALM8,omitempty"`
	ORGNOTICE1                string `json:"ORGNOTICE1,omitempty" url:"ORGNOTICE1,omitempty"`
	ORGNOTICE2                string `json:"ORGNOTICE2,omitempty" url:"ORGNOTICE2,omitempty"`
	ORGNOTICE3                string `json:"ORGNOTICE3,omitempty" url:"ORGNOTICE3,omitempty"`
	ORGNOTICE4                string `json:"ORGNOTICE4,omitempty" url:"ORGNOTICE4,omitempty"`
	ORGNOTICE5                string `json:"ORGNOTICE5,omitempty" url:"ORGNOTICE5,omitempty"`
	ORGNOTICE6                string `json:"ORGNOTICE6,omitempty" url:"ORGNOTICE6,omitempty"`
	ORGNOTICE7                string `json:"ORGNOTICE7,omitempty" url:"ORGNOTICE7,omitempty"`
	ORGNOTICE8                string `json:"ORGNOTICE8,omitempty" url:"ORGNOTICE8,omitempty"`
	ORGADD1                   string `json:"ORGADD1,omitempty" url:"ORGADD1,omitempty"`
	ORGADD2                   string `json:"ORGADD2,omitempty" url:"ORGADD2,omitempty"`
	ORGADD3                   string `json:"ORGADD3,omitempty" url:"ORGADD3,omitempty"`
	ORGADD4                   string `json:"ORGADD4,omitempty" url:"ORGADD4,omitempty"`
	ORGADD5                   string `json:"ORGADD5,omitempty" url:"ORGADD5,omitempty"`
	ORGADD6                   string `json:"ORGADD6,omitempty" url:"ORGADD6,omitempty"`
	ORGADD7                   string `json:"ORGADD7,omitempty" url:"ORGADD7,omitempty"`
	ORGADD8                   string `json:"ORGADD8,omitempty" url:"ORGADD8,omitempty"`
	ORGVMD1                   string `json:"ORGVMD1,omitempty" url:"ORGVMD1,omitempty"`
	ORGVMD2                   string `json:"ORGVMD2,omitempty" url:"ORGVMD2,omitempty"`
	ORGVMD3                   string `json:"ORGVMD3,omitempty" url:"ORGVMD3,omitempty"`
	ORGVMD4                   string `json:"ORGVMD4,omitempty" url:"ORGVMD4,omitempty"`
	ORGVMD5                   string `json:"ORGVMD5,omitempty" url:"ORGVMD5,omitempty"`
	ORGVMD6                   string `json:"ORGVMD6,omitempty" url:"ORGVMD6,omitempty"`
	ORGVMD7                   string `json:"ORGVMD7,omitempty" url:"ORGVMD7,omitempty"`
	ORGVMD8                   string `json:"ORGVMD8,omitempty" url:"ORGVMD8,omitempty"`
	ORGVMDAREA1               string `json:"ORGVMDAREA1,omitempty" url:"ORGVMDAREA1,omitempty"`
	ORGVMDAREA2               string `json:"ORGVMDAREA2,omitempty" url:"ORGVMDAREA2,omitempty"`
	ORGVMDAREA3               string `json:"ORGVMDAREA3,omitempty" url:"ORGVMDAREA3,omitempty"`
	ORGVMDAREA4               string `json:"ORGVMDAREA4,omitempty" url:"ORGVMDAREA4,omitempty"`
	ORGVMDAREA5               string `json:"ORGVMDAREA5,omitempty" url:"ORGVMDAREA5,omitempty"`
	ORGVMDAREA6               string `json:"ORGVMDAREA6,omitempty" url:"ORGVMDAREA6,omitempty"`
	ORGVMDAREA7               string `json:"ORGVMDAREA7,omitempty" url:"ORGVMDAREA7,omitempty"`
	ORGVMDAREA8               string `json:"ORGVMDAREA8,omitempty" url:"ORGVMDAREA8,omitempty"`
	HTTPALM                   string `json:"HTTPALM,omitempty" url:"HTTPALM,omitempty"`
	HTTPALM2                  string `json:"HTTPALM2,omitempty" url:"HTTPALM2,omitempty"`
	HTTPALM3                  string `json:"HTTPALM3,omitempty" url:"HTTPALM3,omitempty"`
	HTTPALM4                  string `json:"HTTPALM4,omitempty" url:"HTTPALM4,omitempty"`
	HTTPALM5                  string `json:"HTTPALM5,omitempty" url:"HTTPALM5,omitempty"`
	HTTPALMURL                string `json:"HTTPALMURL,omitempty" url:"HTTPALMURL,omitempty"`
	HTTPALMURL2               string `json:"HTTPALMURL2,omitempty" url:"HTTPALMURL2,omitempty"`
	HTTPALMURL3               string `json:"HTTPALMURL3,omitempty" url:"HTTPALMURL3,omitempty"`
	HTTPALMURL4               string `json:"HTTPALMURL4,omitempty" url:"HTTPALMURL4,omitempty"`
	HTTPALMURL5               string `json:"HTTPALMURL5,omitempty" url:"HTTPALMURL5,omitempty"`
	HTTPALMUSER               string `json:"HTTPALMUSER,omitempty" url:"HTTPALMUSER,omitempty"`
	HTTPALMUSER2              string `json:"HTTPALMUSER2,omitempty" url:"HTTPALMUSER2,omitempty"`
	HTTPALMUSER3              string `json:"HTTPALMUSER3,omitempty" url:"HTTPALMUSER3,omitempty"`
	HTTPALMUSER4              string `json:"HTTPALMUSER4,omitempty" url:"HTTPALMUSER4,omitempty"`
	HTTPALMUSER5              string `json:"HTTPALMUSER5,omitempty" url:"HTTPALMUSER5,omitempty"`
	UAUTH                     string `json:"UAUTH,omitempty" url:"UAUTH,omitempty"`
	UAUTHNOREG                string `json:"UAUTHNOREG,omitempty" url:"UAUTHNOREG,omitempty"`
	UAUTHMTD                  string `json:"UAUTHMTD,omitempty" url:"UAUTHMTD,omitempty"`
	UNAME                     string `json:"UNAME,omitempty" url:"UNAME,omitempty"`
	HAUTH                     string `json:"HAUTH,omitempty" url:"HAUTH,omitempty"`
	HADD                      string `json:"HADD,omitempty" url:"HADD,omitempty"`
	DHCP                      string `json:"DHCP,omitempty" url:"DHCP,omitempty"`
	NW                        string `json:"NW,omitempty" url:"NW,omitempty"`
	EIP1                      string `json:"EIP1,omitempty" url:"EIP1,omitempty"`
	EIP2                      string `json:"EIP2,omitempty" url:"EIP2,omitempty"`
	EIP3                      string `json:"EIP3,omitempty" url:"EIP3,omitempty"`
	EIP4                      string `json:"EIP4,omitempty" url:"EIP4,omitempty"`
	EMASK1                    string `json:"EMASK1,omitempty" url:"EMASK1,omitempty"`
	EMASK2                    string `json:"EMASK2,omitempty" url:"EMASK2,omitempty"`
	EMASK3                    string `json:"EMASK3,omitempty" url:"EMASK3,omitempty"`
	EMASK4                    string `json:"EMASK4,omitempty" url:"EMASK4,omitempty"`
	EDGW1                     string `json:"EDGW1,omitempty" url:"EDGW1,omitempty"`
	EDGW2                     string `json:"EDGW2,omitempty" url:"EDGW2,omitempty"`
	EDGW3                     string `json:"EDGW3,omitempty" url:"EDGW3,omitempty"`
	EDGW4                     string `json:"EDGW4,omitempty" url:"EDGW4,omitempty"`
	DNS                       string `json:"DNS,omitempty" url:"DNS,omitempty"`
	PRISRV1                   string `json:"PRISRV1,omitempty" url:"PRISRV1,omitempty"`
	PRISRV2                   string `json:"PRISRV2,omitempty" url:"PRISRV2,omitempty"`
	PRISRV3                   string `json:"PRISRV3,omitempty" url:"PRISRV3,omitempty"`
	PRISRV4                   string `json:"PRISRV4,omitempty" url:"PRISRV4,omitempty"`
	SECSRV1                   string `json:"SECSRV1,omitempty" url:"SECSRV1,omitempty"`
	SECSRV2                   string `json:"SECSRV2,omitempty" url:"SECSRV2,omitempty"`
	SECSRV3                   string `json:"SECSRV3,omitempty" url:"SECSRV3,omitempty"`
	SECSRV4                   string `json:"SECSRV4,omitempty" url:"SECSRV4,omitempty"`
	IP6_AUTO                  string `json:"IP6_AUTO,omitempty" url:"IP6_AUTO,omitempty"`
	IP6                       string `json:"IP6,omitempty" url:"IP6,omitempty"`
	IP6_DGW                   string `json:"IP6_DGW,omitempty" url:"IP6_DGW,omitempty"`
	IP6_DHCP                  string `json:"IP6_DHCP,omitempty" url:"IP6_DHCP,omitempty"`
	PRISRV_V6                 string `json:"PRISRV_V6,omitempty" url:"PRISRV_V6,omitempty"`
	SECSRV_V6                 string `json:"SECSRV_V6,omitempty" url:"SECSRV_V6,omitempty"`
	HTTPPORT                  string `json:"HTTPPORT,omitempty" url:"HTTPPORT,omitempty"`
	SPEED                     string `json:"SPEED,omitempty" url:"SPEED,omitempty"`
	RTPSIZE                   string `json:"RTPSIZE,omitempty" url:"RTPSIZE,omitempty"`
	MSS                       string `json:"MSS,omitempty" url:"MSS,omitempty"`
	BWC                       string `json:"BWC,omitempty" url:"BWC,omitempty"`
	EASYIPSETUP               string `json:"EASYIPSETUP,omitempty" url:"EASYIPSETUP,omitempty"`
	DHCP_HOST                 string `json:"DHCP_HOST,omitempty" url:"DHCP_HOST,omitempty"`
	MLUSE                     string `json:"MLUSE,omitempty" url:"MLUSE,omitempty"`
	MLSRV                     string `json:"MLSRV,omitempty" url:"MLSRV,omitempty"`
	MLPOPSRV                  string `json:"MLPOPSRV,omitempty" url:"MLPOPSRV,omitempty"`
	MLAUTH                    string `json:"MLAUTH,omitempty" url:"MLAUTH,omitempty"`
	MLUSER                    string `json:"MLUSER,omitempty" url:"MLUSER,omitempty"`
	MLFRM                     string `json:"MLFRM,omitempty" url:"MLFRM,omitempty"`
	SMTPPORT                  string `json:"SMTPPORT,omitempty" url:"SMTPPORT,omitempty"`
	MLALM1                    string `json:"MLALM1,omitempty" url:"MLALM1,omitempty"`
	MLALM2                    string `json:"MLALM2,omitempty" url:"MLALM2,omitempty"`
	MLALM3                    string `json:"MLALM3,omitempty" url:"MLALM3,omitempty"`
	MLALM4                    string `json:"MLALM4,omitempty" url:"MLALM4,omitempty"`
	MLNOTICE1                 string `json:"MLNOTICE1,omitempty" url:"MLNOTICE1,omitempty"`
	MLNOTICE2                 string `json:"MLNOTICE2,omitempty" url:"MLNOTICE2,omitempty"`
	MLNOTICE3                 string `json:"MLNOTICE3,omitempty" url:"MLNOTICE3,omitempty"`
	MLNOTICE4                 string `json:"MLNOTICE4,omitempty" url:"MLNOTICE4,omitempty"`
	MLSSL                     string `json:"MLSSL,omitempty" url:"MLSSL,omitempty"`
	MLTOADD1                  string `json:"MLTOADD1,omitempty" url:"MLTOADD1,omitempty"`
	MLTOADD2                  string `json:"MLTOADD2,omitempty" url:"MLTOADD2,omitempty"`
	MLTOADD3                  string `json:"MLTOADD3,omitempty" url:"MLTOADD3,omitempty"`
	MLTOADD4                  string `json:"MLTOADD4,omitempty" url:"MLTOADD4,omitempty"`
	MLSUBJECT                 string `json:"MLSUBJECT,omitempty" url:"MLSUBJECT,omitempty"`
	MLBODY                    string `json:"MLBODY,omitempty" url:"MLBODY,omitempty"`
	MLSUBJECTNOTICE           string `json:"MLSUBJECTNOTICE,omitempty" url:"MLSUBJECTNOTICE,omitempty"`
	TIMEADJUST                string `json:"TIMEADJUST,omitempty" url:"TIMEADJUST,omitempty"`
	NTPSVR                    string `json:"NTPSVR,omitempty" url:"NTPSVR,omitempty"`
	NTPADD                    string `json:"NTPADD,omitempty" url:"NTPADD,omitempty"`
	NTPPORT                   string `json:"NTPPORT,omitempty" url:"NTPPORT,omitempty"`
	NTPINTERVAL               string `json:"NTPINTERVAL,omitempty" url:"NTPINTERVAL,omitempty"`
	PORTFORWARD               string `json:"PORTFORWARD,omitempty" url:"PORTFORWARD,omitempty"`
	CAM_SC                    string `json:"CAM_SC,omitempty" url:"CAM_SC,omitempty"`
	UPNP_PORT                 string `json:"UPNP_PORT,omitempty" url:"UPNP_PORT,omitempty"`
	UPNP_STATUS               string `json:"UPNP_STATUS,omitempty" url:"UPNP_STATUS,omitempty"`
	UPNP_PORT_2               string `json:"UPNP_PORT_2,omitempty" url:"UPNP_PORT_2,omitempty"`
	UPNP_STATUS_2             string `json:"UPNP_STATUS_2,omitempty" url:"UPNP_STATUS_2,omitempty"`
	UPNP_ROUTER               string `json:"UPNP_ROUTER,omitempty" url:"UPNP_ROUTER,omitempty"`
	HTTPS_TYPE                string `json:"HTTPS_TYPE,omitempty" url:"HTTPS_TYPE,omitempty"`
	HTTPS_MIE_CAMURLSELFSTATE string `json:"HTTPS_MIE_CAMURLSELFSTATE,omitempty" url:"HTTPS_MIE_CAMURLSELFSTATE,omitempty"`
	HTTPS_CASTATE             string `json:"HTTPS_CASTATE,omitempty" url:"HTTPS_CASTATE,omitempty"`
	HTTPS_PORT                string `json:"HTTPS_PORT,omitempty" url:"HTTPS_PORT,omitempty"`
	HTTPS_SELFSTATE           string `json:"HTTPS_SELFSTATE,omitempty" url:"HTTPS_SELFSTATE,omitempty"`
	DDNS                      string `json:"DDNS,omitempty" url:"DDNS,omitempty"`
	DDHOST                    string `json:"DDHOST,omitempty" url:"DDHOST,omitempty"`
	DDINT                     string `json:"DDINT,omitempty" url:"DDINT,omitempty"`
	MIE_CAMURL                string `json:"MIE_CAMURL,omitempty" url:"MIE_CAMURL,omitempty"`
	MIE_STATUS                string `json:"MIE_STATUS,omitempty" url:"MIE_STATUS,omitempty"`
	SNMPCOM                   string `json:"SNMPCOM,omitempty" url:"SNMPCOM,omitempty"`
	SNMPTITLE                 string `json:"SNMPTITLE,omitempty" url:"SNMPTITLE,omitempty"`
	SNMPLOCATION              string `json:"SNMPLOCATION,omitempty" url:"SNMPLOCATION,omitempty"`
	SNMPCONTACT               string `json:"SNMPCONTACT,omitempty" url:"SNMPCONTACT,omitempty"`
	SNMPVER                   string `json:"SNMPVER,omitempty" url:"SNMPVER,omitempty"`
	SNMPV3_UNAME              string `json:"SNMPV3_UNAME,omitempty" url:"SNMPV3_UNAME,omitempty"`
	SNMPV3_AUTH               string `json:"SNMPV3_AUTH,omitempty" url:"SNMPV3_AUTH,omitempty"`
	SNMPV3_ENCR               string `json:"SNMPV3_ENCR,omitempty" url:"SNMPV3_ENCR,omitempty"`
	DSCP                      string `json:"DSCP,omitempty" url:"DSCP,omitempty"`
	DSCP_AUDIO                string `json:"DSCP_AUDIO,omitempty" url:"DSCP_AUDIO,omitempty"`
	DSCP_COM                  string `json:"DSCP_COM,omitempty" url:"DSCP_COM,omitempty"`
	SHAPING                   string `json:"SHAPING,omitempty" url:"SHAPING,omitempty"`
	SCHEFUNC1                 string `json:"SCHEFUNC1,omitempty" url:"SCHEFUNC1,omitempty"`
	SCHEFUNC2                 string `json:"SCHEFUNC2,omitempty" url:"SCHEFUNC2,omitempty"`
	SCHEFUNC3                 string `json:"SCHEFUNC3,omitempty" url:"SCHEFUNC3,omitempty"`
	SCHEFUNC4                 string `json:"SCHEFUNC4,omitempty" url:"SCHEFUNC4,omitempty"`
	SCHEFUNC5                 string `json:"SCHEFUNC5,omitempty" url:"SCHEFUNC5,omitempty"`
	EXT1_ID                   string `json:"EXT1_ID,omitempty" url:"EXT1_ID,omitempty"`
	EXT1_MON                  string `json:"EXT1_MON,omitempty" url:"EXT1_MON,omitempty"`
	EXT1_TUE                  string `json:"EXT1_TUE,omitempty" url:"EXT1_TUE,omitempty"`
	EXT1_WED                  string `json:"EXT1_WED,omitempty" url:"EXT1_WED,omitempty"`
	EXT1_THU                  string `json:"EXT1_THU,omitempty" url:"EXT1_THU,omitempty"`
	EXT1_FRI                  string `json:"EXT1_FRI,omitempty" url:"EXT1_FRI,omitempty"`
	EXT1_SAT                  string `json:"EXT1_SAT,omitempty" url:"EXT1_SAT,omitempty"`
	EXT1_SUN                  string `json:"EXT1_SUN,omitempty" url:"EXT1_SUN,omitempty"`
	EXT1_TIME1_SCHE1_START_H  string `json:"EXT1_TIME1_SCHE1_START_H,omitempty" url:"EXT1_TIME1_SCHE1_START_H,omitempty"`
	EXT1_TIME1_SCHE1_START_M  string `json:"EXT1_TIME1_SCHE1_START_M,omitempty" url:"EXT1_TIME1_SCHE1_START_M,omitempty"`
	EXT1_TIME1_SCHE1_END_H    string `json:"EXT1_TIME1_SCHE1_END_H,omitempty" url:"EXT1_TIME1_SCHE1_END_H,omitempty"`
	EXT1_TIME1_SCHE1_END_M    string `json:"EXT1_TIME1_SCHE1_END_M,omitempty" url:"EXT1_TIME1_SCHE1_END_M,omitempty"`
	EXT1_TIME1_SCHE1_MODE     string `json:"EXT1_TIME1_SCHE1_MODE,omitempty" url:"EXT1_TIME1_SCHE1_MODE,omitempty"`
	EXT1_TIME2_SCHE1_START_H  string `json:"EXT1_TIME2_SCHE1_START_H,omitempty" url:"EXT1_TIME2_SCHE1_START_H,omitempty"`
	EXT1_TIME2_SCHE1_START_M  string `json:"EXT1_TIME2_SCHE1_START_M,omitempty" url:"EXT1_TIME2_SCHE1_START_M,omitempty"`
	EXT1_TIME2_SCHE1_END_H    string `json:"EXT1_TIME2_SCHE1_END_H,omitempty" url:"EXT1_TIME2_SCHE1_END_H,omitempty"`
	EXT1_TIME2_SCHE1_END_M    string `json:"EXT1_TIME2_SCHE1_END_M,omitempty" url:"EXT1_TIME2_SCHE1_END_M,omitempty"`
	EXT1_TIME2_SCHE1_MODE     string `json:"EXT1_TIME2_SCHE1_MODE,omitempty" url:"EXT1_TIME2_SCHE1_MODE,omitempty"`
	EXT1_TIME1_SCHE2_START_H  string `json:"EXT1_TIME1_SCHE2_START_H,omitempty" url:"EXT1_TIME1_SCHE2_START_H,omitempty"`
	EXT1_TIME1_SCHE2_START_M  string `json:"EXT1_TIME1_SCHE2_START_M,omitempty" url:"EXT1_TIME1_SCHE2_START_M,omitempty"`
	EXT1_TIME1_SCHE2_END_H    string `json:"EXT1_TIME1_SCHE2_END_H,omitempty" url:"EXT1_TIME1_SCHE2_END_H,omitempty"`
	EXT1_TIME1_SCHE2_END_M    string `json:"EXT1_TIME1_SCHE2_END_M,omitempty" url:"EXT1_TIME1_SCHE2_END_M,omitempty"`
	EXT1_TIME1_SCHE2_MODE     string `json:"EXT1_TIME1_SCHE2_MODE,omitempty" url:"EXT1_TIME1_SCHE2_MODE,omitempty"`
	EXT1_TIME2_SCHE2_START_H  string `json:"EXT1_TIME2_SCHE2_START_H,omitempty" url:"EXT1_TIME2_SCHE2_START_H,omitempty"`
	EXT1_TIME2_SCHE2_START_M  string `json:"EXT1_TIME2_SCHE2_START_M,omitempty" url:"EXT1_TIME2_SCHE2_START_M,omitempty"`
	EXT1_TIME2_SCHE2_END_H    string `json:"EXT1_TIME2_SCHE2_END_H,omitempty" url:"EXT1_TIME2_SCHE2_END_H,omitempty"`
	EXT1_TIME2_SCHE2_END_M    string `json:"EXT1_TIME2_SCHE2_END_M,omitempty" url:"EXT1_TIME2_SCHE2_END_M,omitempty"`
	EXT1_TIME2_SCHE2_MODE     string `json:"EXT1_TIME2_SCHE2_MODE,omitempty" url:"EXT1_TIME2_SCHE2_MODE,omitempty"`
	EXT1_TIME1_SCHE3_START_H  string `json:"EXT1_TIME1_SCHE3_START_H,omitempty" url:"EXT1_TIME1_SCHE3_START_H,omitempty"`
	EXT1_TIME1_SCHE3_START_M  string `json:"EXT1_TIME1_SCHE3_START_M,omitempty" url:"EXT1_TIME1_SCHE3_START_M,omitempty"`
	EXT1_TIME1_SCHE3_END_H    string `json:"EXT1_TIME1_SCHE3_END_H,omitempty" url:"EXT1_TIME1_SCHE3_END_H,omitempty"`
	EXT1_TIME1_SCHE3_END_M    string `json:"EXT1_TIME1_SCHE3_END_M,omitempty" url:"EXT1_TIME1_SCHE3_END_M,omitempty"`
	EXT1_TIME1_SCHE3_MODE     string `json:"EXT1_TIME1_SCHE3_MODE,omitempty" url:"EXT1_TIME1_SCHE3_MODE,omitempty"`
	EXT1_TIME2_SCHE3_START_H  string `json:"EXT1_TIME2_SCHE3_START_H,omitempty" url:"EXT1_TIME2_SCHE3_START_H,omitempty"`
	EXT1_TIME2_SCHE3_START_M  string `json:"EXT1_TIME2_SCHE3_START_M,omitempty" url:"EXT1_TIME2_SCHE3_START_M,omitempty"`
	EXT1_TIME2_SCHE3_END_H    string `json:"EXT1_TIME2_SCHE3_END_H,omitempty" url:"EXT1_TIME2_SCHE3_END_H,omitempty"`
	EXT1_TIME2_SCHE3_END_M    string `json:"EXT1_TIME2_SCHE3_END_M,omitempty" url:"EXT1_TIME2_SCHE3_END_M,omitempty"`
	EXT1_TIME2_SCHE3_MODE     string `json:"EXT1_TIME2_SCHE3_MODE,omitempty" url:"EXT1_TIME2_SCHE3_MODE,omitempty"`
	EXT1_TIME1_SCHE4_START_H  string `json:"EXT1_TIME1_SCHE4_START_H,omitempty" url:"EXT1_TIME1_SCHE4_START_H,omitempty"`
	EXT1_TIME1_SCHE4_START_M  string `json:"EXT1_TIME1_SCHE4_START_M,omitempty" url:"EXT1_TIME1_SCHE4_START_M,omitempty"`
	EXT1_TIME1_SCHE4_END_H    string `json:"EXT1_TIME1_SCHE4_END_H,omitempty" url:"EXT1_TIME1_SCHE4_END_H,omitempty"`
	EXT1_TIME1_SCHE4_END_M    string `json:"EXT1_TIME1_SCHE4_END_M,omitempty" url:"EXT1_TIME1_SCHE4_END_M,omitempty"`
	EXT1_TIME1_SCHE4_MODE     string `json:"EXT1_TIME1_SCHE4_MODE,omitempty" url:"EXT1_TIME1_SCHE4_MODE,omitempty"`
	EXT1_TIME2_SCHE4_START_H  string `json:"EXT1_TIME2_SCHE4_START_H,omitempty" url:"EXT1_TIME2_SCHE4_START_H,omitempty"`
	EXT1_TIME2_SCHE4_START_M  string `json:"EXT1_TIME2_SCHE4_START_M,omitempty" url:"EXT1_TIME2_SCHE4_START_M,omitempty"`
	EXT1_TIME2_SCHE4_END_H    string `json:"EXT1_TIME2_SCHE4_END_H,omitempty" url:"EXT1_TIME2_SCHE4_END_H,omitempty"`
	EXT1_TIME2_SCHE4_END_M    string `json:"EXT1_TIME2_SCHE4_END_M,omitempty" url:"EXT1_TIME2_SCHE4_END_M,omitempty"`
	EXT1_TIME2_SCHE4_MODE     string `json:"EXT1_TIME2_SCHE4_MODE,omitempty" url:"EXT1_TIME2_SCHE4_MODE,omitempty"`
	EXT1_TIME1_SCHE5_START_H  string `json:"EXT1_TIME1_SCHE5_START_H,omitempty" url:"EXT1_TIME1_SCHE5_START_H,omitempty"`
	EXT1_TIME1_SCHE5_START_M  string `json:"EXT1_TIME1_SCHE5_START_M,omitempty" url:"EXT1_TIME1_SCHE5_START_M,omitempty"`
	EXT1_TIME1_SCHE5_END_H    string `json:"EXT1_TIME1_SCHE5_END_H,omitempty" url:"EXT1_TIME1_SCHE5_END_H,omitempty"`
	EXT1_TIME1_SCHE5_END_M    string `json:"EXT1_TIME1_SCHE5_END_M,omitempty" url:"EXT1_TIME1_SCHE5_END_M,omitempty"`
	EXT1_TIME1_SCHE5_MODE     string `json:"EXT1_TIME1_SCHE5_MODE,omitempty" url:"EXT1_TIME1_SCHE5_MODE,omitempty"`
	EXT1_TIME2_SCHE5_START_H  string `json:"EXT1_TIME2_SCHE5_START_H,omitempty" url:"EXT1_TIME2_SCHE5_START_H,omitempty"`
	EXT1_TIME2_SCHE5_START_M  string `json:"EXT1_TIME2_SCHE5_START_M,omitempty" url:"EXT1_TIME2_SCHE5_START_M,omitempty"`
	EXT1_TIME2_SCHE5_END_H    string `json:"EXT1_TIME2_SCHE5_END_H,omitempty" url:"EXT1_TIME2_SCHE5_END_H,omitempty"`
	EXT1_TIME2_SCHE5_END_M    string `json:"EXT1_TIME2_SCHE5_END_M,omitempty" url:"EXT1_TIME2_SCHE5_END_M,omitempty"`
	EXT1_TIME2_SCHE5_MODE     string `json:"EXT1_TIME2_SCHE5_MODE,omitempty" url:"EXT1_TIME2_SCHE5_MODE,omitempty"`
	EXT1_TIME1_SCHE6_START_H  string `json:"EXT1_TIME1_SCHE6_START_H,omitempty" url:"EXT1_TIME1_SCHE6_START_H,omitempty"`
	EXT1_TIME1_SCHE6_START_M  string `json:"EXT1_TIME1_SCHE6_START_M,omitempty" url:"EXT1_TIME1_SCHE6_START_M,omitempty"`
	EXT1_TIME1_SCHE6_END_H    string `json:"EXT1_TIME1_SCHE6_END_H,omitempty" url:"EXT1_TIME1_SCHE6_END_H,omitempty"`
	EXT1_TIME1_SCHE6_END_M    string `json:"EXT1_TIME1_SCHE6_END_M,omitempty" url:"EXT1_TIME1_SCHE6_END_M,omitempty"`
	EXT1_TIME1_SCHE6_MODE     string `json:"EXT1_TIME1_SCHE6_MODE,omitempty" url:"EXT1_TIME1_SCHE6_MODE,omitempty"`
	EXT1_TIME2_SCHE6_START_H  string `json:"EXT1_TIME2_SCHE6_START_H,omitempty" url:"EXT1_TIME2_SCHE6_START_H,omitempty"`
	EXT1_TIME2_SCHE6_START_M  string `json:"EXT1_TIME2_SCHE6_START_M,omitempty" url:"EXT1_TIME2_SCHE6_START_M,omitempty"`
	EXT1_TIME2_SCHE6_END_H    string `json:"EXT1_TIME2_SCHE6_END_H,omitempty" url:"EXT1_TIME2_SCHE6_END_H,omitempty"`
	EXT1_TIME2_SCHE6_END_M    string `json:"EXT1_TIME2_SCHE6_END_M,omitempty" url:"EXT1_TIME2_SCHE6_END_M,omitempty"`
	EXT1_TIME2_SCHE6_MODE     string `json:"EXT1_TIME2_SCHE6_MODE,omitempty" url:"EXT1_TIME2_SCHE6_MODE,omitempty"`
	SMARTCODING               string `json:"SMARTCODING,omitempty" url:"SMARTCODING,omitempty"`
	SMARTCODING_2             string `json:"SMARTCODING_2,omitempty" url:"SMARTCODING_2,omitempty"`
	SMARTCODING_3             string `json:"SMARTCODING_3,omitempty" url:"SMARTCODING_3,omitempty"`
	SMART_VMDSENSE            string `json:"SMART_VMDSENSE,omitempty" url:"SMART_VMDSENSE,omitempty"`
	SMART_VMDSENSE2           string `json:"SMART_VMDSENSE2,omitempty" url:"SMART_VMDSENSE2,omitempty"`
	SMART_VMDSENSE3           string `json:"SMART_VMDSENSE3,omitempty" url:"SMART_VMDSENSE3,omitempty"`
	SMART_RINT                string `json:"SMART_RINT,omitempty" url:"SMART_RINT,omitempty"`
	SMART_RINT_2              string `json:"SMART_RINT_2,omitempty" url:"SMART_RINT_2,omitempty"`
	SMART_RINT_3              string `json:"SMART_RINT_3,omitempty" url:"SMART_RINT_3,omitempty"`
	SMART_FORCE_IDR           string `json:"SMART_FORCE_IDR,omitempty" url:"SMART_FORCE_IDR,omitempty"`
	SMART_FORCE_IDR2          string `json:"SMART_FORCE_IDR2,omitempty" url:"SMART_FORCE_IDR2,omitempty"`
	SMART_FORCE_IDR3          string `json:"SMART_FORCE_IDR3,omitempty" url:"SMART_FORCE_IDR3,omitempty"`
	SMART_RINT_MIN            string `json:"SMART_RINT_MIN,omitempty" url:"SMART_RINT_MIN,omitempty"`
	SMART_RINT_MIN_2          string `json:"SMART_RINT_MIN_2,omitempty" url:"SMART_RINT_MIN_2,omitempty"`
	SMART_RINT_MIN_3          string `json:"SMART_RINT_MIN_3,omitempty" url:"SMART_RINT_MIN_3,omitempty"`
	SMART_LEVEL               string `json:"SMART_LEVEL,omitempty" url:"SMART_LEVEL,omitempty"`
	SMART_LEVEL_2             string `json:"SMART_LEVEL_2,omitempty" url:"SMART_LEVEL_2,omitempty"`
	SMART_LEVEL_3             string `json:"SMART_LEVEL_3,omitempty" url:"SMART_LEVEL_3,omitempty"`
	SMART_LTINT               string `json:"SMART_LTINT,omitempty" url:"SMART_LTINT,omitempty"`
	SMART_LTINT_2             string `json:"SMART_LTINT_2,omitempty" url:"SMART_LTINT_2,omitempty"`
	SMART_LTINT_3             string `json:"SMART_LTINT_3,omitempty" url:"SMART_LTINT_3,omitempty"`
	SMART_VIQS                string `json:"SMART_VIQS,omitempty" url:"SMART_VIQS,omitempty"`
	FRAMERATE_MODE            string `json:"FRAMERATE_MODE,omitempty" url:"FRAMERATE_MODE,omitempty"`
	PPIC_FORCE_IDR            string `json:"PPIC_FORCE_IDR,omitempty" url:"PPIC_FORCE_IDR,omitempty"`
	PPIC_MAXSIZE              string `json:"PPIC_MAXSIZE,omitempty" url:"PPIC_MAXSIZE,omitempty"`
	SCD_ALMID                 string `json:"SCD_ALMID,omitempty" url:"SCD_ALMID,omitempty"`
	IVMD_EXT                  string `json:"IVMD_EXT,omitempty" url:"IVMD_EXT,omitempty"`
	SCDULX1                   string `json:"SCDULX1,omitempty" url:"SCDULX1,omitempty"`
	SCDULY1                   string `json:"SCDULY1,omitempty" url:"SCDULY1,omitempty"`
	SCDBRX1                   string `json:"SCDBRX1,omitempty" url:"SCDBRX1,omitempty"`
	SCDBRY1                   string `json:"SCDBRY1,omitempty" url:"SCDBRY1,omitempty"`
	SCDSTATUS1                string `json:"SCDSTATUS1,omitempty" url:"SCDSTATUS1,omitempty"`
	SCDAREA                   string `json:"SCDAREA,omitempty" url:"SCDAREA,omitempty"`
	SCDSENSE                  string `json:"SCDSENSE,omitempty" url:"SCDSENSE,omitempty"`
	SCDTIME                   string `json:"SCDTIME,omitempty" url:"SCDTIME,omitempty"`
	SCDINFO                   string `json:"SCDINFO,omitempty" url:"SCDINFO,omitempty"`
	UDSEI                     string `json:"UDSEI,omitempty" url:"UDSEI,omitempty"`
	LLDP                      string `json:"LLDP,omitempty" url:"LLDP,omitempty"`
	REG_TLS1_1                string `json:"REG_TLS1_1,omitempty" url:"REG_TLS1_1,omitempty"`
	EXT_RAMEXTEND             string `json:"EXT_RAMEXTEND,omitempty" url:"EXT_RAMEXTEND,omitempty"`
	MQTT                      string `json:"MQTT,omitempty" url:"MQTT,omitempty"`
	MQTT_SERVER_ADDR          string `json:"MQTT_SERVER_ADDR,omitempty" url:"MQTT_SERVER_ADDR,omitempty"`
	MQTT_SERVER_PORT          string `json:"MQTT_SERVER_PORT,omitempty" url:"MQTT_SERVER_PORT,omitempty"`
	MQTT_PROTOCOL             string `json:"MQTT_PROTOCOL,omitempty" url:"MQTT_PROTOCOL,omitempty"`
	MQTT_USERNAME             string `json:"MQTT_USERNAME,omitempty" url:"MQTT_USERNAME,omitempty"`
	MQTT_TERMINAL1            string `json:"MQTT_TERMINAL1,omitempty" url:"MQTT_TERMINAL1,omitempty"`
	MQTT_TERMINAL1_TOPIC      string `json:"MQTT_TERMINAL1_TOPIC,omitempty" url:"MQTT_TERMINAL1_TOPIC,omitempty"`
	MQTT_TERMINAL1_PAYLOAD    string `json:"MQTT_TERMINAL1_PAYLOAD,omitempty" url:"MQTT_TERMINAL1_PAYLOAD,omitempty"`
	MQTT_TERMINAL1_RETAIN     string `json:"MQTT_TERMINAL1_RETAIN,omitempty" url:"MQTT_TERMINAL1_RETAIN,omitempty"`
	MQTT_TERMINAL1_QOS        string `json:"MQTT_TERMINAL1_QOS,omitempty" url:"MQTT_TERMINAL1_QOS,omitempty"`
	MQTT_TERMINAL2            string `json:"MQTT_TERMINAL2,omitempty" url:"MQTT_TERMINAL2,omitempty"`
	MQTT_TERMINAL2_TOPIC      string `json:"MQTT_TERMINAL2_TOPIC,omitempty" url:"MQTT_TERMINAL2_TOPIC,omitempty"`
	MQTT_TERMINAL2_PAYLOAD    string `json:"MQTT_TERMINAL2_PAYLOAD,omitempty" url:"MQTT_TERMINAL2_PAYLOAD,omitempty"`
	MQTT_TERMINAL2_RETAIN     string `json:"MQTT_TERMINAL2_RETAIN,omitempty" url:"MQTT_TERMINAL2_RETAIN,omitempty"`
	MQTT_TERMINAL2_QOS        string `json:"MQTT_TERMINAL2_QOS,omitempty" url:"MQTT_TERMINAL2_QOS,omitempty"`
	MQTT_TERMINAL3            string `json:"MQTT_TERMINAL3,omitempty" url:"MQTT_TERMINAL3,omitempty"`
	MQTT_TERMINAL3_TOPIC      string `json:"MQTT_TERMINAL3_TOPIC,omitempty" url:"MQTT_TERMINAL3_TOPIC,omitempty"`
	MQTT_TERMINAL3_PAYLOAD    string `json:"MQTT_TERMINAL3_PAYLOAD,omitempty" url:"MQTT_TERMINAL3_PAYLOAD,omitempty"`
	MQTT_TERMINAL3_RETAIN     string `json:"MQTT_TERMINAL3_RETAIN,omitempty" url:"MQTT_TERMINAL3_RETAIN,omitempty"`
	MQTT_TERMINAL3_QOS        string `json:"MQTT_TERMINAL3_QOS,omitempty" url:"MQTT_TERMINAL3_QOS,omitempty"`
	MQTT_VMD                  string `json:"MQTT_VMD,omitempty" url:"MQTT_VMD,omitempty"`
	MQTT_VMD_TOPIC            string `json:"MQTT_VMD_TOPIC,omitempty" url:"MQTT_VMD_TOPIC,omitempty"`
	MQTT_VMD_PAYLOAD          string `json:"MQTT_VMD_PAYLOAD,omitempty" url:"MQTT_VMD_PAYLOAD,omitempty"`
	MQTT_VMD_QOS              string `json:"MQTT_VMD_QOS,omitempty" url:"MQTT_VMD_QOS,omitempty"`
	MQTT_VMD_RETAIN           string `json:"MQTT_VMD_RETAIN,omitempty" url:"MQTT_VMD_RETAIN,omitempty"`
	MQTT_AUDIO                string `json:"MQTT_AUDIO,omitempty" url:"MQTT_AUDIO,omitempty"`
	MQTT_AUDIO_TOPIC          string `json:"MQTT_AUDIO_TOPIC,omitempty" url:"MQTT_AUDIO_TOPIC,omitempty"`
	MQTT_AUDIO_PAYLOAD        string `json:"MQTT_AUDIO_PAYLOAD,omitempty" url:"MQTT_AUDIO_PAYLOAD,omitempty"`
	MQTT_AUDIO_QOS            string `json:"MQTT_AUDIO_QOS,omitempty" url:"MQTT_AUDIO_QOS,omitempty"`
	MQTT_AUDIO_RETAIN         string `json:"MQTT_AUDIO_RETAIN,omitempty" url:"MQTT_AUDIO_RETAIN,omitempty"`
	MQTT_CMD                  string `json:"MQTT_CMD,omitempty" url:"MQTT_CMD,omitempty"`
	MQTT_CMD_TOPIC            string `json:"MQTT_CMD_TOPIC,omitempty" url:"MQTT_CMD_TOPIC,omitempty"`
	MQTT_CMD_PAYLOAD          string `json:"MQTT_CMD_PAYLOAD,omitempty" url:"MQTT_CMD_PAYLOAD,omitempty"`
	MQTT_CMD_QOS              string `json:"MQTT_CMD_QOS,omitempty" url:"MQTT_CMD_QOS,omitempty"`
	MQTT_CMD_RETAIN           string `json:"MQTT_CMD_RETAIN,omitempty" url:"MQTT_CMD_RETAIN,omitempty"`
	MQTT_VERIFY_CERT          string `json:"MQTT_VERIFY_CERT,omitempty" url:"MQTT_VERIFY_CERT,omitempty"`
	STRATOCAST                string `json:"STRATOCAST,omitempty" url:"STRATOCAST,omitempty"`
	LIVEPAGE_COLOR            string `json:"LIVEPAGE_COLOR,omitempty" url:"LIVEPAGE_COLOR,omitempty"`
	LIVEPAGE_PANELLAYOUT      string `json:"LIVEPAGE_PANELLAYOUT,omitempty" url:"LIVEPAGE_PANELLAYOUT,omitempty"`
	VIDEO_SERVER_IMAGE_MODE   string `json:"video_server.image.mode,omitempty" url:"video_server.image.mode,omitempty"`
	SDNAME                    string `json:"SDNAME,omitempty" url:"SDNAME,omitempty"`
	SDSERIAL                  string `json:"SDSERIAL,omitempty" url:"SDSERIAL,omitempty"`
	SDTOTALTIME               string `json:"SDTOTALTIME,omitempty" url:"SDTOTALTIME,omitempty"`
	SDOVERWRITE               string `json:"SDOVERWRITE,omitempty" url:"SDOVERWRITE,omitempty"`
}
