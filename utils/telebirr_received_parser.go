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

	// Transaction ID (case-insensitive)
	txRe := regexp.MustCompile(`(?i)transaction number (\w+)`)
	if m := txRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.TransactionID = m[1]
	} else {
		return nil, errors.New("Transaction ID not found")
	}

	// Account (case-insensitive)
	accRe := regexp.MustCompile(`(?i)telebirr Account (\d+)`)
	if m := accRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.Account = m[1]
	} else {
		return nil, errors.New("Account not found")
	}

	// Sender Name (relaxed)
	senderRe := regexp.MustCompile(`(?i)from ([A-Za-z ]+?) to your`)
	if m := senderRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.SenderName = strings.TrimSpace(m[1])
	}

	// Amount (case-insensitive)
	amountRe := regexp.MustCompile(`(?i)received ETB ([\d,]+\.\d+)`)
	if m := amountRe.FindStringSubmatch(sms); len(m) > 1 {
		amt, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Amount = amt
	} else {
		return nil, errors.New("Amount not found")
	}

	// Date (yyyy-mm-dd hh:mm:ss)
	dateRe := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`)
	if m := dateRe.FindStringSubmatch(sms); len(m) > 1 {
		dt, _ := time.Parse("2006-01-02 15:04:05", m[1])
		parsed.Date = dt
	}

	// Balance (case-insensitive)
	balRe := regexp.MustCompile(`(?i)current balance is ETB ([\d,]+\.\d+)`)
	if m := balRe.FindStringSubmatch(sms); len(m) > 1 {
		bal, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Balance = bal
	}

	return parsed, nil
}
