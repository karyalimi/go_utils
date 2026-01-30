app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Ambil hash dari DB berdasarkan username
		storedHash, exists := userDatabase[username]
		if !exists {
			return c.Status(401).SendString("Username atau password salah")
		}

		// Bandingkan hash dengan password yang diinput
		err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
		if err != nil {
			return c.Status(401).SendString("Username atau password salah")
		}

		return c.SendString("Selamat datang, " + username + "!")
})
