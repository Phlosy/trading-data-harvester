package connect

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func Connect() (*sql.DB, error) {
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", ClickhouseHost, ClickhousePort)}, // 数据库地址和端口
		Auth: clickhouse.Auth{
			Database: ClickhouseDatabase, // 数据库名称
			Username: ClickhouseUsername, // 用户名
			Password: ClickhousePassword, // 密码
		},
		TLS: nil,
		Settings: clickhouse.Settings{
			"max_execution_time": 60, // 最大执行时间（秒）
		},
		DialTimeout: time.Second * 30, // 连接超时时间
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4, // 压缩方法
		},
		Debug:                true,  // 是否开启调试模式
		BlockBufferSize:      10,    // 块缓冲区大小
		MaxCompressionBuffer: 10240, // 最大压缩缓冲区大小
		ClientInfo: clickhouse.ClientInfo{ // 客户端信息（可选）
			Products: []struct {
				Name    string // 产品名称
				Version string // 产品版本
			}{
				{Name: "trading-data-harvester", Version: "0.1"},
			},
		},
	})
	conn.SetMaxIdleConns(5)            // 最大空闲连接数
	conn.SetMaxOpenConns(10)           // 最大打开连接数
	conn.SetConnMaxLifetime(time.Hour) // 连接的最大生命周期

	return conn, nil
}
