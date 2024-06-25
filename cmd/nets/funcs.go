package nets

import (
	"os"
	"strings"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/nets/sNet"
)

func Prueba() {
	os.Setenv("SERV_Own", "ApiManagement")

	ServiceOwn, err := sNet.NewService("SERV_Own", "")
	if err != nil {
		sLog.Error("sNet.NewService(): %s", err)
		return
	}
	ServiceOwn.Log()

	request := sNet.NewRequest().MethodPost().SetEndPoint("a")

	body := "aqqbPIycNs+ToAXQIX3TInl/ArkJcqbnhmZQpneqxL6iHg+nynYVzxusUuv5Z7qfRkSVmiIXeFSTRL6FcwDiyUJBMU1r3NxarRFjvMTlTGIsbpDwwjNyqYtMfUlkDX/iEmzpnZvxri0fJ2sipA8ouQ=="
	request.SetBody(strings.NewReader(body))

	data, err := request.Call(nil, ServiceOwn, nil)
	if err != nil {
		sLog.Error("request.Call(): %s", err)
		return
	}

	sLog.Info("data: %s", data)
}
