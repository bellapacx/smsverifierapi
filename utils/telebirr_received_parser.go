package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type TelebirrReceivedParser struct{}

func (p *TelebirrReceivedParser) Parse(sms string) (*ParsedSMS, error) {
	parsed := &ParsedSMS{}

	// Transaction ID
	txRe := regexp.MustCompile(`(?i)transaction number is (\w+)`)
	if m := txRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.TransactionID = m[1]
	} else {
		return nil, errors.New("Transaction ID not found")
	}

	// Sender (Name + Phone inside parentheses)
	senderRe := regexp.MustCompile(`(?i)from ([A-Za-z ]+)\((\d+)\)`)
	if m := senderRe.FindStringSubmatch(sms); len(m) > 2 {
		parsed.SenderName = strings.TrimSpace(m[1])
		parsed.Account = m[2] // phone number in parentheses
	}

	// Amount received
	amountRe := regexp.MustCompile(`(?i)received ETB ([\d,]+\.\d+)`)
	if m := amountRe.FindStringSubmatch(sms); len(m) > 1 {
		amt, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Amount = amt
	} else {
		return nil, errors.New("Amount not found")
	}

	// Date (dd/mm/yyyy hh:mm:ss)
	dateRe := regexp.MustCompile(`(\d{2}/\d{2}/\d{4} \d{2}:\d{2}:\d{2})`)
	if m := dateRe.FindStringSubmatch(sms); len(m) > 1 {
		dt, err := time.Parse("02/01/2006 15:04:05", m[1])
		if err == nil {
			parsed.Date = dt
		}
	}

	// Balance
	balRe := regexp.MustCompile(`(?i)Account balance is ETB ([\d,]+\.\d+)`)
	if m := balRe.FindStringSubmatch(sms); len(m) > 1 {
		bal, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Balance = bal
	}

	return parsed, nil
}
