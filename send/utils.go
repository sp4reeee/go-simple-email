package send

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

func setHeaders(req *http.Request, canary string, mailbox string, auth string) {
	req.Header.Set("host", "outlook.live.com")
	req.Header.Set("content-length", "3816")
	req.Header.Set("prefer", "exchange.behavior=\"IncludeThirdPartyOnlineMeetingProviders\"")
	req.Header.Set("x-req-source", "Mail")
	req.Header.Set("x-owa-canary", canary)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("content-type", "application/json; charset=utf-8")
	req.Header.Set("action", "UpdateItem")
	req.Header.Set("x-owa-hosted-ux", "false")
	req.Header.Set("accept", "*/*")
	req.Header.Set("origin", "https://outlook.live.com")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("accept-language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("cookie", fmt.Sprintf("DefaultAnchorMailbox=%s; RPSSecAuth=%s;", mailbox, auth))
}

func generateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	b[8] = b[8]&^0xc0 | 0x80
	b[6] = b[6]&^0xf0 | 0x40

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}
