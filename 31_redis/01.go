package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	c,err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn redis failed",err)
	return 
	}
	defer c.Close()
	_, err = c.Do("Set", "key1",998)
	if err != nil{
		fmt.Printf("err=%v", err)
		return 
	}
	r, err := redis.Int(c.Do("Get", "key1"))
	if err != nil {
		fmt.Println("get key1 failed,",err)
	return 
	}
	fmt.Printf("r=%v",r)
}