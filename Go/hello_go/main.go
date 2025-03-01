package main

func main() {
	// router.go で定義した SetupRouter() を呼び出す
	router := SetupRouter()
	// サーバーの起動
	router.Run(":8080")
}
