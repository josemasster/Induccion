package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/beego/beego/logs"
)

func GetReqNew(endpoint string, route string, target interface{}) error {
	url := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String(endpoint) + route
	var response map[string]interface{}
	var err error
	err = GetJson(url, &response)
	err = ExtracData(response, &target)
	return err
}

/*func Post(endpoint string, route string, target interface{}) error {
	url := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String(endpoint) + route
	var response map[string]interface{}
	var err error
	err = PostJson(url, &response)
	err = ExtracData(response, &target)
	return err
}*/

func GetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()
	return json.NewDecoder(r.Body).Decode(target)
}

/*
	func PostJson(url string, target interface{}) error {
		r, err := http.Post(url)
		if err != nil {
			return err
		}
		defer func() {
			if err := r.Body.Close(); err != nil {
				beego.Error(err)
			}
		}()
		return json.NewDecoder(r.Body).Decode(target)
	}
*/
func ExtracData(response map[string]interface{}, v interface{}) error {
	var err error
	if response["Success"] == false {
		err = errors.New(fmt.Sprint(response["Data"], response["Message"]))
		panic(err)
	}
	datatype := fmt.Sprintf("%v", response["Data"])
	switch datatype {
	case "map[]", "[map[]]":
		break
	default:
		err = FillStrcut(response["Data"], &v)
		response = nil
	}
	return err
}

func FillStrcut(in interface{}, out interface{}) (err error) {
	var str []byte
	if str, err = json.Marshal(in); err != nil {
		return
	}
	err = json.Unmarshal(str, &out)
	return
}

func ErrorController(c beego.Controller, controller string) {
	if err := recover(); err != nil {
		logs.Error(err)
		localError := err.(map[string]interface{})
		c.Data["message"] = (beego.AppConfig.String("appname") + "/" + controller + "/" + (localError["funcion"]).(string))
		c.Data["data"] = (localError["err"])
		if status, ok := localError["status"]; ok {
			c.Abort(status.(string))
		} else {
			c.Abort("500")
		}
	}
}
