//+build windows
package main

import (
	"errors"
	"syscall"
	"unsafe"
)

//#include <stdlib.h>
//#include <stdbool.h>
//#include <stdio.h>
//#include "adl_defines.h"
//#include "adl_structures.h"
//void* ADL_Main_Memory_Alloc(int iSize)
//{
//return malloc(iSize);
//}
import "C"

const (
	ADL_DLL = "atiadlxx.dll"
)

var (
	dll, _                              = syscall.LoadLibrary(ADL_DLL)
	ADL_Main_Control_Create, _          = syscall.GetProcAddress(dll, "ADL_Main_Control_Create")
	ADL_Adapter_NumberOfAdapters_Get, _ = syscall.GetProcAddress(dll, "ADL_Adapter_NumberOfAdapters_Get")
	ADL_Adapter_AdapterInfo_Get, _      = syscall.GetProcAddress(dll, "ADL_Adapter_AdapterInfo_Get") // , LPInfo, size
)

func (adl *Adl) ADL_Main_Control_Create() (err error) {
	result, _, errno := syscall.Syscall(ADL_Main_Control_Create, 2, uintptr(unsafe.Pointer(C.ADL_Main_Memory_Alloc)), 1, 0)
	if result != ADL_OK {
		return errors.New(errno.Error())
	}
	return
}

func (adl *Adl) ADL_Adapter_NumberOfAdapters_Get() (num int, err error) {
	var n int
	result, _, errno := syscall.Syscall(ADL_Adapter_NumberOfAdapters_Get, 1, uintptr(unsafe.Pointer(&n)), 0, 0)

	if result != ADL_OK {
		return 0, errors.New(errno.Error())
	}
	return int(n), nil
}

func (adl *Adl) ADL_Adapter_AdapterInfo_Get(numOfAdapters int) (result []AdapterInfo, err error) {
	adaptersOutputParam := [C.ADL_MAX_DISPLAY_NAME]C.AdapterInfo{}

	ok, _, errno := syscall.Syscall(ADL_Adapter_AdapterInfo_Get, 2, uintptr(unsafe.Pointer(&adaptersOutputParam[0])), unsafe.Sizeof(C.AdapterInfo{})*uintptr(numOfAdapters), 0)
	if ok != ADL_OK {
		return nil, errors.New(errno.Error())
	}

	for _, adapter := range adaptersOutputParam[:numOfAdapters] {
		result = append(result, AdapterInfo{
			ISize:            int(adapter.iSize),
			IAdapterIndex:    int(adapter.iAdapterIndex),
			StrUDID:          convertString(adapter.strUDID),
			IBusNumber:       int(adapter.iBusNumber),
			IDeviceNumber:    int(adapter.iDeviceNumber),
			IFunctionNumber:  int(adapter.iFunctionNumber),
			IVendorID:        int(adapter.iVendorID),
			StrAdapterName:   convertString(adapter.strAdapterName),
			StrDisplayName:   convertString(adapter.strDisplayName),
			IPresent:         int(adapter.iPresent),
			IExist:           int(adapter.iExist),
			StrDriverPath:    convertString(adapter.strDriverPath),
			StrDriverPathExt: convertString(adapter.strDriverPathExt),
			StrPNPString:     convertString(adapter.strPNPString),
			IOSDisplayIndex:  int(adapter.iOSDisplayIndex),
		})
	}
	return
}
