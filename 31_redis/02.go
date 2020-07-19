package main
// 链接池  redis.Pool
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
var pool *redis.Pool
func init(){
	pool = &redis.Pool{
		MaxIdle:8, // 最大空闲数
		MaxActive:0,  // 最大链接数,0为无限制
		IdleTimeout:100, // 最大空闲时间
		Dial:func()(redis.Conn, error){//初始化链接的代码，链接哪个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	// 从pool中取出一个链接
	conn:= pool.Get()
	defer conn.Close()
	defer pool.Close()
	
	_,err :=conn.Do("Set", "name","tomCat")
	if err != nil {
		fmt.Printf("err=%v", err)
	return 
	}
	r, err:=redis.String(conn.Do("Get", "name"))
	fmt.Printf("r=%v", r)

	// 从pool取出链接要保证链接池没有关闭
}