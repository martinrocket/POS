package PosMenu

type posFuncs struct {
	NewCheck    funcType
	RecallCheck funcType
	ClockIn     funcType
	ClockOut    funcType
}

type funcType struct {
	Active  bool
	Visible bool
	Name    string
}

func GetMainMenu() posFuncs {
	var menu posFuncs
	menu.NewCheck = funcType{true, true, "New Check"}
	menu.RecallCheck = funcType{true, true, "Recall Check"}
	menu.ClockIn = funcType{true, true, "Clock In"}
	menu.ClockOut = funcType{false, true, "Clock Out"}
	return menu
}
