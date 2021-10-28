// Code generated by 'yaegi extract github.com/XiaoMengXinX/Music163Api-Go/utils'. DO NOT EDIT.

package symbols

import (
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"reflect"
)

func init() {
	Symbols["github.com/XiaoMengXinX/Music163Api-Go/utils/utils"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ApiRequest":       reflect.ValueOf(utils.ApiRequest),
		"CacheKeyEncrypt":  reflect.ValueOf(utils.CacheKeyEncrypt),
		"ChooseUserAgent":  reflect.ValueOf(utils.ChooseUserAgent),
		"CreateNewRequest": reflect.ValueOf(utils.CreateNewRequest),
		"DetectFileType":   reflect.ValueOf(utils.DetectFileType),
		"EapiDecrypt":      reflect.ValueOf(utils.EapiDecrypt),
		"EapiEncrypt":      reflect.ValueOf(utils.EapiEncrypt),
		"EapiRequest":      reflect.ValueOf(utils.EapiRequest),
		"Format2Params":    reflect.ValueOf(utils.Format2Params),
		"ImageSize":        reflect.ValueOf(utils.ImageSize),
		"MarkerDecrypt":    reflect.ValueOf(utils.MarkerDecrypt),
		"MarkerEncrypt":    reflect.ValueOf(utils.MarkerEncrypt),
		"RandHex":          reflect.ValueOf(utils.RandHex),
		"RawRequest":       reflect.ValueOf(utils.RawRequest),
		"ReadFile":         reflect.ValueOf(utils.ReadFile),
		"SpliceStr":        reflect.ValueOf(utils.SpliceStr),

		// type definitions
		"EapiOption":  reflect.ValueOf((*utils.EapiOption)(nil)),
		"Headers":     reflect.ValueOf((*utils.Headers)(nil)),
		"RequestData": reflect.ValueOf((*utils.RequestData)(nil)),
	}
}