package utils

import (
	"errors"
	"strings"
)

// BankParser interface
type BankParser interface {
	Parse(sms string) (*ParsedSMS, error)
}

// GetBankParser detects the bank and message type and returns the correct parser
func GetBankParser(sms string) BankParser {
	smsLower := strings.ToLower(sms)

	// CBE sent (transfer)
	if strings.Contains(smsLower, "transfered") && strings.Contains(smsLower, "apps.cbe.com.et") {
		return &CBEDepositParser{}
	}

	// CBE received (credited)
	if strings.Contains(smsLower, "credited with") && strings.Contains(smsLower, "ref no") {
		return &CBEReceivedParser{}
	}

	// Telebirr deposit (sent)
	if strings.Contains(smsLower, "you have transferred") && strings.Contains(smsLower, "ethiotelecom") {
		return &TelebirrDepositParser{}
	}

	// Telebirr received
	if strings.Contains(smsLower, "you have received") && strings.Contains(smsLower, "telebirr") {
		return &TelebirrReceivedParser{}
	}

	// Unknown bank
	return nil
}

// ParseBankMulti parses any supported bank SMS and returns a ParsedSMS
func ParseBankMulti(sms string) (*ParsedSMS, error) {
	parser := GetBankParser(sms)
	if parser == nil {
		return nil, errors.New("unsupported bank or unrecognized SMS format")
	}

	return parser.Parse(sms)
}
