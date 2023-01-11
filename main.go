package main

func main() {
	_Item := NewItem("Innosilicon A10 pro 7GB 720 Mh")

	observerFirst := &Customer{id: "catdog240195@gmail.com"}
	observerSecond := &Customer{id: "maksim_moroxov@gmail.com"}

	_Item.Register(observerFirst)
	_Item.Register(observerSecond)

	_Item.UpdateAvailability()
}
