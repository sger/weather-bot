# Weather bot

First create a new bot in Slack more info here [https://api.slack.com/bot-users](https://api.slack.com/bot-users)

Usage:

```
package main

import (
	"fmt"
	wb "github.com/sger/weather-bot"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: weather-bot token\n")
		os.Exit(1)
	}

	ws, r, err := wb.Connect(os.Args[1])
	fmt.Println("bot is running, hit ^C to exit")

	if err != nil {
		log.Fatal(err)
	}

	for {
		m, err := wb.GetMessage(ws)
		if err != nil {
			log.Fatal(err)
		}

		if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+r.Self.Id+">") {
			fields := strings.Fields(m.Text)
			if len(fields) == 2 {
				go func(m wb.Message) {
					m.Text = getWeather(fields[1])
					m.Text += "\nThank you " + "<@" + m.User + ">" + "\n"
					wb.PostMessage(ws, m)
				}(m)
			} else {
				go func(m wb.Message) {
					m.Text = "sorry wrong usage please write @name-of-the-weather-bot: Athens"
					wb.PostMessage(ws, m)
				}(m)
			}
		}
	}
}

func getWeather(city string) string {
	f, err := wb.Search(city)

	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Temperature for today\n Kelvin: %.2fK | Celsius: %.2f°C | Fahrenheit: %.2f°F", f.Info.Temp, f.Info.Temp-273.15, f.Info.Temp*9/5-459.67)
}
```


