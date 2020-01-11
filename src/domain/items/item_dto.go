package items

/*
一度book_storeとして実装し、後からvideo_storeに改修していく方針。
*/

// Item 動画データを想定
type Item struct {
	ID                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

// Description 説明
type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

// Picture 写真
type Picture struct {
	ID  int64  `json:"id"`
	Url string `json:"url"`
}
