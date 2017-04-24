package PosMenu

type posFuncs struct {
	NewCheck    funcType
	RecallCheck funcType
	ClockIn     funcType
	ClockOut    funcType
}

func GetMainMenu() posFuncs {
	var menu posFuncs
	menu.NewCheck.Name = "New Check"
	menu.RecallCheck.Name = "Recal Check"
	menu.ClockIn.Name = "Clock In"
	menu.ClockOut.Name = "Clock Out"
	return menu
}

type funcType struct {
	Active  bool
	Visible bool
	Name    string
}
