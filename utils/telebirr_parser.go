package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// TelebirrParser implements BankParser
type TelebirrParser struct{}

func (p *TelebirrParser) Parse(sms string) (*ParsedSMS, error) {
	parsed := &ParsedSMS{}

	// Transaction ID
	txRe := regexp.MustCompile(`(?i)transaction number is\s+([A-Z0-9]+)`)
	if m := txRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.TransactionID = strings.TrimSpace(m[1])
	} else {
		return nil, errors.New("transaction ID not found")
	}

	// 2️⃣ Account
	accRe := regexp.MustCompile(`Account (\d+\*+\d+|\d+)`)
	if m := accRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.Account = m[1]
	}

	// 3️⃣ Sender / Receiver Name
	nameRe := regexp.MustCompile(`- ([A-Za-z ]+)\.`)
	if m := nameRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.SenderName = strings.TrimSpace(m[1])
	}

	// 4️⃣ Amount
	amountRe := regexp.MustCompile(`ETB ([\d,]+\.\d+)`)
	if m := amountRe.FindStringSubmatch(sms); len(m) > 1 {
		amt, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Amount = amt
	} else {
		return nil, errors.New("Amount not found")
	}

	// 5️⃣ Date
	dateRe := regexp.MustCompile(`on (\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`)
	if m := dateRe.FindStringSubmatch(sms); len(m) > 1 {
		dt, _ := time.Parse("2006-01-02 15:04:05", m[1])
		parsed.Date = dt
	} else {
		parsed.Date = time.Now()
	}

	// 6️⃣ Balance
	balRe := regexp.MustCompile(`balance is ETB ([\d,]+\.\d+)`)
	if m := balRe.FindStringSubmatch(sms); len(m) > 1 {
		bal, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Balance = bal
	}

	// 7️⃣ URL
	urlRe := regexp.MustCompile(`https?://\S+`)
	parsed.URL = urlRe.FindString(sms)

	return parsed, nil
}
