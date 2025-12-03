package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/SridarDhandapani/go-ipro/pkg/api"
	"github.com/SridarDhandapani/go-ipro/pkg/camera"
	"github.com/SridarDhandapani/go-ipro/pkg/enum"
	ce "github.com/SridarDhandapani/go-ipro/pkg/error"
	"github.com/SridarDhandapani/go-ipro/pkg/model"
)

func main() {
	var err error

	c, err := camera.NewCamera(
		camera.WithName("XDV48057"),
		camera.WithUsername("admin"),
		camera.WithPassword("H0xt0nA!"),
		camera.WithAPIOpts(
			api.WithHost("192.168.50.109"),
			api.WithScheme(api.HTTP),
			api.WithPort(80),
			api.WithEndpointPrefix("/cgi-bin"),
		),
	)
	if err != nil {
		panic(err)
	}

	admin := &model.User{Name: "admin", Password: "H0xt0nA!"}

	err = c.RegisterAdmin(admin)(c.APIHandler)
	if err != nil && !errors.Is(err, ce.ErrAlreadyInitialised) {
		panic(err)
	}

	user := &model.User{Name: "liveonly", Password: "L1v30nly", AccessLevel: enum.ACLLiveOnly}
	err = c.RegisterUser(user)(c.APIHandler)
	if err != nil {
		fmt.Println(err)
	}
	//
	// //_ = c.DeleteUser("liveonly", token)(c.APIHandler)
	//
	// newTime := time.Now().Add(2 * 24 * time.Hour)
	// newTime := time.Now().UTC()
	// timeData := model.CameraData{
	// 	TIMEDATE: newTime.Format("2006,01,02,15,04,05"),
	// 	TIMEDISP: enum.TimeDisplay24H,
	// 	OSDSIZE:  enum.OsdSizeSmall,
	// }
	// err = validator.Validate(timeData)
	// //_ = c.SetData(timeData)(c.APIHandler)
	//
	clock := model.Clock{
		TimeDate:      time.Now().Format("2006,01,02,15,04,05"),
		Timezone:      26,
		TimeDisplay:   enum.TimeDisplay24H,
		DisplayFormat: 1,
	}
	err = c.SetData(clock)(c.APIHandler)
	if err != nil {
		fmt.Println(err)
	}
	// err = validator.Validate(clock)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// c, _ := query.Values(clock)
	// fmt.Println(c.Encode())
	//
	title := model.Title{
		CamTitle:          "XDV48059",
		CamIDDisplay:      enum.Off,
		CamID:             "XDV48059",
		CamIDDisplayPlace: enum.OsdPositionUpperLeft,
		DisplaySize:       enum.OsdSizeXSmall,
	}
	err = c.SetData(title)(c.APIHandler)
	if err != nil {
		fmt.Println(err)
	}

	// err = c.Reset("all")(c.APIHandler)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	m := &model.CaptureMode{
		Ratio: "4_3",
		Fps:   "30",
	}
	err = c.SetImageCaptureMode(m)(c.APIHandler)
	if err != nil {
		fmt.Println(err)
	}

	s := &model.Stream{
		Transmit:   enum.On,
		EncodeType: enum.StreamEncodeTypeH264,
		Resolution: enum.StreamResolutionVGA,
		FrameRate:  enum.StreamFrameRate10fps,
	}
	err = c.SetStreamConfig(1, s)(c.APIHandler)
	if err != nil {
		fmt.Println(err)
	}

	err = c.SetStreamConfig(2, &model.Stream{Transmit: enum.Off})(c.APIHandler)
	if err != nil {
		fmt.Println(err)
	}
	err = c.SetStreamConfig(3, &model.Stream{Transmit: enum.Off})(c.APIHandler)
	if err != nil {
		fmt.Println(err)
	}
	err = c.SetStreamConfig(4, &model.Stream{Transmit: enum.Off})(c.APIHandler)
	if err != nil {
		fmt.Println(err)
	}

	info, _ := c.GetInfo()(c.APIHandler)
	fmt.Printf("Info: %v\n", info)
	fmt.Printf("SN: %s\n", info.Serial)

	data, _ := c.GetData()(c.APIHandler)
	fmt.Printf("CamTitle: %v\n", data.CAMTITLE)
	fmt.Printf("Data: %v\n", data)
}
