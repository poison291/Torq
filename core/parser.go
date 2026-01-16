package core

import (
	"errors"
	"strconv"
)

func Parser(data []byte, pos *int) (any, error) {

	if *pos >= len(data) {
		return nil, errors.New("Unexpected End of Data")
	}

	switch data[*pos] {
	case 'i':
		return parseInt(data, pos)
	case 'l':
		return parseList(data, pos)
	case 'd':
		return parseDict(data, pos)
	default:
		if data[*pos] >= '0' && data[*pos] <= '9' {
			return parseString(data, pos)
		}
	}
	return nil, errors.New("invalid bencode format")
}

// Integer Data parsing
func parseInt(data []byte, pos *int) (int64, error) {
	if data[*pos] != 'i' {
		return 0, errors.New("Expected 'i' ")
	}
	*pos++
	start := *pos

	for data[*pos] != 'e' {
		*pos++
		if *pos >= len(data) {
			return 0, errors.New("Unexpected End of data")
		}
	}
	n, err := strconv.ParseInt(string(data[start:*pos]), 10, 64)
	if err != nil {
		return 0, err
	}
	*pos++
	return n, nil
}

func parseString(data []byte, pos *int) (string, error) {
	start := *pos
	for *pos < len(data) && data[*pos] >= '0' && data[*pos] <= '9' {
		*pos++
	}
	if *pos >= len(data) || data[*pos] != ':' {
		return "", errors.New("invalid string: Missing':' after length")
	}
	length, err := strconv.Atoi(string(data[start:*pos]))
	if err != nil {
		return "", err
	}
	*pos++
	if *pos > len(data) {
		return "", errors.New("invalid string: not enough bytes")
	}

	str := string(data[*pos : *pos+length])
	*pos += length
	return str, nil
}

func parseList(data []byte, pos *int) ([]any, error) {
	if data[*pos] != 'l' {
		return nil, errors.New("Expected 'l' at list start")
	}
	*pos++

	var result []any
	for *pos < len(data) && data[*pos] != 'e' {
		item, err := Parser(data, pos)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	if *pos >= len(data) || data[*pos] != 'e' {
		return nil, errors.New("unexpected end of list")
	}
	*pos++
	return result, nil
}

func parseDict(data []byte, pos *int) (map[string]any, error) {
	if data[*pos] != 'd' {
		return nil, errors.New("expected 'd' at dict start")
	}
	*pos++

	result := make(map[string]any)

	for *pos < len(data) && data[*pos] != 'e' {
		key, err := parseString(data, pos)
		if err != nil {
			return nil, err
		}

		value, err := Parser(data, pos)
		if err != nil {
			return nil, err
		}

		result[key] = value
	}

	if *pos >= len(data) || data[*pos] != 'e' {
		return nil, errors.New("unexpected end of dict")
	}
	*pos++ 
	return result, nil
}
