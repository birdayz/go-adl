package main

//#include "adl_defines.h"
import "C"

const (
	ADL_OK = 0
)

type Adl struct{}

type AdapterInfo struct {
	ISize            int
	IAdapterIndex    int
	StrUDID          string
	IBusNumber       int
	IDeviceNumber    int
	IFunctionNumber  int
	IVendorID        int
	StrAdapterName   string
	StrDisplayName   string
	IPresent         int
	IExist           int
	StrDriverPath    string
	StrDriverPathExt string
	StrPNPString     string
	IOSDisplayIndex  int
}

func convertString(in [C.ADL_MAX_DISPLAY_NAME]C.char) string {
	return C.GoString(&in[0])
}
