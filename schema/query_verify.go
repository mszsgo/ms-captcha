package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/mojocn/base64Captcha"
	"github.com/mszsgo/hjson"
)

type VerifyType int64

func (*VerifyType) Description() string {
	return "验证码验证，返回1代表成功，0代表失败"
}

type VerifyTypeArgs struct {
	CaptchaId    string `graphql:"!" description:"验证码编号"`
	CaptchaValue string `graphql:"!" description:"验证码值"`
}

func (*VerifyType) Args() *VerifyTypeArgs {
	return &VerifyTypeArgs{}
}

func (*VerifyType) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *VerifyTypeArgs
		hjson.MapToStruct(p.Args, &args)
		bool := base64Captcha.VerifyCaptcha(args.CaptchaId, args.CaptchaValue)
		if bool {
			i = 1
		} else {
			i = 0
		}
		return i, err
	}
}
