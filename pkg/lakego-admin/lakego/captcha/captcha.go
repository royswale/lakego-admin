package captcha

import (
    "time"
    "image/color"

    "github.com/mojocn/base64Captcha"

    "github.com/deatil/lakego-admin/lakego/redis"
)

// id, b64s, err := New.Generate()
func New(config Config, redis redis.Redis) Captcha {
    ds := base64Captcha.NewDriverString(
        config.Height,
        config.Width,
        config.NoiseCount,
        config.ShowLineOptions,
        config.Length,
        config.Source,
        &color.RGBA{
            R: config.RBGA.R,
            G: config.RBGA.G,
            B: config.RBGA.B,
            A: config.RBGA.A,
        },
        []string{
            config.Fonts,
        },
    )

    driver := ds.ConvertFonts()
    store := CaptchaStore{
        redis:  &redis,
        key:    config.Key,
        config:	config,
    }

    return Captcha{
        Captcha: base64Captcha.NewCaptcha(driver, store),
    }
}

// 颜色
type RBGA struct {
    R uint8
    B uint8
    G uint8
    A uint8
}

// 配置
type Config struct {
    Key string
    ExpireTimes int

    Height int
    Width int
    NoiseCount int
    ShowLineOptions int
    Length int
    Source string
    Fonts string

    RBGA RBGA
}

/**
 * 验证码
 *
 * @create 2021-9-15
 * @author deatil
 */
type Captcha struct {
    *base64Captcha.Captcha
}

type CaptchaStore struct {
    key    string
    redis  *redis.Redis
    config Config
}

func (a CaptchaStore) getKey(v string) string {
    return a.key + ":" + v
}

func (a CaptchaStore) Set(id string, value string) {
    t := time.Second * time.Duration(a.config.ExpireTimes)
    a.redis.Set(a.getKey(id), value, int(t))
}

func (a CaptchaStore) Get(id string, clear bool) string {
    var (
        key = a.getKey(id)
        val string
    )

    err := a.redis.Get(key, &val)
    if err != nil {
        return ""
    }

    if !clear {
        a.redis.Delete(key)
    }

    return val
}

func (a CaptchaStore) Verify(id, answer string, clear bool) bool {
    v := a.Get(id, clear)
    return v == answer
}