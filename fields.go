package fias

import (
	"fmt"
	"time"
)

type FieldID string

const (
	FID_Track1                      FieldID = "$1"
	FID_Track2                      FieldID = "$2"
	FID_Track3                      FieldID = "$3"
	FID_VendorMsgID                 FieldID = "$J"
	FID_UserDefined0                FieldID = "A0"
	FID_UserDefined1                FieldID = "A1"
	FID_UserDefined2                FieldID = "A2"
	FID_UserDefined3                FieldID = "A3"
	FID_UserDefined4                FieldID = "A4"
	FID_UserDefined5                FieldID = "A5"
	FID_UserDefined6                FieldID = "A6"
	FID_UserDefined7                FieldID = "A7"
	FID_UserDefined8                FieldID = "A8"
	FID_UserDefined9                FieldID = "A9"
	FID_AnswerStatus                FieldID = "AS"
	FID_BalanceAmount               FieldID = "BA"
	FID_ItemDescription             FieldID = "BD"
	FID_ItemAmount                  FieldID = "BI"
	FID_POSCheckNumber              FieldID = "C#"
	FID_Cryptogram                  FieldID = "CG"
	FID_CreditLimit                 FieldID = "CL"
	FID_CreditLimitOverrideFlag     FieldID = "CO"
	FID_ClassOfService              FieldID = "CS"
	FID_ClearText                   FieldID = "CT"
	FID_NumberOfCovers              FieldID = "CV"
	FID_Discount1                   FieldID = "D1"
	FID_Discount2                   FieldID = "D2"
	FID_Discount3                   FieldID = "D3"
	FID_Discount4                   FieldID = "D4"
	FID_Discount5                   FieldID = "D5"
	FID_Discount6                   FieldID = "D6"
	FID_Discount7                   FieldID = "D7"
	FID_Discount8                   FieldID = "D8"
	FID_Discount9                   FieldID = "D9"
	FID_Date                        FieldID = "DA"
	FID_DepartmentCode              FieldID = "DC"
	FID_DialedDigits                FieldID = "DD"
	FID_DoNotDisturbStatus          FieldID = "DN"
	FID_DepartureTime               FieldID = "DT"
	FID_Duration                    FieldID = "DU"
	FID_EquipmentNumber             FieldID = "EN"
	FID_EquipmentPoolID             FieldID = "EP"
	FID_EquipmentStatus             FieldID = "ES"
	FID_EquipmentStatusOfSourceRoom FieldID = "ET"
	FID_FolioNumber                 FieldID = "F#"
	FID_ItemDisplayFlag             FieldID = "FD"
	FID_FieldList                   FieldID = "FL"
	FID_ReservationNumber           FieldID = "G#"
	FID_ProfileNumber               FieldID = "G+"
	FID_GuestArrivalDate            FieldID = "GA"
	FID_GuestDepartureDate          FieldID = "GD"
	FID_GuestFirstName              FieldID = "GF"
	FID_GuestGroupNumber            FieldID = "GG"
	FID_GuestLanguage               FieldID = "GL"
	FID_GuestName                   FieldID = "GN"
	FID_GuestPIN                    FieldID = "GP"
	FID_ShareFlag                   FieldID = "GS"
	FID_GuestTitle                  FieldID = "GT"
	FID_GuestVIPStatus              FieldID = "GV"
	FID_UserID                      FieldID = "ID"
	FID_InterfaceFamily             FieldID = "IF"
	FID_KeyCount                    FieldID = "K#"
	FID_KeyCoder                    FieldID = "KC"
	FID_KeyOptions                  FieldID = "KO"
	FID_KeyType                     FieldID = "KT"
	FID_LocatorExpiryTimer          FieldID = "LT"
	FID_NumberOfArticles            FieldID = "M#"
	FID_MinibarArticle              FieldID = "MA"
	FID_MaximumGuestMatch           FieldID = "MX"
	FID_MessageID                   FieldID = "MI"
	FID_MessageLightStatus          FieldID = "ML"
	FID_MeterPulse                  FieldID = "MP"
	FID_MinibarRights               FieldID = "MR"
	FID_MessageText                 FieldID = "MT"
	FID_NoPostFlag                  FieldID = "NP"
	FID_PostingSequenceNumber       FieldID = "P#"
	FID_PostingCallType             FieldID = "PC"
	FID_HotelID                     FieldID = "PH"
	FID_InquiryData                 FieldID = "PI"
	FID_PaymentMethod               FieldID = "PM"
	FID_PrinterPort                 FieldID = "PP"
	FID_PostingType                 FieldID = "PT"
	FID_PostingRoute                FieldID = "PX"
	FID_NumberOfPersons             FieldID = "PU"
	FID_RecordID                    FieldID = "RI"
	FID_MaximumMessageRecordLength  FieldID = "RL"
	FID_RoomNumber                  FieldID = "RN"
	FID_OldRoomNumber               FieldID = "RO"
	FID_RoompaymentMethod           FieldID = "RP"
	FID_RoomMaidStatus              FieldID = "RS"
	FID_RequestType                 FieldID = "RT"
	FID_Subtotal1                   FieldID = "S1"
	FID_Subtotal2                   FieldID = "S2"
	FID_Subtotal3                   FieldID = "S3"
	FID_Subtotal4                   FieldID = "S4"
	FID_Subtotal5                   FieldID = "S5"
	FID_Subtotal6                   FieldID = "S6"
	FID_Subtotal7                   FieldID = "S7"
	FID_Subtotal8                   FieldID = "S8"
	FID_Subtotal9                   FieldID = "S9"
	FID_ServiceCharge               FieldID = "SC"
	FID_SwapFlag                    FieldID = "SF"
	FID_SuiteInfo                   FieldID = "SI"
	FID_SalesOutlet                 FieldID = "SO"
	FID_ServingTime                 FieldID = "ST"
	FID_TableNumber                 FieldID = "T#"
	FID_Tax1                        FieldID = "T1"
	FID_Tax2                        FieldID = "T2"
	FID_Tax3                        FieldID = "T3"
	FID_Tax4                        FieldID = "T4"
	FID_Tax5                        FieldID = "T5"
	FID_Tax6                        FieldID = "T6"
	FID_Tax7                        FieldID = "T7"
	FID_Tax8                        FieldID = "T8"
	FID_Tax9                        FieldID = "T9"
	FID_TotalPostingAmount          FieldID = "TA"
	FID_Time                        FieldID = "TI"
	FID_Tip                         FieldID = "TP"
	FID_TVRights                    FieldID = "TV"
	FID_SuiteInfoForOldRoom         FieldID = "UO"
	FID_VendorVersion               FieldID = "V#"
	FID_VoiceMail                   FieldID = "VM"
	FID_VideoRights                 FieldID = "VR"
	FID_WorkstationID               FieldID = "WS"
	FID_CrossReferenceData          FieldID = "X1"
)

type Field interface {
	Format() string
}

type AlphaNumField struct {
	ID    string
	Value string
}

func (f AlphaNumField) Format() string {
	return fmt.Sprintf("%s%s", f.ID, f.Value)
}

type NumberField struct {
	ID    string
	Value int
}

func (f NumberField) Format() string {
	return fmt.Sprintf("%s%d", f.ID, f.Value)
}

type MonetaryField struct {
	ID    string
	Value float64
}

func (f MonetaryField) Format() string {
	return fmt.Sprintf("%s%.2f", f.ID, f.Value)
}

type DateField struct {
	ID    string
	Value time.Time
}

func (f DateField) Format() string {
	return fmt.Sprintf("%s%s", f.ID, f.Value.Format(FORMAT_DATE))
}

type TimeField struct {
	ID    string
	Value time.Time
}

func (f TimeField) Format() string {
	return fmt.Sprintf("%s%s", f.ID, f.Value.Format(FORMAT_TIME))
}
