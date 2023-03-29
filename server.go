package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"gali.com/config"
	"gali.com/model"
	"github.com/labstack/echo/v4"
)

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}

func main() {
	config.ConnectDB()
	route := echo.New()
	route.POST("customer/create_customer", func(c echo.Context) error {
		costomer := new(model.Customers)
		c.Bind(costomer)
		contentType := c.Request().Header.Get("Content-type")
		if contentType == "application/json" {
			fmt.Println("Request dari json")
			config.DB.Create(costomer)
		} else if strings.Contains(contentType, "multipart/form-data") || contentType == "application/x-www-form-urlencoded" {
			file, err := c.FormFile("custmer_name")
			if err != nil {
				fmt.Println("CostmerName Kosong")
			} else {
				src, err := file.Open()
				if err != nil {
					return err
				}
				defer src.Close()
				dst, err := os.Create(file.Filename)
				if err != nil {
					return err
				}
				defer dst.Close()
				if _, err = io.Copy(dst, src); err != nil {
					return err
				}
				costomer.CustmerName = file.Filename
				fmt.Println("ada file, akan disimpan")
			}
			config.DB.Create(costomer)
		}
		response := struct {
			Message string
			Data    model.Customers
		}{
			Message: "Sukses create custumer",
			Data:    *costomer,
		}
		return c.JSON(http.StatusOK, response)

	})

	route.PUT("customer/update_customer/:customer_code", func(c echo.Context) error {
		customer := new(model.Customers)
		c.Bind(customer)
		response := new(Response)
		if customer.UpdateCustomer(c.Param("customer_code")) != nil {
			response.ErrorCode = 10
			response.Message = "Gagal Update data Customer"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses update data customer"
			response.Data = *customer
			config.DB.First(&customer)
		}
		return c.JSON(http.StatusOK, response)
	})

	// route.DELETE("customer/delete_customer/:customer_code", func(c echo.Context) error {
	// 	customer, _ := model.GetOneByCode(c.Param("customer_code"))
	// 	response := new(Response)

	// 	if customer.DeleteCustomer() != nil {
	// 		response.ErrorCode = 10
	// 		response.Message = "Gagal menghapus data customer"
	// 	} else {
	// 		response.ErrorCode = 0
	// 		response.Message = "sukses menghapus data customer"
	// 		config.DB.Save(&customer)
	// 	}
	// 	return c.JSON(http.StatusOK, response)
	// })
	route.DELETE("customer/delete_customer/:customer_code", func(c echo.Context) error {
		customer := new(model.Customers)
		c.Bind(customer)
		response := new(Response)
		if customer.DeleteByCode(c.Param("customer_code")) != nil {
			response.ErrorCode = 10
			response.Message = "Gagal hapus data Customer"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses hapus data customer"
			// response.Data = *customer
			config.DB.First(&customer)
		}
		return c.JSON(http.StatusOK, response)
	})

	route.GET("customer/search_customer", func(c echo.Context) error {
		response := new(Response)
		costomers, err := model.GetAll(c.QueryParam("keywords"))
		if err != nil {
			response.ErrorCode = 10
			response.Message = "Gagal melihat data customer"
		} else {
			response.ErrorCode = 10
			response.Message = "Sukses melihat data customer"
			response.Data = costomers

		}
		return c.JSON(http.StatusOK, response)
	})
	route.Start(":9000")
}
