package utils

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ParsedSMS is your existing struct
type ParsedSMS struct {
	Account       string
	Amount        float64
	SenderName    string
	Date          time.Time
	TransactionID string
	Balance       float64
	URL           string
}

// ParseBankSMS parses bank credit SMS
func ParseBankSMS(sms string) (*ParsedSMS, error) {
	parsed := &ParsedSMS{}

	// 1️⃣ Account
	accountRe := regexp.MustCompile(`Account (\d+\*+\d+)`)
	if m := accountRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.Account = m[1]
	} else {
		log.Println("Failed to parse Account")
		return nil, nil
	}

	// 2️⃣ Amount
	amountRe := regexp.MustCompile(`Credited with ETB ([\d,]+\.?\d*)`)
	if m := amountRe.FindStringSubmatch(sms); len(m) > 1 {
		amount, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Amount = amount
	} else {
		log.Println("Failed to parse Amount")
		return nil, nil
	}

	// 3️⃣ Sender Name
	senderRe := regexp.MustCompile(`from (.*?), on`)
	if m := senderRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.SenderName = strings.TrimSpace(m[1])
	} else {
		log.Println("Failed to parse Sender")
		return nil, nil
	}

	// 4️⃣ Date (with 'at' between date and time)
	dateRe := regexp.MustCompile(`on (\d{2}/\d{2}/\d{4}) at (\d{2}:\d{2}:\d{2})`)
	if m := dateRe.FindStringSubmatch(sms); len(m) > 2 {
		datetime := m[1] + " " + m[2]
		date, _ := time.Parse("02/01/2006 15:04:05", datetime)
		parsed.Date = date
	} else {
		log.Println("Failed to parse Date")
		return nil, nil
	}

	// 5️⃣ Transaction ID
	txRe := regexp.MustCompile(`Ref No (\w+)`)
	if m := txRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.TransactionID = m[1]
	} else {
		log.Println("Failed to parse Transaction ID")
		return nil, nil
	}

	// 6️⃣ Balance
	balanceRe := regexp.MustCompile(`Your Current Balance is ETB ([\d,]+\.?\d*)`)
	if m := balanceRe.FindStringSubmatch(sms); len(m) > 1 {
		balance, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Balance = balance
	} else {
		log.Println("Failed to parse Balance")
		return nil, nil
	}

	// 7️⃣ URL
	urlRe := regexp.MustCompile(`https?://\S+`)
	parsed.URL = urlRe.FindString(sms)

	return parsed, nil
}

// -------------------------------
// New parser for Verify SMS
// -------------------------------
func ParseVerifySMS(sms string) (*ParsedSMS, error) {
	parsed := &ParsedSMS{}

	// 1️⃣ Transaction ID
	txRe := regexp.MustCompile(`id=(FT\d+[A-Z0-9]+)`)
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

	// 3️⃣ Receiver Name
	nameRe := regexp.MustCompile(`to ([A-Za-z ]+?) on`)
	if m := nameRe.FindStringSubmatch(sms); len(m) > 1 {
		parsed.SenderName = strings.TrimSpace(m[1])
	} else {
		return nil, errors.New("Receiver name not found")
	}

	// 4️⃣ Amount (total of ETB)
	// Match "total of ETB 20", "total of ETB20.00", etc.
	amountRe := regexp.MustCompile(`total of ETB\s*([\d,]+(?:\.\d+)?)`)

	if m := amountRe.FindStringSubmatch(sms); len(m) > 1 {
		amt, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Amount = amt
	} else {
		return nil, errors.New("Amount not found")
	}

	// 5️⃣ Date
	dateRe := regexp.MustCompile(`on (\d{2}/\d{2}/\d{4} at \d{2}:\d{2}:\d{2})`)
	if m := dateRe.FindStringSubmatch(sms); len(m) > 1 {
		dt, _ := time.Parse("02/01/2006 15:04:05", m[1])
		parsed.Date = dt
	} else {
		return nil, errors.New("Date not found")
	}

	// 6️⃣ Balance
	balRe := regexp.MustCompile(`Current Balance is ETB ([\d,]+\.\d+)`)
	if m := balRe.FindStringSubmatch(sms); len(m) > 1 {
		bal, _ := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
		parsed.Balance = bal
	} else {
		return nil, errors.New("Balance not found")
	}

	// 7️⃣ URL
	urlRe := regexp.MustCompile(`https?://\S+`)
	parsed.URL = urlRe.FindString(sms)

	return parsed, nil
}
