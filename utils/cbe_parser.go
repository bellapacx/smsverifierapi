package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// CBEParser implements BankParser
type CBEParser struct{}

func (p *CBEParser) Parse(sms string) (*ParsedSMS, error) {
	parsed := &ParsedSMS{}

	// 1️⃣ Transaction ID: remove last 8 digits after the "id="
	txRe := regexp.MustCompile(`id=(FT\d+[A-Z]{2}\d)`)
	if m := txRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.TransactionID = m[1]
	} else {
		return nil, errors.New("Transaction ID not found")
	}

	// 2️⃣ Account
	accRe := regexp.MustCompile(`account (\d+\*+\d+)`)
	if m := accRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.Account = m[1]
	} else {
		return nil, errors.New("Account not found")
	}

	// 3️⃣ Receiver / Sender Name
	nameRe := regexp.MustCompile(`to ([A-Za-z ]+?) on`)
	if m := nameRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.SenderName = strings.TrimSpace(m[1])
	} else {
		// fallback for received
		senderRe := regexp.MustCompile(`from ([A-Za-z ]+),`)
		if m := senderRe.FindStringSubmatch(sms); len(m) > 1 {
			parsed.SenderName = strings.TrimSpace(m[1])
		}
	}

	// 4️⃣ Amount (total of ETB)
	amountRe := regexp.MustCompile(`total of ETB([\d,]+\.\d+)`)
	if m := amountRe.FindStringSubmatch(sms); len(m) > 1 {
		amt, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Amount = amt
	} else {
		// fallback for received
		amountRe := regexp.MustCompile(`ETB ([\d,]+\.\d+)`)
		if m := amountRe.FindStringSubmatch(sms); len(m) > 1 {
			amt, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
			parsed.Amount = amt
		} else {
			return nil, errors.New("Amount not found")
		}
	}

	// 5️⃣ Date
	dateRe := regexp.MustCompile(`on (\d{2}/\d{2}/\d{4} at \d{2}:\d{2}:\d{2})`)
	if m := dateRe.FindStringSubmatch(sms); len(m) > 1 {
		dt, _ := time.Parse("02/01/2006 15:04:05", m[1])
		parsed.Date = dt
	} else {
		// fallback for received SMS (yyyy-mm-dd hh:mm:ss)
		dateRe := regexp.MustCompile(`on (\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`)
		if m := dateRe.FindStringSubmatch(sms); len(m) > 1 {
			dt, _ := time.Parse("2006-01-02 15:04:05", m[1])
			parsed.Date = dt
		} else {
			parsed.Date = time.Now() // fallback to now
		}
	}

	// 6️⃣ Balance
	balRe := regexp.MustCompile(`Current Balance is ETB ([\d,]+\.\d+)`)
	if m := balRe.FindStringSubmatch(sms); len(m) > 1 {
		bal, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Balance = bal
	}

	// 7️⃣ URL
	urlRe := regexp.MustCompile(`https?://\S+`)
	parsed.URL = urlRe.FindString(sms)

	return parsed, nil
}
