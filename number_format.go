func NumberFormat(n int) string {
	p := message.NewPrinter(message.MatchLanguage("id"))
	return p.Sprintf("%d", n)
}

#usage
engine.AddFunc("NumberFormat", helper.NumberFormat)
