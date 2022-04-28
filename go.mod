module chat_rooms

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/pkg/errors v0.8.1
	github.com/spf13/viper v1.10.1
	go.uber.org/zap v1.17.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.3
	gorm.io/plugin/soft_delete v1.1.0
)
