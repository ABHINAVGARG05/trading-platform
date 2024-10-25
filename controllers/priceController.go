package controllers

import (
	"time"
	"math/rand"

	"github.com/ABHINAVGARG05/trading-platform/database"
	"github.com/ABHINAVGARG05/trading-platform/models"
);

func PriceUpdater(){
	var stocks[] models.Stock
	database.DB.Find(&stocks)
	for _,stock:= range stocks{
		if stock.Demand>stock.Supply{
			dmr := float64(stock.Demand)/float64(stock.Supply+1)
			stock.Price += stock.Price*0.01*dmr
		}else if stock.Supply>stock.Demand{
			dmr := float64(stock.Supply)/float64(stock.Demand+1)
			stock.Price -= stock.Price*0.01*dmr
		}
		database.DB.Save(&stock)
	}
}

func RandomPriceUpdater(stock *models.Stock){
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(2) + 1

	switch randomNum {
	case 1:
		dmr := float64(stock.Demand)/float64(stock.Supply+1)
		stock.Price += stock.Price*0.01*dmr;
	case 2:
		dmr := float64(stock.Supply)/float64(stock.Demand+1)
		stock.Price -= stock.Price*0.01*dmr
	}
	database.DB.Save(&stock)
}


func PriceTicker(){
	ticker := time.NewTicker(1*time.Minute)
	go func(){
		for range ticker.C{
			var stocks[] models.Stock
			database.DB.Find(&stocks)
			for _, stock := range stocks {
                RandomPriceUpdater(&stock)
            }
		}
	}()
}