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
		var tempVal temp 
		data, err := c.GetRawData()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
				"message" : "failed get data",
			})
			c.Abort()
			return
		}

		if data != nil {
			errors := make(map[string]string)

			err := json.Unmarshal(data, &tempVal)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error" : err.Error(),
					"message" : "failed to unmarshal",
				})
				c.Abort()
				return
			}

			mapTemp := make(map[string]interface{})
			mapTemp["name"] = tempVal.Name
			mapTemp["email"] = tempVal.Email
			mapTemp["address"] = tempVal.Address
			mapTemp["phone"] = tempVal.Phone
			mapTemp["gender"] = tempVal.Gender

			for key, value := range mapTemp {
				if reflect.TypeOf(value).Kind() != reflect.String {
					errors[key] = fmt.Sprintf("invalid type. Expected string but got %T", value)
				}
			}
			
			if len(errors) > 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": errors,
					"message" : "invalid data type",
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
