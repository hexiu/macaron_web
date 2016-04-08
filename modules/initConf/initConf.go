package initConf

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

func InitConf() {
	conf, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf)
}
