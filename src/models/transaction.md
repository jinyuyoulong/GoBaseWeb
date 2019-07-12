数据库事务处理 示例

package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Blance  float64
	Version int `xorm:"version"`
}

var engine *xorm.Engine

func init() {
	// 根据名称注册驱动并创建 ORM 引擎
	var err error
	engine, err = xorm.NewEngine("mysql", "root:pwd@tcp(127.0.0.1:3306)/dbname?charset=utf8")
	if err != nil {
		log.Fatalf("fail to create engine：%v", err)
	}
	if err := engine.Sync(new(Account)); err != nil {
		log.Fatalf("fail to sync database：%v", err)
	}
	// 记录SQL语句log
	f, err := os.Create("sql.log")
	if err != nil {
		log.Fatalf("fail to create log file %v", err)
	}
	// defer f.Close() 打开这行就记录不到，需要注释掉
	engine.SetLogger(xorm.NewSimpleLogger(f))
	engine.ShowSQL(true)
}

func transfer(id1, id2 int, blance float64) error {
	account1 := &Account{}
	has1, err := engine.Id(id1).Get(account1)
	if err != nil {
		return err
	} else if !has1 {
		return errors.New("account1 not found")
	}
	account2 := &Account{}
	has2, err := engine.Id(id2).Get(account2)
	if err != nil {
		return err
	} else if !has2 {
		return errors.New("account2 not found")
	}
	if account1.Blance < blance {
		return errors.New("blance not enough")
	}
	account1.Blance -= blance
	account2.Blance += blance

	// 事务处理
	sess := engine.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return errors.New("fail to session begin")
	}
	if _, err := sess.Id(id1).Cols("blance").Update(account1); err != nil {
		sess.Rollback()
		return errors.New("fail to update 1")
	}
	if _, err := sess.Id(id2).Cols("blance").Update(account2); err != nil {
		sess.Rollback()
		return errors.New("fail to update 2")
	}
	return sess.Commit()
}

func main() {
	// 转账
	err := transfer(5, 7, 100)
	if err != nil {
		log.Fatalf("fail to transfer %v", err)
	}
	fmt.Println("transfer OK")
}

