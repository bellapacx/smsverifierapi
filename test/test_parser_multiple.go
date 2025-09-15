package main

import (
	"fmt"
	"sms-verifier/utils"
)

func main() {
	// ----- Test Bank SMS -----
	bankSamples := []string{
		// Original format
		`Dear Lidiya your Account 1*****6708 has been Credited with ETB 7,000.00 from Fitsumberhan Amanuel, on 10/09/2025 at 10:30:30 with Ref No FT25253N10L8 Your Current Balance is ETB 8,970.46. Thank you for Banking with CBE! https://apps.cbe.com.et:100/?id=FT25253N10L815706708`,

		// Extra spaces
		`Dear Lidiya  your Account 1*****6708  has been Credited with ETB 7,000.00  from Fitsumberhan Amanuel,  on 10/09/2025 at 10:30:30 with Ref No FT25253N10L8  Your Current Balance is ETB 8,970.46.`,
	}

	fmt.Println("===== Bank SMS Parser Test =====")
	for i, sms := range bankSamples {
		fmt.Printf("\n--- Sample %d ---\n", i+1)
		parsed, err := utils.ParseBankSMS(sms)
		if err != nil {
			fmt.Println("Error parsing SMS:", err)
			continue
		}
		if parsed == nil {
			fmt.Println("Parsing failed")
			continue
		}
		fmt.Printf("Parsed SMS:\n%+v\n", parsed)
	}

	// ----- Test Verify SMS -----
	verifySamples := []string{
		`Dear Lidiya, You have transfered ETB 20.00 to Yaekob Dinku on 14/09/2025 at 11:52:20 from your account 1*****6708. Your account has been debited with a S.charge of ETB 0 and  15% VAT of ETB0.00, with a total of ETB20. Your Current Balance is ETB 340.46. Thank you for Banking with CBE! https://apps.cbe.com.et:100/?id=FT25258GL7DZ15706708 For feedback click the link https://forms.gle/R1s9nkJ6qZVCxRVu9`,
	}

	fmt.Println("\n===== Verify SMS Parser Test =====")
	for i, sms := range verifySamples {
		fmt.Printf("\n--- Sample %d ---\n", i+1)
		parsed, err := utils.ParseVerifySMS(sms)
		if err != nil {
			fmt.Println("Error parsing SMS:", err)
			continue
		}
		fmt.Printf("Parsed SMS:\n%+v\n", parsed)
	}
}
