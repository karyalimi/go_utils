// --- LOGIN ---
app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		var user User
		// Cari user berdasarkan username di database
		if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
			return c.Status(401).SendString("Username tidak ditemukan")
		}

		// Bandingkan password input dengan hash di DB
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return c.Status(401).SendString("Password salah")
		}

		return c.SendString("Login Berhasil! ID User Anda: " + fmt.Sprint(user.ID))
	})
