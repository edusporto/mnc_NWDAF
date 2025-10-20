module nwdaf.com

go 1.14

replace nwdaf.com/logger => ../logger

replace nwdaf.com/service => ../service

replace nwdaf.com/factory => ../factory

replace nwdaf.com/util => ../util

replace nwdaf.com/consumer => ../consumer

replace nwdaf.com/context => ../context

replace nwdaf.com/mtlf => ../mtlf

replace nwdaf.com/anlf => ../AnLF

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.1 // indirect
	github.com/free5gc/http2_util v1.0.0
	github.com/free5gc/logger_util v1.0.0
	github.com/free5gc/openapi v1.0.0
	github.com/free5gc/path_util v1.0.0
	github.com/free5gc/util v1.0.4
	github.com/free5gc/version v1.0.0
	github.com/gin-gonic/gin v1.7.3
	github.com/google/uuid v1.3.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli v1.22.4
	gopkg.in/yaml.v2 v2.4.0
)
