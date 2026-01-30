// --- REGISTER / SIGNUP ---
	app.Post("/signup", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Hash password
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

		// Simpan ke database menggunakan GORM
		newUser := User{
			Username: username,
			Password: string(hashedPassword),
		}

		result := DB.Create(&newUser)
		if result.Error != nil {
			return c.Status(500).SendString("Gagal mendaftarkan user (mungkin username sudah ada)")
		}

		return c.SendString("User " + username + " berhasil disimpan ke database!")
	})
