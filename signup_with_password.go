app.Post("/signup", func(c *fiber.Ctx) error {
		// Mengambil data langsung dari form-data atau x-www-form-urlencoded
		username := c.FormValue("username")
		password := c.FormValue("password")

		if username == "" || password == "" {
			return c.Status(400).SendString("Username dan password harus diisi")
		}

		// Hash password
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
		
		// Simpan ke DB
		userDatabase[username] = string(hashedPassword)

		return c.SendString("User " + username + " berhasil didaftarkan!")
})
