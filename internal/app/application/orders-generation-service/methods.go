package orders_generation_service

import (
	"fmt"
	"github.com/mkorobovv/my-rest/internal/app/domain/order"
	"math/rand"
	"strings"
	"time"
)

const (
	charset           = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	trackNumberLength = 20
	txLength          = 18
)

func (svc *OrdersGenerationService) Generate() order.Order {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	trackNumber := newTrackNumber(seed)
	locale := newLocale(seed)
	customerId := seed.Int63()
	phoneNumber := newPhoneNumber(seed)
	email := newEmail(seed)
	zipCode := newZipCode(seed)
	address := newAddress(seed, zipCode)
	recipientName := newRecipientName(seed)
	transactionID := newTransactionID(seed)
	currency := newCurrency(seed)
	amount := float64n(1000001)
	deliveryCost := float64n(10001)
	goodsTotal := amount + deliveryCost
	bank, provider := newBankInfo(seed)
	paymentDt := randomTime(seed)

	return order.Order{
		TrackNumber: trackNumber,
		Locale:      locale,
		CustomerID:  customerId,
		Delivery: order.Delivery{
			RecipientName: recipientName,
			PhoneNumber:   phoneNumber,
			Email:         &email,
			ZipCode:       zipCode,
			Address:       address,
		},
		Payment: order.Payment{
			TransactionID: transactionID,
			Currency:      currency,
			Amount:        amount,
			Provider:      provider,
			PaymentDt:     paymentDt,
			DeliveryCost:  deliveryCost,
			GoodsTotal:    goodsTotal,
			Bank:          bank,
		},
		Items: newItems(seed),
	}
}

func newTrackNumber(seed *rand.Rand) string {
	trackNumber := make([]byte, trackNumberLength)
	for i := range trackNumber {
		trackNumber[i] = charset[seed.Intn(len(charset))]
	}

	return string(trackNumber)
}

func newLocale(seed *rand.Rand) string {
	l := locales()

	return l[seed.Intn(len(l))]
}

func locales() []string {
	return []string{
		"Europe/Moscow",
		"Europe/Netherlands",
		"Europe/Germany",
		"Europe/Armenia",
		"Europe/Georgia",
		"Asia/UAE",
	}
}

func newPhoneNumber(seed *rand.Rand) string {
	areaCode := seed.Intn(900) + 100           // 100-999
	exchangeCode := seed.Intn(900) + 100       // 100-999
	subscriberNumber := seed.Intn(9000) + 1000 // 1000-9999

	return fmt.Sprintf("+7%d%d%d", areaCode, exchangeCode, subscriberNumber)
}

func newEmail(seed *rand.Rand) string {
	first := firstNames()
	last := lastNames()
	d := domains()

	firstName := first[seed.Intn(len(first))]
	lastName := last[seed.Intn(len(last))]
	domain := d[seed.Intn(len(d))]

	format := seed.Intn(5)
	switch format {
	case 0:
		return fmt.Sprintf("%s.%s@%s", firstName, lastName, domain)
	case 1:
		return fmt.Sprintf("%s%s@%s", firstName[:1], lastName, domain)
	case 2:
		return fmt.Sprintf("%s_%s@%s", firstName, lastName, domain)
	case 3:
		return fmt.Sprintf("%s%d@%s", firstName, seed.Intn(100), domain)
	default:
		return fmt.Sprintf("%s@%s", firstName, domain)
	}
}

func newZipCode(seed *rand.Rand) string {
	firstPart := seed.Intn(990) + 10
	secondPart := seed.Intn(900) + 100

	return fmt.Sprintf("%d%d", firstPart, secondPart)
}

func newAddress(seed *rand.Rand, zipCode string) string {
	streetNumber := seed.Intn(999) + 1

	streetNames := streets()
	cityNames := cities()
	stateNames := states()

	street := streetNames[seed.Intn(len(streetNames))]
	city := cityNames[seed.Intn(len(cityNames))]
	state := stateNames[seed.Intn(len(stateNames))]

	return fmt.Sprintf("%d %s, %s, %s %s", streetNumber, street, city, state, zipCode)
}

func newRecipientName(seed *rand.Rand) string {
	first := firstNames()
	last := lastNames()

	firstIndex := seed.Intn(len(first))
	lastIndex := seed.Intn(len(last))

	firstName := strings.ToUpper(string(first[firstIndex][0])) + first[firstIndex][1:]
	lastName := strings.ToUpper(string(last[lastIndex][0])) + last[lastIndex][1:]

	return fmt.Sprintf("%s %s", firstName, lastName)
}

func newTransactionID(seed *rand.Rand) string {
	tx := make([]byte, txLength)
	for i := range tx {
		tx[i] = charset[seed.Intn(len(charset))]
	}

	return fmt.Sprintf("TX%s", tx)
}

func newCurrency(seed *rand.Rand) string {
	c := currencies()

	return c[seed.Intn(len(c))]
}

func newBankInfo(seed *rand.Rand) (string, string) {
	b := banks()
	p := providers()

	index := seed.Intn(len(b))

	return b[index], p[index]
}

func randomTime(seed *rand.Rand) time.Time {
	now := time.Now()
	past := now.AddDate(-5, 0, 0)

	delta := now.Unix() - past.Unix()
	randomSeconds := seed.Int63n(delta)

	return past.Add(time.Second * time.Duration(randomSeconds))
}

func newItems(seed *rand.Rand) []order.Item {
	itemsCount := seed.Intn(6)

	items := make([]order.Item, 0, itemsCount)

	for range itemsCount {
		item := newItem()

		items = append(items, item)
	}

	return items
}

func newItem() order.Item {
	chrtID := rand.Int63()
	price := float64n(100000) + 1
	sale := rand.Int63n(71)
	totalPrice := price * (1 - float64(sale)/100)
	nmID := rand.Int63()
	name := itemName()

	return order.Item{
		ChrtID:     chrtID,
		Price:      price,
		Name:       name,
		Sale:       &sale,
		TotalPrice: totalPrice,
		NmID:       &nmID,
	}
}

func itemName() string {
	categories := []string{
		"Electronics", "Clothing", "Home", "Kitchen", "Beauty",
		"Garden", "Toys", "Sports", "Books", "Office",
	}

	adjectives := []string{
		"Premium", "Deluxe", "Smart", "Eco-Friendly", "Wireless",
		"Professional", "Portable", "Rechargeable", "Ergonomic", "Vintage",
		"Modern", "Luxury", "Compact", "Heavy-Duty", "Adjustable",
	}

	nouns := []string{
		"Blender", "Headphones", "Watch", "Chair", "Lamp",
		"Keyboard", "Speaker", "Bag", "Tool", "Mixer",
		"Monitor", "Mat", "Stand", "Brush", "Drill",
	}

	features := []string{
		"with Bluetooth", "with Touch Screen", "5-Piece Set", "Pro Edition",
		"2024 Model", "Extra Large", "Waterproof", "for Home & Office",
		"with 2-Year Warranty", "with Carrying Case", "10-in-1", "LED",
	}

	category := categories[rand.Intn(len(categories))]
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	feature := features[rand.Intn(len(features))]

	switch rand.Intn(5) {
	case 0:
		return fmt.Sprintf("%s %s %s", adjective, noun, feature)
	case 1:
		return fmt.Sprintf("%s %s for %s", adjective, noun, category)
	case 2:
		return fmt.Sprintf("%s %s %s - %s", category, adjective, noun, feature)
	case 3:
		return fmt.Sprintf("%s %s Series %s", brandName(), noun, feature)
	default:
		return fmt.Sprintf("%s %s %s", brandName(), adjective, noun)
	}
}

func brandName() string {
	brands := []string{
		"Apple", "Samsung", "Sony", "Nike", "Adidas",
		"Amazon", "Dell", "HP", "Bosch", "Black+Decker",
		"KitchenAid", "Dyson", "Philips", "Logitech", "Canon",
	}
	return brands[rand.Intn(len(brands))]
}

func float64n(n int64) float64 {
	decimalPart := rand.Int63n(n)
	fractionalPart := rand.Intn(100)

	return float64(decimalPart) + float64(fractionalPart%100)/100
}

func streets() []string {
	return []string{
		"Main St", "Elm St", "Oak St", "Pine St", "Maple Ave",
		"First St", "Second St", "Third St", "Fourth St", "Fifth Ave",
	}
}

func cities() []string {
	return []string{
		"New York", "Los Angeles", "Chicago", "Houston", "Phoenix",
		"Philadelphia", "San Antonio", "San Diego", "Dallas", "San Jose",
	}
}

func states() []string {
	return []string{
		"NY", "CA", "IL", "TX", "AZ", "PA", "FL", "OH", "GA", "MI",
	}
}

func firstNames() []string {
	return []string{
		"john", "jane", "alex", "emily", "michael", "sarah",
		"david", "lisa", "robert", "jessica", "william", "amanda",
	}
}

func lastNames() []string {
	return []string{
		"smith", "johnson", "williams", "brown", "jones",
		"miller", "davis", "garcia", "rodriguez", "wilson",
	}
}

func domains() []string {
	return []string{
		"gmail.com", "yahoo.com", "outlook.com", "hotmail.com",
		"protonmail.com", "icloud.com", "aol.com", "mail.com", "example.com",
	}
}

func currencies() []string {
	return []string{
		"USD", "AED", "RUB", "EUR", "GBP", "CNY",
	}
}

func providers() []string {
	return []string{
		"wbpay", "tpay", "sberpay", "alphapay", "vtbpay", "ozonpay",
	}
}

func banks() []string {
	return []string{
		"WBBank", "TBank", "Sber", "Alpha", "VTB", "OZONBank",
	}
}
