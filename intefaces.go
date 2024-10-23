package fias

import "fmt"

func GetIF(typ InterfaceID) string {

	return fmt.Sprintf("%s%s", FID_InterfaceFamily, typ)
}

type InterfaceID string

const (
	IF_CallAccounting        InterfaceID = "CA" // Call Accounting
	IF_KeyServiceSystem      InterfaceID = "DL" // Key Services System (Door Locking)
	IF_EnergyManagement      InterfaceID = "EM" // Energy Management
	IF_Minibar               InterfaceID = "MB" // Minibar
	IF_PBXGateway            InterfaceID = "PB" // TMS / PBX Gateway
	IF_POS                   InterfaceID = "PO" // POS
	IF_ExtendedVideoService  InterfaceID = "VI" // Pay TV / Extended Video Services
	IF_VoiceMail             InterfaceID = "VM" // Voice Mail
	IF_Miscellaneous         InterfaceID = "MS" // Miscellaneous / Data Retrieval System
	IF_InroomInternetSystems InterfaceID = "WW" // In-Room Internet Systems
)
