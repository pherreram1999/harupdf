package harupdf

import "fmt"

const (
	// Array related errors
	ErrArrayCountErr           = 0x1001 // Internal error. Data consistency was lost
	ErrArrayItemNotFound       = 0x1002 // Internal error. Data consistency was lost
	ErrArrayItemUnexpectedType = 0x1003 // Internal error. Data consistency was lost

	// Binary and data errors
	ErrBinaryLengthErr = 0x1004 // Data length > HPDF_LIMIT_MAX_STRING_LEN
	ErrCannotGetPallet = 0x1005 // Cannot get pallet data from PNG image

	// Dictionary errors
	ErrDictCountErr             = 0x1007 // Dictionary elements > HPDF_LIMIT_MAX_DICT_ELEMENT
	ErrDictItemNotFound         = 0x1008 // Internal error. Data consistency was lost
	ErrDictItemUnexpectedType   = 0x1009 // Internal error. Data consistency was lost
	ErrDictStreamLengthNotFound = 0x100A // Internal error. Data consistency was lost

	// Document errors
	ErrDocEncryptDictNotFound = 0x100B // HPDF_SetEncryptMode() or HPDF_SetPermission() called before password set
	ErrDocInvalidObject       = 0x100C // Internal error. Data consistency was lost
	ErrDuplicateRegistration  = 0x100E // Tried to re-register a registered font
	ErrExceedJWWCodeNumLimit  = 0x100F // Cannot register a character to the Japanese word wrap characters list

	// Encryption errors
	ErrEncryptInvalidPassword = 0x1011 // Owner and user password issues
	ErrUnknownClass           = 0x1013 // Internal error. Data consistency was lost
	ErrExceedGstateLimit      = 0x1014 // Stack depth > HPDF_LIMIT_MAX_GSTATE

	// Memory and file errors
	ErrFailedToAllocMem = 0x1015 // Memory allocation failed
	ErrFileIOError      = 0x1016 // File processing failed
	ErrFileOpenError    = 0x1017 // Cannot open a file

	// Font errors
	ErrFontExists             = 0x1019 // Tried to load a font that has been registered
	ErrFontInvalidWidthsTable = 0x101A // Font-file format is invalid
	ErrInvalidAFMHeader       = 0x101B // Cannot recognize header of afm file
	ErrInvalidAnnotation      = 0x101C // Specified annotation handle is invalid

	// Image and color errors
	ErrInvalidBitPerComponent = 0x101E // Invalid bit-per-component for mask-image
	ErrInvalidCharMatricsData = 0x101F // Cannot recognize char-matrics-data of afm file
	ErrInvalidColorSpace      = 0x1020 // Invalid color space parameters

	// General validation errors
	ErrInvalidCompressionMode = 0x1021 // Invalid value in HPDF_SetCompressionMode()
	ErrInvalidDateTime        = 0x1022 // Invalid date-time value
	ErrInvalidDestination     = 0x1023 // Invalid destination handle
	ErrInvalidDocument        = 0x1025 // Invalid document handle
	ErrInvalidDocumentState   = 0x1026 // Invalid function for present state
	ErrInvalidEncoder         = 0x1027 // Invalid encoder handle
	ErrInvalidEncoderType     = 0x1028 // Wrong font and encoder combination

	// Encoding and encryption errors
	ErrInvalidEncodingName  = 0x102B // Invalid encoding name
	ErrInvalidEncryptKeyLen = 0x102C // Invalid encryption key length

	// Font definition errors
	ErrInvalidFontdefData = 0x102D // Invalid font handle or unsupported format
	ErrInvalidFontdefType = 0x102E // Internal error. Data consistency was lost
	ErrInvalidFontName    = 0x102F // Font not found

	// Image errors
	ErrInvalidImage    = 0x1030 // Unsupported image format
	ErrInvalidJPEGData = 0x1031 // Unsupported image format
	ErrInvalidNData    = 0x1032 // Cannot read postscript-name from afm

	// Object errors
	ErrInvalidObject    = 0x1033 // Invalid object or internal error
	ErrInvalidObjID     = 0x1034 // Internal error. Data consistency was lost
	ErrInvalidOperation = 0x1035 // Invalid operation with mask-image

	// Document structure errors
	ErrInvalidOutline   = 0x1036 // Invalid outline handle
	ErrInvalidPage      = 0x1037 // Invalid page handle
	ErrInvalidPages     = 0x1038 // Invalid pages handle
	ErrInvalidParameter = 0x1039 // Invalid parameter value

	// Image format errors
	ErrInvalidPNGImage      = 0x103B // Invalid PNG format
	ErrInvalidStream        = 0x103C // Internal error. Data consistency was lost
	ErrMissingFileNameEntry = 0x103D // Missing _FILE_NAME entry

	// TrueType errors
	ErrInvalidTTCFile  = 0x103F // Invalid TTC format
	ErrInvalidTTCIndex = 0x1040 // Invalid font index
	ErrInvalidWXData   = 0x1041 // Cannot read width-data from afm

	// Miscellaneous errors
	ErrItemNotFound     = 0x1042 // Internal error. Data consistency was lost
	ErrLibPNGError      = 0x1043 // PNGLIB error
	ErrNameInvalidValue = 0x1044 // Internal error. Data consistency was lost
	ErrNameOutOfRange   = 0x1045 // Internal error. Data consistency was lost

	// Page errors
	ErrPagesMissingKidsEntry   = 0x1049 // Internal error. Data consistency was lost
	ErrPageCannotFindObject    = 0x104A // Internal error. Data consistency was lost
	ErrPageCannotGetRootPages  = 0x104B // Internal error. Data consistency was lost
	ErrPageCannotRestoreGState = 0x104C // No graphics-states to restore
	ErrPageCannotSetParent     = 0x104D // Internal error. Data consistency was lost
	ErrPageFontNotFound        = 0x104E // Current font not set
	ErrPageInvalidFont         = 0x104F // Invalid font handle
	ErrPageInvalidFontSize     = 0x1050 // Invalid font size
	ErrPageInvalidGMode        = 0x1051 // Invalid graphics mode
	ErrPageInvalidIndex        = 0x1052 // Internal error. Data consistency was lost
	ErrPageInvalidRotateValue  = 0x1053 // Rotation not multiple of 90
	ErrPageInvalidSize         = 0x1054 // Invalid page size
	ErrPageInvalidXObject      = 0x1055 // Invalid image handle
	ErrPageOutOfRange          = 0x1056 // Value out of range

	// Additional errors
	ErrRealOutOfRange       = 0x1057 // Value out of range
	ErrStreamEOF            = 0x1058 // Unexpected EOF
	ErrStreamReadlnContinue = 0x1059 // Internal error. Data consistency was lost
	ErrStringOutOfRange     = 0x105B // Text too long
	ErrThisFuncWasSkipped   = 0x105C // Function skipped due to errors

	// TrueType font errors
	ErrTTFCannotEmbeddingFont = 0x105D // Cannot embed font (license)
	ErrTTFInvalidCmap         = 0x105E // Cannot find unicode cmap
	ErrTTFInvalidFormat       = 0x105F // Unsupported TTF format
	ErrTTFMissingTable        = 0x1060 // Missing necessary table

	// Unsupported features
	ErrUnsupportedFontType   = 0x1061 // Internal error. Data consistency was lost
	ErrUnsupportedFunc       = 0x1062 // Function not supported
	ErrUnsupportedJPEGFormat = 0x1063 // Unsupported JPEG format
	ErrUnsupportedType1Font  = 0x1064 // Invalid PFB file

	// Additional document errors
	ErrXrefCountErr         = 0x1065 // Internal error. Data consistency was lost
	ErrZlibError            = 0x1066 // ZLIB error
	ErrInvalidPageIndex     = 0x1067 // Invalid page index
	ErrInvalidURI           = 0x1068 // Invalid URI
	ErrPagelayoutOutOfRange = 0x1069 // Invalid page layout

	// Page mode and annotation errors
	ErrPagemodeOutOfRange      = 0x1070 // Invalid page mode
	ErrPageNumStyleOutOfRange  = 0x1071 // Invalid page number style
	ErrAnnotInvalidIcon        = 0x1072 // Invalid annotation icon
	ErrAnnotInvalidBorderStyle = 0x1073 // Invalid border style
	ErrPageInvalidDirection    = 0x1074 // Invalid page direction
	ErrInvalidFont             = 0x1075 // Invalid font handle
)

func GetErrorMessage(code int) string {
	messages := map[int]string{
		// Array related errors
		ErrArrayCountErr:           "Internal error: Data consistency was lost in array count",
		ErrArrayItemNotFound:       "Internal error: Array item not found",
		ErrArrayItemUnexpectedType: "Internal error: Array item has unexpected type",

		// Binary and data errors
		ErrBinaryLengthErr: "Data length exceeds HPDF_LIMIT_MAX_STRING_LEN",
		ErrCannotGetPallet: "Cannot get pallet data from PNG image",

		// Dictionary errors
		ErrDictCountErr:             "Dictionary elements exceed HPDF_LIMIT_MAX_DICT_ELEMENT",
		ErrDictItemNotFound:         "Internal error: Dictionary item not found",
		ErrDictItemUnexpectedType:   "Internal error: Dictionary item has unexpected type",
		ErrDictStreamLengthNotFound: "Internal error: Stream length not found in dictionary",

		// Document errors
		ErrDocEncryptDictNotFound: "Encryption dictionary not found: Set password before encryption mode or permissions",
		ErrDocInvalidObject:       "Internal error: Invalid document object",
		ErrDuplicateRegistration:  "Attempted to register an already registered font",
		ErrExceedJWWCodeNumLimit:  "Cannot register character to Japanese word wrap characters list",

		// Encryption errors
		ErrEncryptInvalidPassword: "Invalid password: Either NULL owner password or matching owner/user passwords",
		ErrUnknownClass:           "Internal error: Unknown class",
		ErrExceedGstateLimit:      "Graphics state stack depth exceeds HPDF_LIMIT_MAX_GSTATE",

		// Memory and file errors
		ErrFailedToAllocMem: "Memory allocation failed",
		ErrFileIOError:      "File I/O operation failed",
		ErrFileOpenError:    "Cannot open file",

		// Font errors
		ErrFontExists:             "Attempted to load an already registered font",
		ErrFontInvalidWidthsTable: "Font file format is invalid or internal data consistency lost",
		ErrInvalidAFMHeader:       "Cannot recognize AFM file header",
		ErrInvalidAnnotation:      "Invalid annotation handle specified",

		// Image and color errors
		ErrInvalidBitPerComponent: "Invalid bit-per-component in mask image",
		ErrInvalidCharMatricsData: "Cannot recognize character metrics data in AFM file",
		ErrInvalidColorSpace:      "Invalid color space parameter or incompatible with mask image",

		// General validation errors
		ErrInvalidCompressionMode: "Invalid compression mode specified",
		ErrInvalidDateTime:        "Invalid date-time value specified",
		ErrInvalidDestination:     "Invalid destination handle specified",
		ErrInvalidDocument:        "Invalid document handle specified",
		ErrInvalidDocumentState:   "Function invalid in current document state",
		ErrInvalidEncoder:         "Invalid encoder handle specified",
		ErrInvalidEncoderType:     "Incompatible font and encoder combination",

		// Encoding and encryption errors
		ErrInvalidEncodingName:  "Invalid encoding name specified",
		ErrInvalidEncryptKeyLen: "Invalid encryption key length",

		// Font definition errors
		ErrInvalidFontdefData: "Invalid font handle or unsupported font format",
		ErrInvalidFontdefType: "Internal error: Invalid font definition type",
		ErrInvalidFontName:    "Font with specified name not found",

		// Image errors
		ErrInvalidImage:    "Unsupported image format",
		ErrInvalidJPEGData: "Unsupported JPEG format",
		ErrInvalidNData:    "Cannot read PostScript name from AFM file",

		// Object errors
		ErrInvalidObject:    "Invalid object specified or internal consistency lost",
		ErrInvalidObjID:     "Internal error: Invalid object ID",
		ErrInvalidOperation: "Invalid operation: Cannot set color mask for masked image",

		// Document structure errors
		ErrInvalidOutline:   "Invalid outline handle specified",
		ErrInvalidPage:      "Invalid page handle specified",
		ErrInvalidPages:     "Invalid pages handle specified",
		ErrInvalidParameter: "Invalid parameter value specified",

		// Image format errors
		ErrInvalidPNGImage:      "Invalid PNG image format",
		ErrInvalidStream:        "Internal error: Invalid stream",
		ErrMissingFileNameEntry: "Internal error: Missing _FILE_NAME entry for delayed loading",

		// TrueType errors
		ErrInvalidTTCFile:  "Invalid TrueType Collection file format",
		ErrInvalidTTCIndex: "TrueType Collection index exceeds number of included fonts",
		ErrInvalidWXData:   "Cannot read width data from AFM file",

		// Miscellaneous errors
		ErrItemNotFound:     "Internal error: Item not found",
		ErrLibPNGError:      "Error returned from PNGLIB while loading image",
		ErrNameInvalidValue: "Internal error: Invalid name value",
		ErrNameOutOfRange:   "Internal error: Name out of range",

		// Page errors
		ErrPagesMissingKidsEntry:   "Internal error: Missing kids entry in pages",
		ErrPageCannotFindObject:    "Internal error: Cannot find page object",
		ErrPageCannotGetRootPages:  "Internal error: Cannot get root pages",
		ErrPageCannotRestoreGState: "Cannot restore graphics state: No states to restore",
		ErrPageCannotSetParent:     "Internal error: Cannot set page parent",
		ErrPageFontNotFound:        "No current font is set",
		ErrPageInvalidFont:         "Invalid font handle specified for page",
		ErrPageInvalidFontSize:     "Invalid font size specified",
		ErrPageInvalidGMode:        "Invalid graphics mode for operation",
		ErrPageInvalidIndex:        "Internal error: Invalid page index",
		ErrPageInvalidRotateValue:  "Page rotation value must be multiple of 90",
		ErrPageInvalidSize:         "Invalid page size specified",
		ErrPageInvalidXObject:      "Invalid XObject/image handle specified",
		ErrPageOutOfRange:          "Specified value is out of range for page",

		// Additional errors
		ErrRealOutOfRange:       "Real number value is out of range",
		ErrStreamEOF:            "Unexpected end of file in stream",
		ErrStreamReadlnContinue: "Internal error: Stream read line continuation",
		ErrStringOutOfRange:     "String length exceeds limit",
		ErrThisFuncWasSkipped:   "Function skipped due to previous errors",

		// TrueType font errors
		ErrTTFCannotEmbeddingFont: "Cannot embed font due to license restrictions",
		ErrTTFInvalidCmap:         "Cannot find Unicode cmap in TrueType font",
		ErrTTFInvalidFormat:       "Unsupported TrueType font format",
		ErrTTFMissingTable:        "Required table missing in TrueType font",

		// Unsupported features
		ErrUnsupportedFontType:   "Internal error: Unsupported font type",
		ErrUnsupportedFunc:       "Function not supported or internal error",
		ErrUnsupportedJPEGFormat: "Unsupported JPEG format",
		ErrUnsupportedType1Font:  "Failed to parse Type1 font file (PFB)",

		// Additional document errors
		ErrXrefCountErr:         "Internal error: Cross-reference count error",
		ErrZlibError:            "Error in ZLIB processing",
		ErrInvalidPageIndex:     "Invalid page index specified",
		ErrInvalidURI:           "Invalid URI specified",
		ErrPagelayoutOutOfRange: "Page layout value is out of range",

		// Page mode and annotation errors
		ErrPagemodeOutOfRange:      "Page mode value is out of range",
		ErrPageNumStyleOutOfRange:  "Page number style is out of range",
		ErrAnnotInvalidIcon:        "Invalid annotation icon specified",
		ErrAnnotInvalidBorderStyle: "Invalid annotation border style specified",
		ErrPageInvalidDirection:    "Invalid page direction specified",
		ErrInvalidFont:             "Invalid font handle specified",
	}

	if msg, ok := messages[code]; ok {
		return msg
	}
	return "Unknown error code: " + fmt.Sprintf("0x%04X", code)
}
