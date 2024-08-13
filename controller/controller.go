package controller

import (
	"fmt"
	"log"

	"tcstorego/database"
	"tcstorego/model"

	"github.com/gofiber/fiber/v2"
)

func HelloGo(c *fiber.Ctx) error {
	return c.SendString("hello world!")
}

func Dududumdum(c *fiber.Ctx) error {
	return c.SendStatus(fiber.ErrBadRequest.Code)
}

func Findtc(c *fiber.Ctx) error {
	storyID := c.Query("StoryID")
	version := c.Query("version")
	applicationName := c.Query("applicationName")
	description := c.Query("description")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	if storyID != "" || version != "" || applicationName != "" || description != "" || startDate != "" || endDate != "" {
		//condition search
		fmt.Println("have")

	} else {

		//search all data
		fmt.Println("don't have")
	}

	fmt.Println(c.Queries())
	return c.Status(fiber.StatusOK).JSON(c.Queries())
}

func Addtc(c *fiber.Ctx) error {
	tc := new(model.Testcase)
	if err := c.BodyParser(tc); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if tc.StoryID == "" || tc.ApplicationName == "" || tc.FileName == "" {
		return c.Status(400).JSON("Required field are missing")
	}
	result := database.DBCon.Create(&tc)
	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}
	return c.SendStatus(fiber.StatusOK)
}

func Edittc(c *fiber.Ctx) error {
	tc := new(model.Testcase)
	if err := c.BodyParser(tc); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if tc.StoryID == "" || tc.ApplicationName == "" || tc.FileName == "" {
		return c.Status(400).JSON("Required field are missing")
	}
	updateTC := model.Testcase{
		StoryID:         tc.StoryID,
		Version:         tc.Version,
		ApplicationName: tc.ApplicationName,
		Description:     tc.Description,
		Date:            tc.Date,
	}

	database.DBCon.Model(&model.Testcase{}).Where("TestCaseID = ?", tc.TestCaseID).Updates(updateTC)
	//This function has not been tested yet.
	return c.SendStatus(fiber.StatusOK)
}
//TODO: check c.Query that can recieve value from front end
func Deletetc(c *fiber.Ctx) error {
	TcID := c.Query("TestCaseID")
	if TcID != "" {
		log.Println("it have value", TcID)
		database.DBCon.Delete(&model.Testcase{},TcID)
	}
	log.Panicln("did not find :DELETE method")
	return c.SendStatus(fiber.StatusOK)

}

