package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	InitDB()

	e := echo.New()

	// Aktifkan middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://202.43.249.105:3000/"},                                      // Allow frontend
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}, // Allowed HTTP methods
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},                   // Allowed headers
	}))

	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, API!"})
	})

	e.GET("/students", GetStudents)

	e.Logger.Fatal(e.Start(":1323"))
}

func GetStudents(c echo.Context) error {
	rows, err := db.Query("SELECT id, name, email, age FROM student")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Email, &s.Age); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		students = append(students, s)
	}

	return c.JSON(http.StatusOK, students)
}
