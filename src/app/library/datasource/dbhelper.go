package datasource

// 初始化 engine ，创建单例

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql" // 使用MySQL的隐式驱动
	"github.com/go-xorm/xorm"
	"github.com/pelletier/go-toml"
	"ielpm.cn/projectweb/src/app/config"
	"ielpm.cn/projectweb/src/app/config/diserver"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

// InstanceMaster 数据库主实例
func InstanceMaster() *xorm.Engine {

	if masterEngine != nil {
		return masterEngine
	}
	// WARNNING:
	// 互斥锁，如果有多个线程同时访问，且masterEngine == nil 的时候，第一线程创建完后，后续线程依然认为 masterEngine == nil,此时需要在多一次判断
	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	container := diserver.GetDI().Container
	var tomlC *toml.Tree
	container.Invoke(func(config *config.Config) {
		tomlC = config.New()
	})

	//  读取配置文件的数据
	driver := tomlC.Get("database.dirver").(string)
	configTree := tomlC.Get(driver).(*toml.Tree)
	userName := configTree.Get("databaseUsername").(string)
	password := configTree.Get("databasePassword").(string)
	databaseName := configTree.Get("databaseName").(string)
	dbHost := configTree.Get("databaseHost").(string)
	dbPort := configTree.Get("databasePort").(string)

	connet := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", userName, password, dbHost, dbPort, databaseName)

	fmt.Println("connetion = ", connet)
	var dbDirver string
	dbDirver = tomlC.Get("database.dirver").(string)
	engine, err := xorm.NewEngine(dbDirver, connet)
	if err != nil {
		log.Fatal("dbhelper.instanceMaster error=%s", err)
	}
	masterEngine = engine
	return masterEngine
}
