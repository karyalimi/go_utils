//time
	now := time.Now().UTC()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	wibTime := now.In(loc)

                    //usage
                    "time":        wibTime.Format("02/01/2006 15:04:05"),
