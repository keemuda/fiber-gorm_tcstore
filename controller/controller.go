package controller

import (
	"fmt"
	"log"
	"os"
	"time"

	"tcstorego/database"
	"tcstorego/model"

	"github.com/gofiber/fiber/v2"
)

func FindTestcase(c *fiber.Ctx) error {
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
		testcase := new(model.Testcase)
		result := database.DBCon.Find(&testcase)
		log.Println(result.RowsAffected)
		log.Println(testcase)
		if result.Error != nil {
			return c.Status(500).JSON(result.Error)
		}
		return c.Status(200).JSON(testcase)
	}

	fmt.Println(c.Queries())
	return c.Status(fiber.StatusOK).JSON(c.Queries())
}

func AddTestcase(c *fiber.Ctx) error {
	testcase := new(model.Testcase)
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error read MultipartForm: ", err)
		return c.Status(400).JSON("Invalid multipart form data")
	}

	testcase = extractformTestcase(form.Value)
	if testcase.StoryID == "" || testcase.ApplicationName == "" || testcase.FileName == "" {
		return c.Status(400).JSON("Required field are missing")
	}

	files := form.File["File"]
	if len(files) == 0 {
		return c.Status(400).JSON("No file uploaded with the field 'File'")
	}

	result := database.DBCon.Create(&testcase)
	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}
	file := files[0]
	c.SaveFile(file, "./excelFile/"+file.Filename)

	return c.SendStatus(fiber.StatusCreated)
}

func extractformTestcase(values map[string][]string) *model.Testcase {
	tc := &model.Testcase{
		StoryID:         values["StoryID"][0],
		ApplicationName: values["ApplicationName"][0],
		FileName:        values["FileName"][0],
	}
	if v, ok := values["Version"]; ok {
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

func EditTestcase(c *fiber.Ctx) error {
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
	return c.SendStatus(fiber.StatusOK)
}

func DeleteTestcase(c *fiber.Ctx) error {
	TestCaseID := c.Query("TestCaseID")
	if TestCaseID != "" {
		log.Println("it have value", TestCaseID)
		database.DBCon.Delete(&model.Testcase{}, TestCaseID)
	}
	log.Println("deleteTestCase: did not find value TestCaseID")
	return c.Status(400).JSON("Required field are missing: TestCaseID")
}

func GetTestcaseFile(c *fiber.Ctx) error {
	fileName := c.Query("Filename")
	filepath := fmt.Sprintf("./excel/%s", fileName)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fmt.Sprintf("File (%s) not found", fileName))
	}
	return c.SendFile(filepath)
}
