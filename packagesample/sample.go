package packagesample

func PrintExemplo() {
	print("teste exemplo package\n")
}

//function not visible out package
func printExemplo() {
	print("teste exemplo package\n")
}

type House struct {
	QtdRoms int
	Meters  int
}

//start low case visible intern package
type houseinternt struct {
	QtdRoms int
	//this property is visible just internally
	meters int
}

//construct start low case, so it's visible just internally
func (h House) construct() {

}
