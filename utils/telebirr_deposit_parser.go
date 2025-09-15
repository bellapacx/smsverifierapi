package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type TelebirrDepositParser struct{}

func (p *TelebirrDepositParser) Parse(sms string) (*ParsedSMS, error) {
	parsed := &ParsedSMS{}

	// Transaction ID
	txRe := regexp.MustCompile(`transaction number is (\w+)`)
	if m := txRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.TransactionID = m[1]
	} else {
		return nil, errors.New("Transaction ID not found")
	}

	// Account (optional)
	accRe := regexp.MustCompile(`\((\d+\*+\d+)\)`)
	if m := accRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.Account = m[1]
	}

	// Receiver Name
	nameRe := regexp.MustCompile(`to ([A-Za-z ]+) \(`)
	if m := nameRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.SenderName = strings.TrimSpace(m[1])
	}

	// Amount
	amountRe := regexp.MustCompile(`transferred ETB ([\d,]+\.\d+)`)
	if m := amountRe.FindStringSubmatch(sms); len(m) > 1 {
		amt, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Amount = amt
	} else {
		return nil, errors.New("Amount not found")
	}

	// Date
	dateRe := regexp.MustCompile(`on (\d{2}/\d{2}/\d{4} \d{2}:\d{2}:\d{2})`)
	if m := dateRe.FindStringSubmatch(sms); len(m) > 1 {
		dt, _ := time.Parse("02/01/2006 15:04:05", m[1])
		parsed.Date = dt
	}

	// Balance
	balRe := regexp.MustCompile(`balance is ETB ([\d,]+\.\d+)`)
	if m := balRe.FindStringSubmatch(sms); len(m) > 1 {
		bal, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Balance = bal
	}

	// URL
	urlRe := regexp.MustCompile(`https?://\S+`)
	parsed.URL = urlRe.FindString(sms)

	return parsed, nil
}
