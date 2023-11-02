package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var root string = "C:\\Users\\bryan\\Documents\\TSwiftProjects"
var filesOpened map[string]string = map[string]string{}
var num int = len(filesOpened) - 1
var name string = "NULL"
var mu sync.Mutex

type Files struct{}

type OpenFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type SetCurrentN struct {
	Num  int    `json:"num"`
	Name string `json:"name"`
}

type Close struct {
	Name string `json:"name"`
}

type Save struct {
	Content string `json:"content"`
}

func (f *Files) OpenFile(ctx *fiber.Ctx) error {
	var reqBody OpenFile
	if err := ctx.BodyParser(&reqBody); err != nil {
		return err
	}
	name = reqBody.Name
	content := reqBody.Content
	filePath := fmt.Sprintf("%v\\%v", root, name)
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"response": err.Error(),
		})
	}
	key := fmt.Sprintf("%v_%v", root, name)
	filesOpened[strings.Replace(key, "\\", ":", -1)] = content
	num = len(filesOpened) - 1
	return ctx.JSON(fiber.Map{
		"num": num,
	})
}

func (f *Files) OpenedFiles(ctx *fiber.Ctx) error {
	response, err := json.Marshal(filesOpened)
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"response": fmt.Sprintf("%v", string(response)),
	})
}

func (f *Files) SetCurrentN(ctx *fiber.Ctx) error {
	var reqBody SetCurrentN
	if err := ctx.BodyParser(&reqBody); err != nil {
		return err
	}
	num = reqBody.Num - 1
	name = reqBody.Name
	return ctx.JSON(fiber.Map{
		"num": num,
	})
}

func (f *Files) GetCurrentN(ctx *fiber.Ctx) error {
	return ctx.JSON(&fiber.Map{
		"num": num,
	})
}

func (f *Files) Close(ctx *fiber.Ctx) error {
	var reqBody Close
	if err := ctx.BodyParser(&reqBody); err != nil {
		return err
	}
	delete(filesOpened, strings.Replace(fmt.Sprintf("%v_%v", root, reqBody.Name), "\\", ":", -1))
	if len(filesOpened) > 0 {
		num = 0
		for k := range filesOpened {
			name = strings.Split(k, "_")[0]
			break
		}
	} else {
		num = -1
		name = "NULL"
	}
	return ctx.JSON(fiber.Map{
		"num": num,
	})
}

func (f *Files) Save(ctx *fiber.Ctx) error {
	var reqBody Save
	if err := ctx.BodyParser(&reqBody); err != nil {
		return err
	}
	if name == "NULL" {
		return ctx.JSON(fiber.Map{
			"response": "unsaved",
		})
	}
	filePath := filepath.Join(root, name)
	err := ioutil.WriteFile(filePath, []byte(reqBody.Content), 0644)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"response": err.Error(),
		})
	}
	key := fmt.Sprintf("%v_%v", root, name)
	filesOpened[strings.Replace(key, "\\", ":", -1)] = reqBody.Content
	return ctx.JSON(fiber.Map{
		"response": "saved",
	})
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
