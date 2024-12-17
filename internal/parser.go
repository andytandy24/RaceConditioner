package internal

import (
	"log"
	"strings"
)

func ParseHeaders(headers string) map[string]string {
	var headerMap = make(map[string]string)

	if headers != "" {
		for _, header := range strings.Split(headers, ";") {
			parts := strings.SplitN(header, "=", 2)
			if len(parts) == 2 {
				headerMap[parts[0]] = parts[1]
			} else {
				log.Fatalf("Invalid header format: %s", header)
			}
		}
		return headerMap
	} else {
		return headerMap
	}
}
