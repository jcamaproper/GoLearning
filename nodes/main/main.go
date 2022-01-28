package main

import (
	"context"
	"fmt"
	"time"
	"utils/chrome"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	ch "github.com/chromedp/chromedp"
)

const (
	tableSelector = "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div > div.h7vnx2-1.bFzXgL > table"
	tbodySelector = "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div > div.h7vnx2-1.bFzXgL > table > tbody"
	trSelector    = "#__next > div > div.main-content > div.sc-57oli2-0.comDeo.cmc-body-wrapper > div > div > div.h7vnx2-1.bFzXgL > table > tbody > tr"
)

type cryptoCurreny struct {
	currencyName      string
	currentPrice      string
	marketCap         string
	circulatingSupply string
}

func main() {

	//start := time.Now()
	// Create a context
	ctx, cancel, err := chrome.StartChrome(false)
	if err != nil {
		println(err.Error())
		return
	}
	defer cancel()
	//fmt.Println(ctx)

	//Oper URL

	url := "https://coinmarketcap.com/"
	if err := ch.Run(ctx,
		ch.Navigate(url)); err != nil {
		println(err.Error())
		return
	}

	//scroll to view all the children elements

	//read nodes
	var rows []*cdp.Node
	if err := ch.Run(ctx,
		ch.WaitReady(tbodySelector, ch.ByQuery),
		ch.Nodes(tbodySelector, &rows, ch.ByQuery),
	); err != nil {
		println(err.Error())
		return
	}

	//Get all the child nodes for TBODY
	if err := ch.Run(ctx,
		ch.ActionFunc(func(c context.Context) error {
			// Get the children of the table
			return dom.RequestChildNodes(rows[0].NodeID).WithDepth(-1).Do(c)
		}),
		ch.Sleep(1*time.Second),
	); err != nil {
		//return err
		println(err.Error())
		return
	}

	//Iterate tbody
	for _, tr := range rows[0].Children {

		//td ->         div      ->   a      -> span
		priceElem := tr.Children[3].Children[0].Children[0].Children[0]
		priceNodeID := []cdp.NodeID{priceElem.NodeID}
		var price string

		if err := ch.Run(ctx,
			ch.Text(priceNodeID, &price, ch.ByNodeID),
		); err != nil {
			//return err
			println(err.Error())
		}
		fmt.Println(price)

	}

	/* for i, row := range rows {
		fmt.Println(i, "value", row.NodeValue, "type", row.NodeType.String(), "childcount", row.ChildNodeCount)
	} */

	//fmt.Println("Row found: ", len(rows))

	//fmt.Printf("\nTook: %f secs\n", time.Since(start).Seconds())

}
