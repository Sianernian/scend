package main

import (
	_ "BeegoProject123/db_mysql"
	_ "BeegoProject123/routers"
	"github.com/astaxie/beego"
)

func main() {

	//config :=beego.AppConfig
	//appName :=config.String("appname")
	//port ,err:=config.Int("httpport")
	//if err !=nil{
	//	panic("项目配置文件解析失败，请检查配置文件")
	//}
	//fmt.Println("应用名称：",appName)
	//fmt.Println("端口号：",port)

	beego.Run()
}

/*                 ____________
                  |            \
                 |   _      _   \
                |       <        \
               |      ~~~~~~      \
               ————————————————————

           \`,""   ,'7"r-..__/ \
		  ,'\ `, ,',','    _/   \
		 /   \  7 / /     (   \ |
		J     \/ j  L______\  / |
		L   __JF"""/""\"\_,    /
		L,-"|❤|  | ❤ |  L_  _/
		F   \_ /  \__/   `-  j|
			.-'    `"""    ,' |          _..====.._
			\__/         r"_  A        ,' _..---.._`,
			 `-.______,,-L// / \  ___,' ,'_..:::.. `,`,
					  j   / / / 7"    `-<""=:'  '':. \ \
					 /   <,' /  F  . i , \   `,    :T W I
					 |    \,'  /    >X<  |     \   :| | L
					 |     `._/    ' ! ` |      I  :| |  G
					  \           \     /       |  :H T  |
					 __>-.__   __,-\   |        |  S P   |
					/     /|   | \  \  \_       | 'A R   |
				   /   __/ |   |  L  L   \      K./ /    L
				  /   |    |   4  I  I    |    E./ /   ,'
				 J    \    I    L F  I    |    // / _,'
		_________|     |   F    |J   F    |   //_/-'
		\   __   |    /   J     |/  J     |  /="
		\\  \_\  \__,' \  F     |   F     |
		\\\_____________\/      F__/      F
		 \\|  你是真的帅 |_____/  (______/
		  \/__别不信___/

*/
