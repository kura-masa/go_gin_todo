package main

import (
	"go_gin_todo/config"
	"go_gin_todo/models"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	//ログ設定
	f, _ := os.Create(config.Config.LogFile)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	r := gin.Default()
	// htmlファイル（テンプレート）を読み込んでる（レンダリング）
	r.LoadHTMLGlob("templates/**/*")

	// todo create
	r.POST("/todos/save", func(c *gin.Context) {
		// contentの内容をdbに入れる。
		models.CreateTodo(c.PostForm("content"))
		// このパスにリダイレクトする
		c.Redirect(http.StatusMovedPermanently, "/todos/list")
	})

	// todo create
	r.POST("/todos/update", func(c *gin.Context) {
		//strconv.Atoiで文字から数字に
		//PostFormでフォームの値を持ってこれる
		id, _ := strconv.Atoi(c.PostForm("id"))
		content := c.PostForm("content")
		todo, _ := models.GetTodo(id)
		todo.Content = content
		models.UpdateTodo(todo)

		c.Redirect(http.StatusMovedPermanently, "/todos/list")
	})

	//todo edit 更新画面を表示する
	r.GET("/todos/edit", func(c *gin.Context) {
		// Queryは、URLから情報を取ってくる
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			log.Fatalln(err)
		}
		todo, _ := models.GetTodo(id)

		// HTMLメソッドは、tmplファイルをhtmlファイルに変更して表示する
		// その際、情報を追加できる
		c.HTML(http.StatusOK, "edit.tmpl", gin.H{ //ここ　元々edit.tmplだった
			"title": "Todo",
			"todo":  todo,
		})
	})

	//todo delete リストを削除するもの
	r.GET("/todos/destroy", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			log.Fatalln(err)
		}
		models.DeleteTodo(id)

		c.Redirect(http.StatusMovedPermanently, "/todos/list")
	})

	//todo list リストを表示
	//GETの第一引数の"/todos/list"はパス
	//ユーザがこのパスにアクセスしたら、（そのURLに飛ぶ）
	//以下のハンドラー関数が実行される。
	//https://twitter.com/kuramasa19は. 左からプロトコル・ドメイン・パス
	r.GET("/todos/list",
		//ハンドラー関数
		func(c *gin.Context) {
			var todos []models.Todo

			//todosスライスの先頭のアドレス?とtodosの情報（スラライスと型）を引数として渡して、
			//それと一致する情報を見つけて、todosに格納している？
			models.Db.Find(&todos)

			c.HTML(http.StatusOK, "list.tmpl", gin.H{
				"title": "Todo",
				"todos": todos,
			})
		})

	return r
}

func main() {

	r := setupRouter()
	r.Run(":8080")
	//サーバーを立ち上げる
	//ずっとユーザからのリクエストがあったら対応する。

	//ユーザーからのリクエストとは、（ボタンを押すなどして）特定のURLにアクセスすること。
	//そのURL(パス)が、上のパスと一致したら、そこのハンドラー関数を実行する

	//URLは、ユーザーからのリクエスト
}
