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
	ID   uint   `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
	Age  int    `json:"age" gorm:"age"`
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
	e.DELETE("user/:id", u.DeleteUser)
}

// CreateUser ...
func (u *User) CreateUser(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}

	err := db.Create(&u).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u.ID)
}

// GetUser ...
func (u *User) GetUser(c echo.Context) error {
	id := c.Param("id")
	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

// DeleteUser ...
func (u *User) DeleteUser(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}
	err := db.Delete(&u).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Deleted")
}

var db *gorm.DB

func main() {
	db = InitDB()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	e := echo.New()

	InitMiddleware(e)

	u := new(User)
	InitRouting(e, u)

	e.Start(":8085")
}
