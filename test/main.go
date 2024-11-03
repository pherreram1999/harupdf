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
	page.AddText("Me canso ganzo mendez mendoza", 14)
	page.AddText("Ostia joder", 14)
	page.AddText("Esto es otro parrafo demostrativo", 14)

	pdfdoc.SaveFile("demo.pdf")

}
