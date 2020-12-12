package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User ...
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// InitDB ...
func InitDB() *gorm.DB {
	dsn := "root:root@tcp(db:3306)/originalcode?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// InitMiddleware ...
func InitMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

// InitRouting ...
func InitRouting(e *echo.Echo, u *User) {
	e.POST("user", u.CreateUser)
	e.GET("user/:id", u.GetUser)
	e.GET("users", u.GetUsers)
	e.PUT("user/:id", u.UpdateUser)
	e.DELETE("user/:id", u.DeleteUser)
}

// CreateUser ...
func (u *User) CreateUser(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.String(http.StatusOK, "You are "+"Name:"+u.Name)
}

// GetUser ...
func (u *User) GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "You are "+id)
}

// GetUsers ...
func (u *User) GetUsers(c echo.Context) error {
	email := c.QueryParam("email")
	return c.String(http.StatusOK, "You got All Users."+email)
}

// UpdateUser ...
func (u *User) UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Updated")
}

// DeleteUser ...
func (u *User) DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "Deleted")
}

func main() {
	db := InitDB()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	// defer func() {
	// 	if db != nil {
	// 		if err := db.Close(); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}
	// }()

	e := echo.New()

	InitMiddleware(e)

	u := new(User)
	InitRouting(e, u)

	e.Start(":8085")
}
