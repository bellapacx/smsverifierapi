package main

import (
	"fmt"
	"sms-verifier/utils"
)

func main() {
	tests := []struct {
		name   string
		sms    string
		parser interface {
			Parse(string) (*utils.ParsedSMS, error)
		}
	}{
		{
			name:   "CBE sent",
			sms:    `Dear Fitsumberhan, You have transfered ETB 7,000.00 to Lidiya Mulugeta on 10/09/2025 at 10:30:30 from your account 1*****5934. Your account has been debited with a S.charge of ETB 0 and 15% VAT of ETB0.00, with a total of ETB7000. Your Current Balance is ETB 1,237.60. Thank you for Banking with CBE! https://apps.cbe.com.et:100/?id=FT25253N10L850735934 For feedback click the link https://forms.gle/R1s9nkJ6qZVCxRVu9`,
			parser: &utils.CBEDepositParser{},
		},
		{
			name:   "CBE credited",
			sms:    `Dear Lidiya your Account 1*****6708 has been Credited with ETB 7,000.00 from Fitsumberhan Amanuel, on 10/09/2025 at 10:30:30 with Ref No FT25253N10L8 Your Current Balance is ETB 8,970.46. Thank you for Banking with CBE! https://apps.cbe.com.et:100/?id=FT25253N10L815706708`,
			parser: &utils.CBEReceivedParser{},
		},
		{
			name: "Telebirr deposit",
			sms: `Dear Lidia 
You have transferred ETB 220.00 to Mekdes Addisu (2519****0690) on 14/09/2025 11:36:27. Your transaction number is CIE8TH4R6K. The service fee is ETB 1.74 and 15% VAT on the service fee is ETB 0.26. Your current E-Money Account balance is ETB 1,335.00. To download your payment information please click this link: https://transactioninfo.ethiotelecom.et/receipt/CIE8TH4R6K.

Thank you for using telebirr
Ethio telecom`,
			parser: &utils.TelebirrDepositParser{},
		},
		{
			name: "Telebirr received",
			sms: `Dear Lidia,
You have received ETB 850.00 by transaction number CIE8TGIIDE on 2025-09-14 11:03:42 from Commercial Bank of Ethiopia to your telebirr Account 251914555740 - Lidia Mulugeta Gebrekiros. Your current balance is ETB 850.00.
Thank you for using telebirr
Ethio telecom`,
			parser: &utils.TelebirrReceivedParser{},
		},
	}

	for _, test := range tests {
		fmt.Printf("\n=== %s ===\n", test.name)
		parsed, err := test.parser.Parse(test.sms)
		if err != nil {
			fmt.Println("Error parsing SMS:", err)
			continue
		}
		fmt.Printf("Parsed SMS:\n%+v\n", parsed)
	}
}
