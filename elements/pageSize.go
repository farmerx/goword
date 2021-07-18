package elements

type PageSize struct {
	// Page Width
	WidthAttr float64
	// Page Height
	HeightAttr float64
	// Page Orientation
	OrientAttr PageOrientation
	// Printer Paper Code
	CodeAttr *int64
}
