app.Post("/import", func(c *fiber.Ctx) error {
		// 1. Ambil file dari form-data
		fileHeader, err := c.FormFile("document")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "File tidak ditemukan"})
		}

		// 2. Buka file yang diupload
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal membuka file"})
		}
		defer file.Close()

		// 3. Baca file Excel menggunakan Excelize
		f, err := excelize.OpenReader(file)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Format Excel tidak valid"})
		}
		defer f.Close()

		// Ambil nama sheet pertama
		sheetName := f.GetSheetName(0)
		
		// 4. Ambil semua baris dari sheet tersebut
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal membaca baris"})
		}

		var dataUsers []User

		// 5. Iterasi baris (asumsi baris 1 adalah Header, jadi kita mulai dari index 1)
		for i, row := range rows {
			if i == 0 || len(row) < 2 { // Skip header atau baris kosong
				continue
			}

			user := User{
				NIK:  row[0], // Kolom A
				Nama: row[1], // Kolom B
			}
			dataUsers = append(dataUsers)
			
			// Di sini kamu bisa tambahkan logika simpan ke Database (GORM/SQL)
			fmt.Printf("Importing: %s - %s\n", user.NIK, user.Nama)
		}

		return c.JSON(fiber.Map{
			"message": "Import berhasil",
			"total":   len(dataUsers),
			"data":    dataUsers,
		})
	})
