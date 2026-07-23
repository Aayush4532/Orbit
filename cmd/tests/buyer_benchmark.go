package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func BuyerBenchmark() {
	GetLiveEvents()
	GetEventProducts()
	createBuyers()

	BenchmarkStart()

	var wg sync.WaitGroup

	for i := 1; i <= BuyerCount; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			BuyerWorker(id)
		}(i)
	}

	wg.Wait()

	BenchmarkFinish()
}

func createBuyers() {
	type job struct{ index int }

	jobs := make(chan job, BuyerCount)
	var wg sync.WaitGroup

	for w := 0; w < WorkerCount; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				client := NewClient()
				email := BuyerPrefix + strconv.Itoa(j.index) + "@gmail.com"
				SignupBuyer(client, email, BuyerPassword)
			}
		}()
	}

	for i := 1; i <= BuyerCount; i++ {
		jobs <- job{index: i}
	}
	close(jobs)

	wg.Wait()
}

func BuyerWorker(id int) {
	client := NewClient()
	email := BuyerPrefix + strconv.Itoa(id) + "@gmail.com"

	SigninBuyer(client, email, BuyerPassword)
	productID := RandomProduct()

	BuyProduct(client, productID)
}

func GetLiveEvents() {

	body, status, err := buyerClient.Get(
		"/api/buyer/events",
	)

	Assert(err)
	AssertStatus(http.StatusOK, status)

	resp := Decode[LiveEventsResponse](body)

	AssertTrue(
		len(resp.Events) > 0,
		"Get Live Events",
		"no live events found",
	)
}

func GetEventProducts() {

	body, status, err := buyerClient.Get(
		"/api/buyer/event/" + EventID,
	)

	Assert(err)
	AssertStatus(http.StatusOK, status)

	resp := Decode[BuyerProductsResponse](body)

	AssertTrue(
		len(resp.Products) > 0,
		"Get Event Products",
		"no products returned",
	)

	ProductIDs = ProductIDs[:0]

	for _, p := range resp.Products {
		ProductIDs = append(ProductIDs, p.ProductID)
	}
}

func SignupBuyer(client *Client, email, password string) {

	req := map[string]any{
		"firstName": "Benchmark",
		"lastName":  "Buyer",
		"emailId":   email,
		"age":       25,
		"password":  password,
		"role":      "buyer",
	}

	body, status, err := client.Post(
		"/api/auth/signup",
		req,
	)

	if err != nil {
		IncError()
		return
	}

	if status != http.StatusCreated {
		IncError()
		return
	}

	resp := Decode[AuthResponse](body)

	if resp.User.ID == "" {
		IncError()
	}
}

func SigninBuyer(client *Client, email, password string) {

	req := map[string]any{
		"emailId":  email,
		"password": password,
	}

	body, status, err := client.Post(
		"/api/auth/signin",
		req,
	)

	if err != nil {
		IncError()
		return
	}

	if status != http.StatusOK {
		IncError()
		return
	}

	resp := Decode[AuthResponse](body)

	if resp.User.ID == "" {
		IncError()
	}
}

func RandomProduct() string {
	return ProductIDs[rand.Intn(len(ProductIDs))]
}

func BuyProduct(client *Client, productID string) {

	start := time.Now()

	body, status, err := client.Post(
		"/api/buyer/event/"+EventID+"/purchase/"+productID,
		nil,
	)

	AddLatency(time.Since(start))

	if err != nil {
		IncError()
		return
	}

	switch status {

	case http.StatusCreated:

		resp := Decode[PurchaseResponse](body)

		if resp.Status != "" {
			IncSuccess()
		} else {
			IncError()
		}

	case http.StatusConflict:
		resp := Decode[ErrorResponse](body)
		switch resp.Error {
		case "sold out":
			IncSoldOut()
		case "already booked":
			IncAlreadyBooked()
		default:
			IncError()
		}
	default:
		IncError()
	}
}
