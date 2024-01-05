# go-simple-email

## Introduction 🌟
The go-simple-email library is a practical and efficient tool for sending emails using a Hotmail account 📬. This Go library is ideal for businesses and individual developers seeking a straightforward solution for email management 🚀.

## Installation

To install go-simple-email, use go get:

### ```go get -u github.com/sp4reeee/go-simple-email/send```

## Quick Start

Using go-simple-email is simple! Here's an example to get you started:

```
package main

import (
	"io"
	"os"
	"strings"

	"github.com/sp4reeee/go-simple-email/send"
)

func main() {
	file, err := os.Open("password-recovery.html")
	if err != nil {
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return
	}

	send.Send("exemple@hotmail.com", "password", "email@to.sent", "Example: Password Reset", strings.ReplaceAll(string(content), `"`, `\"`))
}
```
✨ In this example, the library is used to send a password recovery email 🛠️. The content of the email is read from a file, and any double quotes in the content are replaced with \" before sending the email 📤.

## Important Note

1 - Hotmail Account and Sender Name: When using go-simple-email with a Hotmail account, ensure that the first name and last name associated with the Hotmail account match the name of your enterprise 🏢. This is particularly important for maintaining professionalism and brand consistency in your email communications 📩.

2 - Special Character Handling: The library intelligently handles special characters like quotes ("). When sending emails, any quotes in the email content are automatically replaced with \" to ensure proper formatting and delivery of the email content 📝.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

The go-simple-email library simplifies the process of sending emails through a Hotmail account with its straightforward API and smart handling of special characters 🌐. It's a valuable tool for businesses and developers who need a reliable and easy-to-use email solution 🛠️.
