package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func JsonValidEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		type temp struct {
			Name    interface{}
			Email   interface{}
			Address interface{}
			Phone   interface{}
			Gender  interface{}
		}
		var tempVal temp
		data, err := c.GetRawData()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "failed get data",
			})
			c.Abort()
			return
		}

		if data != nil {
			errors := make(map[string]string)

			err := json.Unmarshal(data, &tempVal)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   err.Error(),
					"message": "failed to unmarshal",
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
					"error":   errors,
					"message": "invalid data type",
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

func JsonValidOpportunity() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Resource struct {
			Qty             interface{} `json:"qty" structs:"qty"`
			Position        interface{} `json:"position" structs:"position"`
			Level           interface{} `json:"level" structs:"level"`
			Ctc             interface{} `json:"ctc" structs:"ctc"`
			ProjectDuration interface{} `json:"project_duration" structs:"project_duration"`
		}

		type JsonDataOpportunity struct {
			Code            interface{} `json:"code" structs:"code"`
			ClientCode      interface{} `json:"client_code" structs:"client_code"`
			PicEmail        interface{} `json:"pic_email" structs:"pic_email"`
			OpportunityName interface{} `json:"opportunity_name" structs:"opportunity_name"`
			Description     interface{} `json:"description" structs:"description"`
			SalesEmail      interface{} `json:"sales_email" structs:"sales_email"`
			Status          interface{} `json:"status" structs:"status"`
			LastModified    interface{} `json:"last_modified" structs:"last_modified"`
			Resources       []Resource  `json:"resources" structs:"resources"`
		}

		var tempVal JsonDataOpportunity
		data, err := c.GetRawData()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "failed get data",
			})
			c.Abort()
			return
		}

		if data != nil {
			errors := make(map[string]string)

			err := json.Unmarshal(data, &tempVal)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   err.Error(),
					"message": "failed to unmarshal",
				})
				c.Abort()
				return
			}

			mapTemp := structs.Map(tempVal)
			interResources := mapTemp["resources"]
			delete(mapTemp, "resources")

			for key, value := range mapTemp {
				if reflect.TypeOf(value).Kind() != reflect.String {
					if reflect.TypeOf(value).Kind() == reflect.Float64 {
						errors[key] = "invalid type. Expected string but got number"
					} else {
						errors[key] = fmt.Sprintf("invalid type. Expected string but got %T", value)
					}
				}
			}

			resources, ok := interResources.([]interface{})
			if !ok {
				fmt.Printf("invalid type. Expected []map[string]interface{} but got %T", interResources)
			}

			for idx, res := range resources {
				resource, ok := res.(map[string]interface{})

				if !ok {
					fmt.Println("Resource is not a map")
				}

				for key, value := range resource {
					switch key {
					case "project_duration", "qty", "ctc":
						if value != nil {
							if reflect.TypeOf(value).Kind() != reflect.Float64 {
								errors["resources."+strconv.Itoa(idx)+"."+key] = fmt.Sprintf("invalid type. Expected number but got %T", value)
							}
						}
					case "position", "level":
						if reflect.TypeOf(value).Kind() != reflect.String {
							if reflect.TypeOf(value).Kind() == reflect.Float64 {
								errors["resources."+strconv.Itoa(idx)+"."+key] = "invalid type. Expected string but got number"
							} else {
								errors["resources."+strconv.Itoa(idx)+"."+key] = fmt.Sprintf("invalid type. Expected string but got %T", value)
							}
						}
					default:
						errors["resources"] = "invalid type"
					}
				}
			}
			if len(errors) > 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   errors,
					"message": "invalid data type",
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
