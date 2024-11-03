package main

import "harupdf"

func main() {
	pdfdoc, err := harupdf.New()
	if err != nil {
		panic(err)
	}
	defer pdfdoc.Close()

	page := pdfdoc.NewPage()

	page.AddText("Titulo sobre el documento", 24)
	for i := 0; i < 25; i++ {
		page.AddText("Huevos", 16)
		page.AddLine()
	}

	pdfdoc.SaveFile("demo.pdf")

}
