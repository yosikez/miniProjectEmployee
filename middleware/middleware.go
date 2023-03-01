package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func JsonValid() gin.HandlerFunc {
	return func(c *gin.Context) {
		type temp struct {
			Name      interface{}    
			Email     interface{}    
			Address   interface{}
			Phone     interface{}
			Gender    interface{}
		}
		var ne temp 
		data, err := c.GetRawData()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
			c.Abort()
			return
		}

		if data != nil {
			errrs := make(map[string]string)

			err := json.Unmarshal(data, &ne)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error" : err.Error(),
					"message" : "failed to unmarshal",
				})
				c.Abort()
				return
			}

			mapTemp := make(map[string]interface{})
			mapTemp["name"] = ne.Name
			mapTemp["email"] = ne.Email
			mapTemp["address"] = ne.Address
			mapTemp["phone"] = ne.Phone
			mapTemp["gender"] = ne.Gender

			for key, value := range mapTemp {
				if reflect.TypeOf(value).Kind() != reflect.String {
					errrs[key] = fmt.Sprintf("Invalid type. Expected string but got %T", value)
				}
			}
			
			if len(errrs) > 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": errrs,
				})
				c.Abort()
				return
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
				c.Next()
			}
		}
		c.Next()
	}
}
