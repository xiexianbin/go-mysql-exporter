/*
Copyright (c) 2022 xiexianbin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/xiexianbin/go-mysql-exporter/collectors"
)

var (
	h         bool
	addr      string
	mysqlAddr string
	user      string
	password  string
	db        string
	dsn       string
	version   bool
)

func init() {
	flag.BoolVar(&h, "help", false, "show help message")
	flag.StringVar(&addr, "addr", "0.0.0.0:9306", "server url")
	flag.StringVar(&mysqlAddr, "mysqlAddr", "127.0.0.1:3306", "mysql url")
	flag.StringVar(&user, "u", "root", "mysql user")
	flag.StringVar(&password, "p", "root", "mysql password")
	flag.StringVar(&db, "d", "mysql", "mysql db")
	flag.BoolVar(&version, "v", false, "version info")
	flag.Usage = func() {
		fmt.Println(`Usage: mysql-exporter -addr 0.0.0.0:9306 -mysqlAddr 127.0.0.1:3306 -u root -p root -d mysql`)
		flag.PrintDefaults()
	}
	flag.Parse()

	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=PRC&parseTime=True", user, password, mysqlAddr, db)
}

func main() {
	if h == true {
		flag.Usage()
		return
	}
	if version == true {
		fmt.Println("v0.1.0")
		return
	}

	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// 注册指标
	//mysqlConnectInfo := prometheus.NewGauge(prometheus.GaugeOpts{
	//	Name: "mysql_connect_info",
	//	Help: "mysql connect info",
	//	ConstLabels: prometheus.Labels{
	//		"addr": mysqlAddr,
	//	},
	//})
	//prometheus.MustRegister(mysqlConnectInfo)
	//mysqlConnectInfo.Set(1)
	prometheus.MustRegister(prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "mysql_connect_info",
		Help: "mysql connect info",
		ConstLabels: prometheus.Labels{
			"addr": mysqlAddr,
		},
	}, func() float64 {
		return 1
	}))

	prometheus.MustRegister(collectors.NewPingCollector(db))
	prometheus.MustRegister(collectors.NewVersionController(db))
	prometheus.MustRegister(collectors.NewThreadsCollector(db))
	prometheus.MustRegister(collectors.NewComCollector(db))

	// 注册控制器
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("/metrics"))
	})

	// 启动web服务
	log.Println("license on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
