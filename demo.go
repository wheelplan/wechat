package wechat

import (
	"time"
)

/*var client Client = Client{
	Corpid:    "wwd327d0ea50c0da**",
	Appsecret: "1kJw7t38aUxb3doG7IaLGXiX6DGobUllCLvoMf1eg**",
	Agentid:   "10000**",
	Touser:    "@all",
}*/

func (c Client) BuyDogeCard(executedQty, avgPrice, cumQuote string) (err error) {

	//executed := fmt.Sprintf("%v", executedQty)
	//avg := fmt.Sprintf("%v", avgPrice)
	//cum := fmt.Sprintf("%v", cumQuote)

	nowtime := time.Now().Format("2006-01-02 15:04:05 .000")

	card := Card{
		Title: "🦁 DOGECOIN",
		Description: "<div class=\\\"gray\\\">" + nowtime + "</div><br><div class=\\\"normal\\\">" +
			"EXTRA !  You bought " + executedQty + " Dogecoins, the average price is " + avgPrice +
			" , and the total consumption is 💲" + cumQuote +
			" 💵💵💵</div><br><div class=\\\"highlight\\\">                           ☀️ I hope everything goes well .</div>",
		Url: "https://rocc.top",
		//Btntxt: "more",
	}

	err = c.SendCardOld(card)

	return

}

func (c Client) SellDogeCard(executedQty, balance string) (err error) {

	//executed := fmt.Sprintf("%v", executedQty)
	//avg := fmt.Sprintf("%v", avgPrice)
	//cum := fmt.Sprintf("%v", cumQuote)

	nowtime := time.Now().Format("2006-01-02 15:04:05 .000")

	card := Card{
		Title: "🐮 DOGECOIN",
		Description: "<div class=\\\"gray\\\">" + nowtime + "</div><br><div class=\\\"normal\\\">" +
			"EXTRA !  You sold " + executedQty + " Dogecoins, the current contract account balance is 💲" + balance +
			" 💵💵💵</div><br><div class=\\\"highlight\\\">                           🐥 I hope everything goes well .</div>",
		Url: "https://rocc.top",
		//Btntxt: "more",
	}

	err = c.SendCardOld(card)

	return

}

func (c Client) PriceWarnCard(coin, price string) (err error) {

	nowtime := time.Now().Format("2006-01-02 15:04:05 .000")

	card := Card{
		Title: "🐚 " + coin + "-usdt",
		Description: "<div class=\\\"gray\\\">" + nowtime + "</div><br><div class=\\\"normal\\\">" +
			"WARN! The current price of " + coin + " is 💲" + price +
			"</div><br><div class=\\\"highlight\\\">                   ☀️ News provided by 🚀@Launch</div>",
		Url: "https://rocc.top",
		//Btntxt: "more",
	}

	err = c.SendCardOld(card)

	return

}

func (c Client) WarnText(context string) (err error) {

	err = c.SendTextOld(context)

	return

}
