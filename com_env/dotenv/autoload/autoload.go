package autoload

import (
	"log"

	"github.com/zhangxiaofeng05/com/com_env/dotenv"
)

// import _ "github.com/zhangxiaofeng05/com/com_env/dotenv/autoload"
func init() {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}
