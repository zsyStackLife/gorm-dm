**一、引入依赖**

**go get github.com/zsyStackLife/gorm-dm@v1.0.5**

**二、使用示例**

```
import (
	dm "github.com/zsyStackLife/gorm-dm"
	_ "github.com/zsyStackLife/gorm-dm/dm8"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	_, _ = gorm.Open(dm.Open(dm.BuildDMUrl("username", "paasword", "host:port",
		map[string]string{"schema": "dbname"})), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error)})
}
```

