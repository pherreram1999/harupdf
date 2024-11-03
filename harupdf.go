package harupdf

/*
#cgo LDFLAGS: -lhpdf
#include <hpdf.h>

extern void errorHandler(HPDF_STATUS error_no, HPDF_STATUS detail_no, void *user_data);


// con la funcion expuesta desde go, se la podemos pasar C
static HPDF_Doc HPDF_New_WithCallback() {
    return HPDF_New(errorHandler, NULL);
}
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type (
	CursorPosition struct {
		x, y, lastSize C.HPDF_REAL
	}

	PDFDocument struct {
		cPdf  C.HPDF_Doc
		cFont C.HPDF_Font
	}

	Page struct {
		pdfDoc         *PDFDocument
		cursorPosition CursorPosition
		PaddingLeft    C.HPDF_REAL
		PaddingRight   C.HPDF_REAL
		Height, Width  C.HPDF_REAL
		cPage          C.HPDF_Page
	}
)

// necesita export para que puede ser llamada desde C

//export errorHandler
func errorHandler(error_no, detail_no C.HPDF_STATUS, userdata unsafe.Pointer) {
	errorMsg := GetErrorMessage(int(error_no))
	fmt.Printf("PDF Error: %s (Code: 0x%04X, Detail: 0x%04X)\n",
		errorMsg, error_no, detail_no)
}

func New() (*PDFDocument, error) {
	cPdf := C.HPDF_New_WithCallback()
	if cPdf == nil {
		return nil, errors.New("No fue posible crear la instacia HDPF")
	}

	return &PDFDocument{
		cPdf:  cPdf,
		cFont: C.HPDF_GetFont(cPdf, C.CString("Helvetica"), nil),
	}, nil
}

func (doc *PDFDocument) SaveFile(path string) error {
	if C.HPDF_SaveToFile(doc.cPdf, C.CString(path)) != C.HPDF_OK {
		errors.New("No fue posible escribir el archivo")
	}

	return nil
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

func (doc *PDFDocument) NewPage() *Page {
	cpage := C.HPDF_AddPage(doc.cPdf)
	p := &Page{
		pdfDoc:       doc,
		cPage:        cpage,
		Height:       C.HPDF_Page_GetHeight(cpage),
		Width:        C.HPDF_Page_GetWidth(cpage),
		PaddingLeft:  10,
		PaddingRight: 10,
	}

	p.cursorPosition.y = p.Height
	p.cursorPosition.x = p.Width

	return p
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
	p.cursorPosition.lastSize = C.HPDF_REAL(size)
	p.cursorPosition.y -= p.cursorPosition.lastSize
	C.HPDF_Page_TextOut(
		p.cPage,
		p.PaddingLeft,
		p.cursorPosition.y,
		C.CString(text),
	)
	C.HPDF_Page_EndText(p.cPage)
}

func (p *Page) AddLine() {
	// 							(x, y)
	pad := p.cursorPosition.lastSize / 2
	//p.cursorPosition.y -= p.cursorPosition.lastSize
	C.HPDF_Page_MoveTo(p.cPage, p.PaddingLeft, p.cursorPosition.y-pad)                     // origen
	C.HPDF_Page_LineTo(p.cPage, p.cursorPosition.x-p.PaddingRight, p.cursorPosition.y-pad) // final
	C.HPDF_Page_Stroke(p.cPage)
	p.cursorPosition.y -= p.cursorPosition.lastSize - pad
}
