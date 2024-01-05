package send

import "fmt"

func Send(email string, password string, to string, subject string, data string) {

	cookies := get_cookies(email, password)

	resp := send_email(to, subject, data, cookies)

	fmt.Println(resp)
}
