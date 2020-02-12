package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/mojocn/base64Captcha"
	"github.com/mszsgo/hjson"
)

type CaptchaNumberBase64Type struct {
	CaptchaId   string `description:"验证码编号，验证时需要提供"`
	Base64Image string `description:"验证码Base64编码图片"`
}

func (*CaptchaNumberBase64Type) Description() string {
	return "数字验证码对象"
}

type CaptchaNumberBase64TypeArgs struct {
	Len    int `defaultValue:"6" description:"数字验证码长度"`
	Width  int `defaultValue:"240" description:"验证码图片宽度，px"`
	Height int `defaultValue:"80" description:"验证码图片高度，px"`
}

func (*CaptchaNumberBase64Type) Args() *CaptchaNumberBase64TypeArgs {
	return &CaptchaNumberBase64TypeArgs{}
}

func (*CaptchaNumberBase64Type) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *CaptchaNumberBase64TypeArgs
		hjson.MapToStruct(p.Args, &args)

		// 数字验证码配置
		var configD = base64Captcha.ConfigDigit{
			Height:     args.Height,
			Width:      args.Width,
			MaxSkew:    0.7,
			DotCount:   80,
			CaptchaLen: args.Len,
		}

		captchaId, capD := base64Captcha.GenerateCaptcha("", configD)
		base64Image := base64Captcha.CaptchaWriteToBase64Encoding(capD)

		return &CaptchaNumberBase64Type{
			CaptchaId:   captchaId,
			Base64Image: base64Image,
		}, err

	}
}
