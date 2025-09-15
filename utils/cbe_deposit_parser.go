package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type CBEDepositParser struct{}

func (p *CBEDepositParser) Parse(sms string) (*ParsedSMS, error) {
	parsed := &ParsedSMS{}

	// Transaction ID (remove last 8 digits after id=)
	txRe := regexp.MustCompile(`id=(FT\w+)`)
	if m := txRe.FindStringSubmatch(sms); len(m) > 1 {
		fullID := m[1]
		if len(fullID) > 8 {
			parsed.TransactionID = fullID[:len(fullID)-8] // remove last 8 digits
		} else {
			parsed.TransactionID = fullID
		}
	} else {
		return nil, errors.New("Transaction ID not found")
	}

	// Account
	accRe := regexp.MustCompile(`account (\d+\*+\d+)`)
	if m := accRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.Account = m[1]
	} else {
		return nil, errors.New("Account not found")
	}

	// Receiver Name
	nameRe := regexp.MustCompile(`to ([A-Za-z ]+?) on`)
	if m := nameRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.SenderName = strings.TrimSpace(m[1])
	}

	// Amount
	amountRe := regexp.MustCompile(`total of ETB\s*([\d,]+(?:\.\d+)?)`)
	if m := amountRe.FindStringSubmatch(sms); len(m) > 1 {
		amt, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Amount = amt
	} else {
		return nil, errors.New("Amount not found")
	}

	// Date
	dateRe := regexp.MustCompile(`on (\d{2}/\d{2}/\d{4} at \d{2}:\d{2}:\d{2})`)
	if m := dateRe.FindStringSubmatch(sms); len(m) > 1 {
		dt, _ := time.Parse("02/01/2006 15:04:05", m[1])
		parsed.Date = dt
	} else {
		return nil, errors.New("Date not found")
	}

	// Balance
	balRe := regexp.MustCompile(`Current Balance is ETB ([\d,]+\.\d+)`)
	if m := balRe.FindStringSubmatch(sms); len(m) > 1 {
		bal, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Balance = bal
	}

	// URL
	urlRe := regexp.MustCompile(`https?://\S+`)
	parsed.URL = urlRe.FindString(sms)

	return parsed, nil
}
