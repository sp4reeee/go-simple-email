package send

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/sp4reeee/go-simple-email/constants"
)

func send_email(to string, subject string, data string, cookies constants.Cookies) string {

	req, err := http.NewRequest("POST", constants.BASE_URL_SEND, bytes.NewBuffer([]byte(fmt.Sprintf(constants.POST_DATA, to, to, subject, data))))
	if err != nil {
		return "Error: Email not sent"
	}

	setHeaders(req, cookies.Canary, cookies.Mailbox, cookies.Auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Error : Email not sent"
	}

	defer resp.Body.Close()

	return "Success : Email sent"

}

func get_cookies(email string, password string) constants.Cookies {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var cobranID, err = generateUUID()
	var cookies constants.Cookies

	var temp_cookies []*network.Cookie
	err = chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf(constants.BASE_URL_LOGIN, cobranID)),

		chromedp.WaitVisible(constants.SELECTOR_BUTTON, chromedp.ByQuery),

		chromedp.SendKeys(constants.SELECTOR_EMAIL, email, chromedp.ByQuery),
		chromedp.Click(constants.SELECTOR_BUTTON, chromedp.ByQuery),

		chromedp.WaitVisible(constants.SELECTOR_BUTTON, chromedp.ByQuery),

		chromedp.SendKeys(constants.SELECTOR_PASSWORD, password, chromedp.ByQuery),
		chromedp.Sleep(500*time.Millisecond),
		chromedp.Click(constants.SELECTOR_BUTTON, chromedp.ByQuery),

		chromedp.WaitVisible(constants.SELECTOR_REMEMBER_ME, chromedp.ByQuery),

		chromedp.Click(constants.SELECTOR_REMEMBER_ME, chromedp.ByQuery),
		chromedp.Sleep(1500*time.Millisecond),

		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			temp_cookies, err = network.GetCookies().Do(ctx)
			return err
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, cookie := range temp_cookies {
		if cookie.Name == "X-OWA-CANARY" {
			cookies.Canary = cookie.Value
		}
		if cookie.Name == "DefaultAnchorMailbox" {
			cookies.Mailbox = cookie.Value
		}
		if cookie.Name == "RPSSecAuth" {
			cookies.Auth = cookie.Value
		}
	}

	return cookies
}
