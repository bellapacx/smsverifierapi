package main

import (
	"fmt"
	"sms-verifier/utils"
)

func main() {
	samples := []string{
		// Original format
		`Dear Lidiya your Account 1*****6708 has been Credited with ETB 7,000.00 from Fitsumberhan Amanuel, on 10/09/2025 at 10:30:30 with Ref No FT25253N10L8 Your Current Balance is ETB 8,970.46. Thank you for Banking with CBE! https://apps.cbe.com.et:100/?id=FT25253N10L815706708`,

		// Extra spaces
		`Dear Lidiya  your Account 1*****6708  has been Credited with ETB 7,000.00  from Fitsumberhan Amanuel,  on 10/09/2025 at 10:30:30 with Ref No FT25253N10L8  Your Current Balance is ETB 8,970.46.`,

		// No URL
		`Account 1*****6708 has been Credited with ETB 7,000.00 from Fitsumberhan Amanuel, on 10/09/2025 at 10:30:30 with Ref No FT25253N10L8 Your Current Balance is ETB 8,970.46.`,

		// No “Dear Lidiya”
		`Account 1*****6708 has been Credited with ETB 7,500.00 from John Doe, on 11/09/2025 at 12:15:00 with Ref No FT25253N11J1 Your Current Balance is ETB 9,500.00. https://bank.example.com/id=FT25253N11J1`,
	}

	for i, sms := range samples {
		fmt.Printf("\n=== Sample %d ===\n", i+1)
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
}
