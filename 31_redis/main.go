package main

func main() {
	
}
/*
	1. 添加 key-value [set]
	2. 查看当前redis的所有key  [keys *]
	3. 获取key对应的值 [get key]
	4. 切换redis数据库 [select index]
	5. 如何查看当前数据库key-value数量 [dbsize]
	6. 清空当前数据库的key-val和清空所有数据库的key-val [flushdb/flushall]
	
	字符串操作 set/get/del
	超时废除 setex key second(单位:s) value
	同时设置多个key-value  mset key value key value ...
						  mget key key key ...

	Hash 类似go中的map	hset key field value      exmple:hset user1 name "smith"
						hgetall key  // 获取key中的所有key-value
						hmset key name jerry age 110 job "java coder"  // 一次性全部设置
						hmget key name age job        // 一次性获取全部
						hexists key name  // key 中有没有name字段

	List 链表 有序 可重复	lpush key value value value ... // 左插 rpush 右插
						lrange key 0 -1 //0到-1
						lpop/rpop   删除单个元素
						del 删除整个list
						LLen key // 返回key长度

	Set	 集合 无序 不可重复 	sadd emails ton@sohu.com jack@qq.com  // 存
						  	smembers emails       // 取所有值
						  	sismember             // 判断值是否存在
						  	srem                  // 删除指定的值


*/