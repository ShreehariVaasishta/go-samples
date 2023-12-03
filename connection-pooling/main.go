package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const MAX_CONNECTIONS = 5

type ConnectionPool struct {
	Connections chan *sql.DB
}

func NewConnectionPool(size int, dsn string) (*ConnectionPool, error) {
	pool := &ConnectionPool{
		Connections: make(chan *sql.DB, size),
	}

	for i := 0; i < size; i++ {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, err
		}
		if err = db.Ping(); err != nil {
			return nil, err
		}

		pool.Connections <- db
	}

	return pool, nil
}

func (p *ConnectionPool) GetConnection() *sql.DB {
	return <-p.Connections
}

func (p *ConnectionPool) ReleaseConnection(conn *sql.DB) {
	p.Connections <- conn
}

func withConnPool() {
	println("!!!Running with connection pooling.")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "127.0.0.1", 3308, "afw_live_1")
	pool, err := NewConnectionPool(MAX_CONNECTIONS, dsn) // Adjust pool size as needed
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
	}

	r := gin.Default()

	r.GET("/query", func(c *gin.Context) {
		db := pool.GetConnection()
		defer pool.ReleaseConnection(db)
		time.Sleep(1 * time.Minute)

		// Perform a sample database operation
		var username string
		err := db.QueryRow("SELECT username FROM users LIMIT 1").Scan(&username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database query failed",
				"msg":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":      "Query successful",
			"query_result": username,
		})
	})

	r.Run(":9082")
}

func withoutConnPool() {
	println("!!!Running without connection pooling.")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "127.0.0.1", 3308, "afw_live_1")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/query", func(c *gin.Context) {
		time.Sleep(1 * time.Minute)
		var username string
		err := db.QueryRow("SELECT username FROM users LIMIT 1").Scan(&username)
		if err != nil {
			println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database query failed",
				"msg":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":      "Query successful",
			"query_result": username,
		})
	})

	r.Run(":9082")
}

func main() {
	// withConnPool()
	withoutConnPool()
}
