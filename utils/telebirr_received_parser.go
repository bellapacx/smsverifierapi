package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// TelebirrReceivedParser handles received money SMS for Telebirr
type TelebirrReceivedParser struct{}

func (p *TelebirrReceivedParser) Parse(sms string) (*ParsedSMS, error) {
	parsed := &ParsedSMS{}

	// Transaction ID
	txRe := regexp.MustCompile(`(?i)transaction number is\s+([A-Z0-9]+)`)
	if m := txRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.TransactionID = strings.TrimSpace(m[1])
	} else {
		return nil, errors.New("transaction ID not found")
	}

	// 2️⃣ Account (inside parentheses after sender)
	accRe := regexp.MustCompile(`\((\d+\*+\d+)\)`)
	if m := accRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.Account = m[1]
	} else {
		return nil, errors.New("Account not found")
	}

	// 3️⃣ Sender Name
	senderRe := regexp.MustCompile(`from ([A-Za-z ]+)\(`)
	if m := senderRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.SenderName = strings.TrimSpace(m[1])
	}

	// 4️⃣ Amount
	amountRe := regexp.MustCompile(`received ETB ([\d,]+\.\d+)`)
	if m := amountRe.FindStringSubmatch(sms); len(m) > 1 {
		amt, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Amount = amt
	} else {
		return nil, errors.New("Amount not found")
	}

	// 5️⃣ Date (DD/MM/YYYY HH:MM:SS)
	dateRe := regexp.MustCompile(`on (\d{2}/\d{2}/\d{4} \d{2}:\d{2}:\d{2})`)
	if m := dateRe.FindStringSubmatch(sms); len(m) > 1 {
		dt, _ := time.Parse("02/01/2006 15:04:05", m[1])
		parsed.Date = dt
	} else {
		parsed.Date = time.Now() // fallback
	}

	// 6️⃣ Balance
	balRe := regexp.MustCompile(`balance (?:is )?ETB ([\d,]+\.\d+)`)
	if m := balRe.FindStringSubmatch(sms); len(m) > 1 {
		bal, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Balance = bal
	}

	return parsed, nil
}
