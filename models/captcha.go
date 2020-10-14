package models

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

//CaptchaBody json request body.
type CaptchaBody struct {
	ID            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// GenerateCaptcha 生成验证码
func (e *CaptchaBody) GenerateCaptcha() (string, string, error) {
	var driver base64Captcha.Driver
	//create base64 encoding captcha
	switch e.CaptchaType {
	case "audio":
		e.DriverAudio = base64Captcha.DefaultDriverAudio
		driver = e.DriverAudio
	case "string":
		e.DriverString = base64Captcha.NewDriverString(46, 140, 2, 2, 4, "234567890abcdefghjkmnpqrstuvwxyz", &color.RGBA{240, 240, 246, 246}, []string{})
		driver = e.DriverString.ConvertFonts()
	case "math":
		e.DriverMath = base64Captcha.NewDriverMath(46, 140, 2, 2, &color.RGBA{240, 240, 246, 246}, []string{})
		driver = e.DriverMath.ConvertFonts()
	case "chinese":
		driver = e.DriverChinese.ConvertFonts()
	default:
		e.DriverDigit = base64Captcha.DefaultDriverDigit
		driver = e.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	return id, b64s, err
}

// VerifyCaptcha 验证码验证
func (e *CaptchaBody) VerifyCaptcha(id string, code string) bool {
	return store.Verify(id, code, true)
}
