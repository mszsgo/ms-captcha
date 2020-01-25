package graph

import (
	"github.com/graphql-go/graphql"
	"github.com/mojocn/base64Captcha"
)

func init() {
	QueryFields(_query)
}

var _query = graphql.Fields{
	"captcha": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "CaptchaType",
			Fields: graphql.Fields{
				"verify": &graphql.Field{
					Description: "图片验证码校验, 成功返回true，失败返回false ，此功能非前端使用，由服务端调用",
					Type:        graphql.Boolean,
					Args: graphql.FieldConfigArgument{
						"captchaId":    &graphql.ArgumentConfig{Type: NewNonNullString, DefaultValue: nil, Description: "验证码编号，获取验证码时返回"},
						"captchaValue": &graphql.ArgumentConfig{Type: NewNonNullString, DefaultValue: nil, Description: "用户输入的验证码值"},
					},
					Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
						captchaId := p.Args["captchaId"].(string)
						captchaValue := p.Args["captchaValue"].(string)
						bool := base64Captcha.VerifyCaptcha(captchaId, captchaValue)
						return bool, err
					},
				},
				"number": &graphql.Field{
					Type: graphql.NewObject(graphql.ObjectConfig{
						Name:       "CaptchaNumberBase64Type",
						Interfaces: nil,
						Fields: graphql.Fields{
							"captchaId":   &graphql.Field{Type: String, Description: "验证码编号，验证时需要使用"},
							"base64Image": &graphql.Field{Type: String, Description: "验证码图片Base64编码"},
						},
						IsTypeOf:    nil,
						Description: "验证码图片Base64编码与验证码图片编号",
					}),
					Args: graphql.FieldConfigArgument{
						"len":    &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 6, Description: "生成验证码字符长度"},
						"width":  &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 240, Description: "生成验证码像素宽度"},
						"height": &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 80, Description: "生成验证码像素高度"},
					},
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						len := p.Args["len"].(int)
						width := p.Args["width"].(int)
						height := p.Args["height"].(int)

						// 数字验证码配置
						var configD = base64Captcha.ConfigDigit{
							Height:     height,
							Width:      width,
							MaxSkew:    0.7,
							DotCount:   80,
							CaptchaLen: len,
						}

						captchaId, capD := base64Captcha.GenerateCaptcha("", configD)
						base64Image := base64Captcha.CaptchaWriteToBase64Encoding(capD)

						type NumberCaptcha struct {
							CaptchaId   string `json:"captchaId"`
							Base64Image string `json:"base64Image"`
						}
						return &NumberCaptcha{
							CaptchaId:   captchaId,
							Base64Image: base64Image,
						}, e
					},
					Description: "获取数字验证码Base64编码字符串与验证码编号",
				},
			},
			IsTypeOf:    nil,
			Description: "提供图片验证码生成与验证",
		}),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return "", e
		},
		Description: "图片验证码服务，提供图片验证码生成与验证功能",
	},
}
