package main


import "github.com/clear-street/clst"

func main() {
    api := clst.NewAPI("TEST")
    api.Insert({
      "type": "bilateral_trade",
      "timestamp": 1556544618,
      "client_trade_id": "042919-1",
      "date": 20190304,
      "account_id": 100016,
      "mic": "XNAS",
      "contra_mpid": "AABB",
      "contra_dtc_num": "0123",
      "contra_side_qualifier": "close",
      "exec_mpid": "CSMM",
      "capacity": "principal",
      "quantity": "100",
      "price": "140.00",
      "instrument": {
        "identifier": "ATRA",
        "identifier_type": "ticker",
        "currency": "USD",
        "country": "USA"
      },
      "side": {
        "direction": "buy"
      }
    })
}