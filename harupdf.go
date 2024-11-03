package harupdf

/*
#cgo LDFLAGS: -lhpdf
#include <hpdf.h>
*/
import "C"
import (
	"errors"
)

type (
	CursorPosition struct {
		x, y C.HPDF_REAL
	}

	PDFDocument struct {
		cPdf  C.HPDF_Doc
		cFont C.HPDF_Font
	}

	Page struct {
		pdfDoc         *PDFDocument
		cursorPosition CursorPosition
		cPage          C.HPDF_Page
	}
)

func New() (*PDFDocument, error) {
	cPdf := C.HPDF_New(nil, nil)
	if cPdf == nil {
		return nil, errors.New("No fue posible crear la instacia HDPF")
	}

	return &PDFDocument{
		cPdf:  cPdf,
		cFont: C.HPDF_GetFont(cPdf, C.CString("Helvetica"), nil),
	}, nil
}

func (doc *PDFDocument) SaveFile(path string) {
	C.HPDF_SaveToFile(doc.cPdf, C.CString(path))
}

func (p *PDFDocument) SetFontSize() {

}

func (p *PDFDocument) cGetPDF() C.HPDF_Doc {
	return p.cPdf
}

func (p *PDFDocument) cGetFont() C.HPDF_Font {
	return p.cFont
}

func (p *PDFDocument) Close() {
	C.HPDF_Free(p.cPdf)
}

func (p *PDFDocument) NewPage() *Page {
	cpage := C.HPDF_AddPage(p.cPdf)
	return &Page{
		pdfDoc: p,
		cPage:  cpage,
		cursorPosition: CursorPosition{
			y: C.HPDF_Page_GetHeight(cpage),
		},
	}
}

// page
func (p *Page) GetDocument() *PDFDocument {
	return p.pdfDoc
}

func (p *Page) SetFontSize(size float32) {
	cfont := p.GetDocument().cGetFont()
	C.HPDF_Page_SetFontAndSize(p.cPage, cfont, C.float(size))
}

func (p *Page) AddText(text string, size float32) {
	/*
		las posiciones para y es de abajo hacia arriba
		mientras que x es de izquierda a derecha
	*/
	p.SetFontSize(size)
	C.HPDF_Page_BeginText(p.cPage)
	// con esto vamos actualizado la ubicacion del ultimo elemento
	p.cursorPosition.y -= C.HPDF_REAL(size)
	C.HPDF_Page_TextOut(
		p.cPage,
		0,
		p.cursorPosition.y,
		C.CString(text),
	)
	C.HPDF_Page_EndText(p.cPage)
}
