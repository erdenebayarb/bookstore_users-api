package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/erdenebayarb/bookstore_users-api/domain/users"
	"github.com/erdenebayarb/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(bytes))
	if err != nil {
		//TODO: Handle error
		return
	}

	if err := json.Unmarshal([]byte(bytes), &user); err != nil {
		switch err := err.(type) {
		case *json.SyntaxError:
			fmt.Printf("syntax error at offset %d: %v\n", err.Offset, err)
		case *json.UnmarshalTypeError:
			fmt.Printf("type error at offset %d: %v\n", err.Offset, err)
		case nil:
			// no error
		default:
			fmt.Printf("error %v\n", err)
		}
		//TODO: Handle json error
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: Handle user creation error
		return
	}

	fmt.Println(string(bytes), "user dta")
	fmt.Println(err)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")

}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")
// }
