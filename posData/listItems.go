package posData

// Item exports items received from Google Spreadsheet
type Item struct {
	ItemID    string
	ItemName  string
	ItemPrice string
}

// ListItems returns slices of items
func ListItems() []Item {

	v := Start()
	var m Item
	var iList []Item

	for _, Row := range v {
		m.ItemID = Row[0].(string)
		m.ItemName = Row[1].(string)
		m.ItemPrice = Row[2].(string)
		iList = append(iList, m)

	}

	return iList
}
