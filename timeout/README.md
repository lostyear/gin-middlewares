## timeout middleware
gin handler timeout middleware

### useage

```golang
import "github.com/lostyear/gin-middlewares/timeout"
```

```golang
r := gin.New()
r.Use(timeout.TimeoutMiddleware(time.Second * 2, "timeout message"))
```
