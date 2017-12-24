package main

import (
	// "JetSMSDemo/smsService"
	"fmt"

	smsService "github.com/NetYazilim/goSMSService"
)

// gowsdl -o JetSMS.go -p smsService  https://www.jetsms.net/ws/soapsms.asmx?WSDL

func main() {

	url := "https://www.jetsms.net/ws/soapsms.asmx"

	client := smsService.NewSMSServiceSoap(url, true, nil)

	req := &smsService.SendSMSSingle{User: "User", Password: "Password", Originator: "Originator", Receipents: "Receipents", Messages: "Message"}
	//	var resp *smsService.SendSMSSingleResponse
	//	var err error
	resp, err := client.SendSMSSingle(req)

	if err != nil {
		fmt.Println("\n\nHata: ", err)
	} else {

		fmt.Printf("Sonuç: %v\n", resp)
		//		fmt.Printf("ID       : %v\n", resp.SendSMSResult.ID)
		//		fmt.Printf("XMLName  : %v\n", resp.SendSMSResult.XMLName)
	}
}
