package captcha

import (
	"DouyinParser/pkg/e"
	"bytes"
	"encoding/json"
	"github.com/mojocn/base64Captcha"
	"sync"
)

// Captcha 验证码
type Captcha struct {
	Id            string
	CaptchaType   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

var instance *Captcha
var once sync.Once

// GetInstance 获取一个实例.
func GetInstance() *Captcha {
	once.Do(func() {
		instance = &Captcha{}
	})
	return instance
}

// SetConfig 载入验证码配置.
func (c *Captcha) SetConfig(cfg string) {
	// 载入配置函数
	var loadCfg = func(cfg string, captcha *Captcha) error {
		decoder := json.NewDecoder(bytes.NewBufferString(cfg))
		err := decoder.Decode(c)
		return err
	}

	if cfg == "" || e.HasError(loadCfg(cfg, c)) {
		// 如果没有配置 或报错，则使用默认的配置
		defaultCfg := `{
			"CaptchaType": "string",
			"DriverAudio": {"Length": 6, "Language": "zh"},
			"DriverString": {"Height": 60, "Width": 230, "ShowLineOptions": 14, "NoiseCount": 480, "Source": "234568wertyupkjhfdsazxcvnm", "Length": 6, "Fonts": ["Flim-Flam.ttf", "DeborahFancyDress.ttf", "actionj.ttf", "RitaSmith.ttf", "chromohv.ttf", "ApothecaryFont.ttf"], "BgColor": {"R": 0, "G": 0, "B": 0, "A": 0}},
			"DriverMath": {"Height": 60, "Width": 240, "ShowLineOptions": 14, "NoiseCount": 0, "Length": 6, "Fonts": ["wqy-microhei.ttc"], "BgColor": {"R": 0, "G": 0, "B": 0, "A": 0}},
			"DriverChinese": {"Height": 60, "Width": 320, "ShowLineOptions": 0, "NoiseCount": 0, "Source": "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,,不想要,的值", "Length": 2, "Fonts": ["wqy-microhei.ttc"], "BgColor": {"R": 125, "G": 125, "B": 0, "A": 118}},
			"DriverDigit": {"Height": 80, "Width": 240, "Length": 5, "MaxSkew": 0.7, "DotCount": 80}
		}`
		_ = loadCfg(defaultCfg, c)
	}
}

// Generate 创建一个验证码..
func (c *Captcha) Generate() (string, string, error) {
	var driver base64Captcha.Driver
	switch c.CaptchaType {
	case "audio":
		driver = c.DriverAudio
	case "string":
		driver = c.DriverString.ConvertFonts()
	case "math":
		driver = c.DriverMath.ConvertFonts()
	case "chinese":
		driver = c.DriverChinese.ConvertFonts()
	default:
		driver = c.DriverDigit
	}
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	return id, b64s, err
}

// Verify 验证码校验，clear 参数为 true 时，则验证后失效.
func (c *Captcha) Verify(id string, verifyValue string, clear bool) bool {
	return store.Verify(id, verifyValue, clear)
}

// Get 通过验证码 id 获取验证码，clear 参数为 true 时，则获取后失效.
func (c *Captcha) Get(id string, clear bool) string {
	return store.Get(id, clear)
}
