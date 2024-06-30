package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HelloGo(c *fiber.Ctx) error {
	return c.SendString("hello world!")
}

func Dududumdum(c *fiber.Ctx) error {
	return c.SendStatus(fiber.ErrBadRequest.Code)
}

func Findtc(c *fiber.Ctx)error {
	storyID := c.Query("StoryID")
	version := c.Query("version")
	applicationName := c.Query("applicationName")
	description := c.Query("description")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	if(storyID!=""||version!=""||applicationName!=""||description!=""||startDate!=""||endDate!=""){
		fmt.Println("have")
	}else{
		fmt.Println("don't have")
	}
	
	fmt.Println(c.Queries())
	return c.Status(fiber.StatusOK).JSON(c.Queries())
}

func Addtc(c *fiber.Ctx)error{
	return c.SendStatus(fiber.StatusOK)
}

func Edittc(c *fiber.Ctx)error{
	return c.SendStatus(fiber.StatusOK)
}

func Deletetc(c *fiber.Ctx)error{
	return c.SendStatus(fiber.StatusOK)
}

