package controller

import (
	"fmt"
	"log"
	"time"

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

// TODO: find the best way to recive file and data mutipart form data? or just 2 api? (2 api is easier way i think)
func Addtc(c *fiber.Ctx) error {
	/*tc := new(model.Testcase)
	if err := c.BodyParser(tc); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if tc.StoryID == "" || tc.ApplicationName == "" || tc.FileName == "" {
		return c.Status(400).JSON("Required field are missing")
	}
	result := database.DBCon.Create(&tc)
	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}*/
	form, err := c.MultipartForm();
	if err != nil {
		log.Println("Error read MultipartForm: ", err)
		return c.Status(400).JSON("Invalid multipart form data")
	}
	files := form.File["File"]
	if len(files) == 0 {
		return c.Status(400).JSON("No file uploaded with the field 'File'")
	}
	file := files[0] //First file! I mean that only one file sent from user. :)
	c.SaveFile(file,"./excelFile/"+file.Filename)
	
	return c.SendStatus(fiber.StatusOK)
}

func extractformTC(values map[string][]string) *model.Testcase {
	tc := &model.Testcase{
		StoryID: values["StoryID"][0],
		ApplicationName: values["ApplicationName"][0],
		FileName: values["FileName"][0],
	}
	if v, ok := values["Version"]; ok{
		tc.Version = &v[0]
	}
	if d, ok := values["Date"]; ok {
		date, _ := time.Parse(time.RFC3339, d[0])
		tc.Date = &date
	}

	if desc, ok := values["Description"]; ok {
		tc.Description = &desc[0]
	}

	return tc
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

// TODO: check c.Query that can recieve value from front end
func Deletetc(c *fiber.Ctx) error {
	TcID := c.Query("TestCaseID")
	if TcID != "" {
		log.Println("it have value", TcID)
		database.DBCon.Delete(&model.Testcase{}, TcID)
	}
	log.Println("did not find value in :DELETE method")
	return c.SendStatus(fiber.StatusOK)

}
