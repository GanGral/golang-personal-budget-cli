package module1

// Budget stores budget information
var Bugdet struct {
	Max   float32
	Items []Item
}

// Item stores item information
var Item struct {
	Description string
	Price       float32
}
