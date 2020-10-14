package db

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	DefaultPoolIdle = 20              // 默认空闲连接数量
	DefaultLifeTime = time.Minute * 5 // 默认连接超时时间
)

var pg = Postgresql{
	Host:     "localhost",
	Port:     "5432",
	User:     "postgres",
	Password: "123456",
	Name:     "lai",
}

type Postgresql struct {
	Host     string // 192.168.3.235
	Port     string // 5432
	User     string // postgres
	Password string // 123456
	Name     string // lai
}

var (
	sqlArray = map[string]string{}
	dbArray  = [...]string{"postgres"}
	dc       = dataCenter{}
)

func Init() {
	pgConnection := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", pg.Host, pg.User, pg.Name, pg.Password)
	sqlArray["postgres"] = pgConnection

	for k, _ := range dbArray {
		if dc.pool == nil {
			dc.createDataCenter()
		}
		dc.pool[dbArray[k]] = CreatePool(dbArray[k]) // 将连接池放入数据中心
	}
	log.Println("db init finish.")
}

func init() {
	pgConnection := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", pg.Host, pg.User, pg.Name, pg.Password)
	sqlArray["postgres"] = pgConnection

	for k, _ := range dbArray {
		if dc.pool == nil {
			dc.createDataCenter()
		}
		dc.pool[dbArray[k]] = CreatePool(dbArray[k]) // 将连接池放入数据中心
	}
	log.Println("db init finish.")
}

// 数据中心
type dataCenter struct {
	pool map[string]*pl
}

type pl struct {
	name         string        // 数据库类别名称 postgres mysql
	MaxlifeTime  time.Duration // 空闲连接的时间(连接复用的时长)
	MaxIdleConns int           // 空闲连接数据
	sql          string        // 数据库连接语句
	Con          *gorm.DB      // 连接对象
}

// 获取一个数据库连接
func Pop(dbName string) *gorm.DB {
	dbName = strings.ToLower(dbName)
	if checkDBName(dbName) {
		if _, ok := dc.pool[dbName]; ok {
			return dc.pool[dbName].getConnection()
		}
	}
	return nil
}

// 检查数据库是否支持
func checkDBName(dbName string) (ok bool) {
	for i, _ := range dbArray {
		if strings.EqualFold(dbName, dbArray[i]) {
			ok = true
			return
		}
	}
	return
}

// 创建数据中心
func (dc *dataCenter) createDataCenter() {
	dc.pool = make(map[string]*pl)
}

// 创建连接池
func CreatePool(name string) (p *pl) {
	p = &pl{
		name:         name,
		MaxlifeTime:  DefaultLifeTime,
		MaxIdleConns: DefaultPoolIdle,
		sql:          getSql(name),
	}
	p.build()

	return
}

// 初始化连接池
func (p *pl) build() {
	db, err := p.createConnection()
	if err != nil {
		log.Println(p.sql)
		log.Panic("failed to create database connection, Error =", err)
	}
	p.Con = db

	return
}

// 获取一个数据库连接语句
func getSql(dbType string) string {
	return sqlArray[dbType]
}

// 创建数据库连接
func (p *pl) createConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open(p.name, p.sql)
	if err != nil {
		return
	}

	// 数据库连接的存活时间设置应比连接池连接的存活时间长，由连接池来维护连接
	db.DB().SetConnMaxLifetime(p.MaxlifeTime)
	db.DB().SetMaxIdleConns(p.MaxIdleConns)

	return
}

func (p *pl) getConnection() *gorm.DB {
	return p.Con
}
