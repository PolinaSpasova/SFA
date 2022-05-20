package main

import (
	"database/sql"
	"encoding/json"
	"final/cmd"
	"fmt"
	"strconv"

	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
)

type List struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Task struct {
	Id        int64  `json:"id,omitempty"`
	Text      string `json:"text,omitempty"`
	ListId    int64  `json:"listId,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

type Response struct {
	Weather []struct {
		Description string `json:"description"`
	}
	Main struct {
		Temp float64 `json:"Temp"`
	}
	City string `json:"name"`
}

type Weather struct {
	FormatedTemp string
	Description  string
	City         string
}

func main() {

	os.Remove("data.db")
	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`Create table tasks (
                taskId INTEGER PRIMARY KEY AUTOINCREMENT,
                text   VARCHAR(32),
                listId INTEGER,
                completed INTEGER
            )`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`Create table lists (
        listId INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(32)
        )`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`Create table users (
        userId INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(32),
		password VarChar(32)
        )`)
	if err != nil {
		log.Fatal(err)
	}

	router := echo.New()
	//	g := router.Group("/localhost:3000")

	router.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		// if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
		// 	subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
		// 	return true, nil
		//}
		if username == "polina" && password == "030711" {
			return true, nil
		}

		return false, nil
	}))

	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		// This is a sample demonstration of how to attach middlewares in Echo
		return func(ctx echo.Context) error {
			log.Println("Echo middleware was called")
			return next(ctx)
		}
	})

	// Add your handler (API endpoint) registrations here
	router.GET("/api", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello, World!")
	})

	router.GET("/api/lists", func(ctx echo.Context) error {
		rows, err := db.Query("select * from lists")
		if err != nil {
			log.Fatal(err)
		}
		list := []List{}
		for rows.Next() {
			var l List
			rows.Scan(&l.Id, &l.Name)
			list = append(list, l)
		}

		return ctx.JSON(200, list)
	})

	router.POST("/api/lists", func(ctx echo.Context) error {
		l := new(List)
		l.Name = ctx.FormValue("name:")
		if err := ctx.Bind(l); err != nil {
			return err
		}
		res, err := db.Exec("insert into lists(name) values ($1)", l.Name)
		if err != nil {
			log.Fatal(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		l.Id = id
		return ctx.JSON(200, l)
	})

	router.DELETE("/api/lists/:id", func(ctx echo.Context) error {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err)
		}
		_, err = db.Exec("delete from tasks where listId=$1 ", id)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec("delete from lists where listId=$1 ", id)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.String(200, "successful operation")
	})

	router.DELETE("/api/tasks/:id", func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err)
		}
		_, err = db.Exec("delete from tasks where taskId=$1 ", id)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.String(200, "Succesfully deleted")
	})

	router.PATCH("/api/tasks/:id", func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err)
		}
		t := new(Task)
		t.Completed, _ = strconv.ParseBool(ctx.FormValue("completed:"))
		if err := ctx.Bind(t); err != nil {
			return err
		}

		_, err = db.Exec("update tasks set completed = $1 where taskId = $2", t.Completed, id)
		if err != nil {
			log.Fatal(err)
		}

		rows, err := db.Query("select * from tasks where taskId=$1", id)
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			rows.Scan(&t.Id, &t.Text, &t.ListId, &t.Completed)
		}

		return ctx.JSON(200, t)
	})

	router.POST("/api/lists/:id/tasks", func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err)
		}
		t := new(Task)
		t.Text = ctx.FormValue("text:")
		if err := ctx.Bind(t); err != nil {
			return err
		}
		t.ListId = int64(id)
		t.Completed = false
		res, err := db.Exec("insert into tasks(text,listId,completed) values ($1,$2,$3)", t.Text, id, t.Completed)
		if err != nil {
			log.Fatal(err)
		}
		t.Id, err = res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(200, t)
	})

	router.GET("/api/lists/:id/tasks", func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err)
		}

		rows, err := db.Query("select * from tasks where listId=$1", id)
		if err != nil {
			log.Fatal(err)
		}
		tasks := []Task{}
		for rows.Next() {
			var t Task
			rows.Scan(&t.Id, &t.Text, &t.ListId, &t.Completed)
			tasks = append(tasks, t)
		}

		return ctx.JSON(200, tasks)
	})

	router.GET("/api/weather", func(ctx echo.Context) error {
		var lat, lon int
		for key, v := range ctx.Request().Header {
			if key == "Lat" {
				lat, err = strconv.Atoi(v[0])
				if err != nil {
					log.Fatal(err)
				}
			}
			if key == "Lon" {
				lon, err = strconv.Atoi(v[0])
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		u := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%d&lon=%d&appid=6ffc6e72ec8fcee8ecde61518fc3ddf9", lat, lon)
		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			log.Fatal(err)
		}
		httpRes, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer httpRes.Body.Close()

		var res Response
		json.NewDecoder(httpRes.Body).Decode(&res)
		w := new(Weather)
		w.FormatedTemp = fmt.Sprintf("%.2f C", res.Main.Temp-273.15)
		w.Description = res.Weather[0].Description
		w.City = res.City
		fmt.Println(w)
		return ctx.JSON(http.StatusOK, w)
	})

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
