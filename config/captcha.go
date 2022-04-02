package config

import "DouyinParser/pkg/config"

func init() {
	config.Add("captcha", config.StrMap{
		// 验证码样式，具体的配置可以到 https://captcha.mojotv.cn/ 网站 request 提交请求中抓取
		// 确定好自己满意的验证码样式后，把 JSON 直接完整复制粘贴过来就可以了
		"style_json": config.Env("CAPTCHA_STYLE_JSON", `{"ShowLineOptions":["2","4","8"],"CaptchaType":"string","Id":"","VerifyValue":"duku3m","DriverAudio":{"Length":6,"Language":"zh"},"DriverString":{"Height":60,"Width":230,"ShowLineOptions":14,"NoiseCount":480,"Source":"234568wertyupkjhfdsazxcvnm","Length":6,"Fonts":["Flim-Flam.ttf","DeborahFancyDress.ttf","actionj.ttf","RitaSmith.ttf","chromohv.ttf","ApothecaryFont.ttf"],"BgColor":{"R":0,"G":0,"B":0,"A":0}},"DriverMath":{"Height":60,"Width":240,"ShowLineOptions":14,"NoiseCount":0,"Length":6,"Fonts":["wqy-microhei.ttc"],"BgColor":{"R":0,"G":0,"B":0,"A":0}},"DriverChinese":{"Height":60,"Width":320,"ShowLineOptions":0,"NoiseCount":0,"Source":"设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,,不想要,的值","Length":2,"Fonts":["wqy-microhei.ttc"],"BgColor":{"R":125,"G":125,"B":0,"A":118}},"DriverDigit":{"Height":80,"Width":240,"Length":5,"MaxSkew":0.7,"DotCount":80}}`),
	})
}
