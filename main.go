package main

import (
	"fmt"
	"os"

	"github.com/Tencent-YouTu/Go_sdk"
	"github.com/satori/go.uuid"
	"github.com/labstack/echo"
	"net/http"
	"strings"
	"errors"
	"encoding/base64"
	"github.com/labstack/echo/middleware"
)

func main() {
	//Register your app on http://open.youtu.qq.com
	//Get the following details
	appID := uint32(0)
	secretID := ""
	secretKey := ""
	userID := uuid.NewV4().String()

	as, err := youtu.NewAppSign(appID, secretID, secretKey, userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewAppSign() failed: %s\n", err)
		return
	}

	//imgData, err := ioutil.ReadFile("test.jpeg")
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "ReadFile() failed: %s\n", err)
	//	return
	//}

	//yt := youtu.Init(as, youtu.TencentYunHost)
	yt := youtu.Init(as, youtu.DefaultHost)
	e := echo.New()

	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://detect.vdo.pub", "http://localhost:8020","http://127.0.0.1:8020"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.POST("/api/detect_face", func(c echo.Context) error {

		face:=c.FormValue("face")

		if len(strings.TrimSpace(face))==0 {
			return errors.New("This face data is not matched!")
		}
		imgData, err :=base64.StdEncoding.DecodeString(face)
		if err != nil {
			return err
		}

		df, err := yt.DetectFace(imgData, false, 0)

		if err != nil {
			fmt.Fprintf(os.Stderr, "DetectFace() failed: %s", err)
			return err
		}
		fmt.Printf("df: %#v\n", df)
		return c.JSON(http.StatusOK, df)
	})
	e.Logger.Fatal(e.Start(":1323"))

}
