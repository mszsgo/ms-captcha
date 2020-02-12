package schema

import (
	"github.com/graphql-go/graphql"
)

type Query struct {
	Captcha *CaptchaQuery `description:"验证码"`
}

type CaptchaQuery struct {
	Number *CaptchaNumberBase64Type `description:"数字验证码，返回base64编码的图片，可使用img标签直接显示"`

	Verify *VerifyType `description:"验证码验证，限服务端使用"`
}

func (*CaptchaQuery) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		return "", err
	}
}
