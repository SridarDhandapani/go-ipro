package model

import "github.com/SridarDhandapani/go-ipro/pkg/enum"

type Stream struct {
	Transmit         enum.OnOff                  `url:"h264_transmit,omitempty" validate:"regexp=^[01]*$"`
	EncodeType       enum.StreamEncodeType       `url:"encode_type,omitempty" validate:"regexp=^[12]*$"`
	Resolution       enum.StreamResolution       `url:"h264_resolution,omitempty" validate:"regexp=^(320|400|640|800|1280|1600|1920|2048|3072|2560|3840|2192|2992)*$"`
	Priority         enum.StreamPriority         `url:"f_priority,omitempty" validate:"regexp=^[0124]*$"`
	FrameRate        enum.StreamFrameRate        `url:"nr_framerate,omitempty" validate:"regexp=^(1|3|5|7\\.5|10|12|15|20|30|60)*$"`
	Bandwidth        enum.StreamBandwidth        `url:"h264_bandwidth,omitempty" validate:"regexp=^(64|128|256|384|512|768|1024|1536|2048|3072|4096|6144|8192|10240|12288|14336|16384|20480|24576)*$"`
	Quality          enum.StreamQuality          `url:"h264_quality,omitempty" validate:"regexp=^(fine|normal|low|[0-9])*$"`
	TransmissionType enum.StreamTransmissionType `url:"h264_unimulti,omitempty" validate:"regexp=^(uni|multi|uni_manual)*$"`
	MulticastAddr    string                      `url:"multicast_addr,omitempty" validate:"regexp=^((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)?)*$"`
	MulticastPort    string                      `url:"multicast_port,omitempty" validate:"regexp=^(10(2[4-9]|[3-9][0-9])|[2-9][0-9]{3}|[1-4][0-9]{4}|50000)*$"`
	MulticastTTL     string                      `url:"multicast_ttl,omitempty" validate:"regexp=^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])*$"`
}
