module github.com/deatil/lakego-doak

go 1.16

replace github.com/deatil/go-filesystem => ./../go-filesystem

require (
	github.com/casbin/casbin/v2 v2.37.4
	github.com/deatil/go-filesystem v0.0.3
	github.com/fsnotify/fsnotify v1.5.1
	github.com/gin-gonic/gin v1.7.4
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.9.0
	github.com/go-redis/cache/v8 v8.4.3
	github.com/go-redis/redis/v8 v8.11.4
	github.com/golang-jwt/jwt/v4 v4.1.0
	github.com/google/uuid v1.3.0
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/mojocn/base64Captcha v1.3.5
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.16
)

require (
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/fatih/color v1.13.0
	github.com/flosch/pongo2/v4 v4.0.2
	github.com/go-redis/redis/extra/redisotel/v8 v8.11.4
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/tebeka/strftime v0.1.5 // indirect
	go.uber.org/dig v1.13.0
	golang.org/x/sys v0.0.0-20211020064051-0ec99a608a1b // indirect
)