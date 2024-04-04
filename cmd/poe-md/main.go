package main

import "poe-ninja-md-db/internals/poeninja"

func main() {
	hc := poeninja.NewHttpClient("Necropolis")
	tattoos, err := hc.GetItemsOverview(poeninja.ItemTypeTattoo)
	if err != nil {
		panic(err)
	}

	for _, tattoo := range tattoos.Lines {
		println(tattoo.Name)
	}
}
