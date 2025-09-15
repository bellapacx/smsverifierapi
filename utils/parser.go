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
	re := regexp.MustCompile(`Account (\d+\*+\d+) .* Credited with ETB ([\d,\.]+) from (.+?), on (\d{2}/\d{2}/\d{4} \d{2}:\d{2}:\d{2}) with Ref No (\w+) Your Current Balance is ETB ([\d,\.]+)`)

	matches := re.FindStringSubmatch(sms)
	if len(matches) != 7 {
		log.Println("SMS parsing failed:", sms)
		return nil, nil
	}

	amount, _ := strconv.ParseFloat(strings.ReplaceAll(matches[2], ",", ""), 64)
	date, _ := time.Parse("02/01/2006 15:04:05", matches[4])
	balance, _ := strconv.ParseFloat(strings.ReplaceAll(matches[6], ",", ""), 64)

	return &ParsedSMS{
		Account:       matches[1],
		Amount:        amount,
		SenderName:    matches[3],
		Date:          date,
		TransactionID: matches[5],
		Balance:       balance,
	}, nil
}
