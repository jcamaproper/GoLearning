package main

import (
	"fmt"
	"time"
	"utils/chrome"

	ch "github.com/chromedp/chromedp"
)

// Selectos to perform clicks

const (
	//login elements
	loginButton  = "#__next > div > div.main-content > div.bywovg-0.kuGegY > div:nth-child(1) > div > div.sc-111rrsy-0.qbrWo > button.x0o17e-0.qrNYy"
	loginWindow  = "body > div.t8xka3-0.clxZon.modalOpened > div > div"
	emailField   = "body > div.t8xka3-0.clxZon.modalOpened > div > div > div > div:nth-child(3) > input"
	passField    = "body > div.t8xka3-0.clxZon.modalOpened > div > div > div > div.sc-1htht4q-3.kHSKLo.last > div.sc-1srpo52-0.ciHfxX > input"
	loginButtonW = "body > div.t8xka3-0.clxZon.modalOpened > div > div > div > div.sc-1htht4q-6.kUXNCx > button"

	//Home Elements
	homePage      = "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div"
	btcPriceLabel = "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div > div.h7vnx2-1.bFzXgL > table > tbody > tr:nth-child(1) > td:nth-child(4) > div > a > span"

	// credentials
	email = "jcamacholpz@gmail.com"
	pass  = "Juan1996"
)

func main() {

	// Create a context
	ctx, cancel, err := chrome.StartChrome(false)
	if err != nil {
		println(err.Error())
		return
	}
	defer cancel()
	fmt.Println(ctx)

	//Oper URL

	url := "https://coinmarketcap.com/"
	if err := ch.Run(
		ctx,
		ch.Navigate(url)); err != nil {
		println(err.Error())
		return
	}

	// click on login button

	if err := ch.Run(ctx,
		// Wait ready, until find the element on page
		// inputs #selector and function ByQuery, this can change to Path, etc
		ch.WaitReady(loginButton, ch.ByQuery),
		ch.Click(loginButton, ch.ByQuery),
	); err != nil {
		println(err.Error())
		return
	}

	// enter credentials

	if err := ch.Run(
		ctx,
		// Wait ready, until find the element on page
		ch.WaitReady(loginWindow, ch.ByQuery),
		ch.Click(emailField, ch.ByQuery),
		//ch.SetValue(emailField, email, ch.ByQuery),
		ch.Sleep(1*time.Second),
		ch.KeyEvent(email),
		ch.Click(passField, ch.ByQuery),
		//ch.SetValue(passField, pass, ch.ByQuery),
		ch.Sleep(1*time.Second),
		ch.KeyEvent(pass),
		ch.Click(loginButtonW, ch.ByQuery),
	); err != nil {
		println(err.Error())
		return
	}

	//pause to pass CAPTCHA manually
	time.Sleep(15 * time.Second)

	var btcPrice string
	// get BTC current price
	if err := ch.Run(ctx,
		//ch.WaitReady(homePage, ch.ByQuery),
		ch.Value(btcPriceLabel, &btcPrice, ch.ByQuery),
	); err != nil {
		println(err.Error())
		return
	}

	fmt.Println("Bitcoin price is: ", btcPrice)

	//time.Sleep(1 * time.Second)
}
