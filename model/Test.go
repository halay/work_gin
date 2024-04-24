package model

import "fmt"

func RedisString() {
	var key = "halay"
	err := RedisDb.Set(key, "hello redis", 0).Err()
	if err != nil {
		fmt.Println("set redis string err:", err)
		return
	} else {
		fmt.Println("set redis string success")
	}
	result, err := RedisDb.Get(key).Result()
	if err != nil {
		fmt.Println("get redis string err:", err)
		return
	}
	fmt.Println("get redis string success:", result)

	result2, err := RedisDb.Get("name").Result()
	fmt.Println("err2:", err)
	fmt.Println("result2:", result2)
}

type Articles struct {
	Title      string
	Content    string
	View       int
	Favourites int
}

func RedisHash() {
	articles := Articles{"halay", "hello halay content", 10, 0}
	articleKey := "halay_hash"
	err := RedisDb.HMSet(articleKey, ToStringMap(&articles)).Err()
	if err != nil {
		fmt.Println("set hash map err:", err)
		return
	}
	outMap, err := RedisDb.HGetAll(articleKey).Result()
	if err != nil {
		fmt.Println("get hash map err:", err)
		return
	}
	for index, value := range outMap {
		fmt.Printf("\n %s:%s", index, value)
	}
	view := RedisDb.HIncrBy(articleKey, "View", 1).Val()
	fmt.Println("get hash map view:", view)
	return

}
func ToStringMap(m *Articles) map[string]interface{} {
	var HashMap = make(map[string]interface{}, 0)
	HashMap["Title"] = m.Title
	HashMap["Content"] = m.Content
	HashMap["View"] = m.View
	HashMap["Favourites"] = m.Favourites
	return HashMap
}
