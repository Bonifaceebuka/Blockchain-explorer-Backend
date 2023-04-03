package controllers

import (
	"github.com/Bonifaceebuka/Blockchain-explorer-Backend/config"
	"github.com/gofiber/fiber/v2"
)

func init() {
	// DBConnection = config.GetDBConnection()
	config.LoadEnv()
}

func Home(c *fiber.Ctx) error {
	c.Status(200)

	return c.JSON(fiber.Map{
		"msg": "API service is fully up!",
	})

}

// func Signup(c *fiber.Ctx) error {
// 	formData, err := c.MultipartForm()
// 	emailRegex := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
// 	if err != nil {
// 		log.Fatal("Unable to parse form data")
// 	}

// 	name := strings.TrimSpace(formData.Value["name"][0])
// 	email := strings.TrimSpace(formData.Value["email"][0])
// 	password := formData.Value["password"][0]

// 	if len(strings.ReplaceAll(name, " ", "")) == 0 {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"msg": "Your name is required",
// 		})
// 	}

// 	if len(email) == 0 {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"msg": "Your email is required",
// 		})
// 	}

// 	if !emailRegex.MatchString(email) {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"msg": "Invalid email address was supplied",
// 		})
// 	}
// 	if len(password) == 0 {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"msg": "Your password is required",
// 		})
// 	}

// 	if len(password) < 3 {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"msg": "Your password must be more than 3 characters",
// 		})
// 	}

// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	// image := userData["image"]
// 	user := models.User{
// 		Name:     name,
// 		Email:    email,
// 		Password: hashedPassword,
// 	}

// 	DBConnection.Create(&user)
// 	c.Status(201)
// 	return c.JSON(fiber.Map{
// 		"msg": "Account created successfully!",
// 	})
// }
