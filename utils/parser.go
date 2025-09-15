package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ParsedSMS struct {
	Account       string
	Amount        float64
	SenderName    string
	Date          time.Time
	TransactionID string
	Balance       float64
}

func ParseBankSMS(sms string) (*ParsedSMS, error) {
	// Regex: capture account, amount, sender, date, transaction ID, balance
	re := regexp.MustCompile(`Account (\d+\*+\d+) has been Credited with ETB ([\d,\.]+) from (.*?), on (\d{2}/\d{2}/\d{4} \d{2}:\d{2}:\d{2}) with Ref No (\w+) Your Current Balance is ETB ([\d,\.]+)`)

	matches := re.FindStringSubmatch(sms)
	if len(matches) != 7 {
		log.Println("SMS parsing failed:", sms)
		return nil, nil
	}

	amount, err := strconv.ParseFloat(strings.ReplaceAll(matches[2], ",", ""), 64)
	if err != nil {
		log.Println("Failed to parse amount:", matches[2])
		return nil, err
	}

	date, err := time.Parse("02/01/2006 15:04:05", matches[4])
	if err != nil {
		log.Println("Failed to parse date:", matches[4])
		return nil, err
	}

	balance, err := strconv.ParseFloat(strings.ReplaceAll(matches[6], ",", ""), 64)
	if err != nil {
		log.Println("Failed to parse balance:", matches[6])
		return nil, err
	}

	return &ParsedSMS{
		Account:       matches[1],
		Amount:        amount,
		SenderName:    matches[3],
		Date:          date,
		TransactionID: matches[5],
		Balance:       balance,
	}, nil
}
