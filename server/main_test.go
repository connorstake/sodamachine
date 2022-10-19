package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/mvpmatch/server/models"
	"github.com/stretchr/testify/assert"
)

type TokenResponse struct {
	Token string `json:"token"`
}

var buyerUsername = "testBuyer"
var sellerUsername = "testSeller"
var password = "12345678"

func getUserInfoHelper(username string) models.User {
	token := setupToken(username, password)
	jsonBody := []byte("{}")
	getUserBodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/secured/user", getUserBodyReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token.Token)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(bodyBytes))

	var user models.User

	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		panic("Unable to convert the JSON string to a struct")
	}

	fmt.Println(user)
	return user

}

func getProductsBySellerHelper() (productsReturn []models.Product) {
	token := setupToken(sellerUsername, password)
	jsonBody := []byte("{}")
	getUserBodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/secured/products", getUserBodyReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token.Token)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(bodyBytes))

	var products struct{ Products []models.Product }

	err = json.Unmarshal(bodyBytes, &products)
	if err != nil {
		panic("Unable to convert the JSON string to a struct")
	}

	fmt.Println(products)
	return products.Products

}

func setupToken(username string, password string) (r TokenResponse) {
	creds := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)
	jsonBody := []byte(creds)
	tokenBodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/token", tokenBodyReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)

	var tr TokenResponse

	err = json.Unmarshal(bodyBytes, &tr)
	if err != nil {
		panic("Unable to convert the JSON string to a struct")
	}

	return tr

}

func TestCreateUsers(t *testing.T) {
	jsonBody := []byte(`{"username": "testBuyer", "password": "12345678", "role": "buyer"}`)
	buyerBodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/user/register", buyerBodyReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)

	var buyer models.User

	err = json.Unmarshal(bodyBytes, &buyer)
	if err != nil {
		fmt.Println("Unable to convert the JSON string to a struct")
	}

	assert.Equal(t, 201, res.StatusCode)

	jsonBody = []byte(`{"username": "testSeller", "password": "12345678", "role": "seller"}`)
	sellerBodyReader := bytes.NewReader(jsonBody)
	req, _ = http.NewRequest("POST", "http://localhost:8080/api/user/register", sellerBodyReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err = client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	bodyBytes, err = ioutil.ReadAll(res.Body)

	var seller models.User

	err = json.Unmarshal(bodyBytes, &seller)
	if err != nil {
		fmt.Println("Unable to convert the JSON string to a struct")
	}

	assert.Equal(t, 201, res.StatusCode)

}

func TestAddProduct(t *testing.T) {
	token := setupToken(sellerUsername, password)
	user := getUserInfoHelper(sellerUsername)
	body := fmt.Sprintf(`{"sellerID": %v, "amountAvailable": 150, "productName": "Sprite", "cost": 5}`, user.ID)
	jsonBody := []byte(body)
	createProdBodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/secured/product", createProdBodyReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token.Token)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(res)
}

func TestDeposit(t *testing.T) {
	jsonBody := []byte(`{"depositAmount": 100}`)
	depositBodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/secured/deposit", depositBodyReader)
	token := setupToken(buyerUsername, password)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token.Token)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(bodyBytes))

	assert.Equal(t, 200, res.StatusCode)

	//TEST FAIL
	jsonBody = []byte(`{"depositAmount": 90}`)
	depositBodyReader = bytes.NewReader(jsonBody)
	req, _ = http.NewRequest("POST", "http://localhost:8080/api/secured/deposit", depositBodyReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token.Token)

	res, err = client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	bodyBytes, err = ioutil.ReadAll(res.Body)
	fmt.Println(string(bodyBytes))

	assert.Equal(t, 400, res.StatusCode)

}

func TestBuyProduct(t *testing.T) {
	token := setupToken(buyerUsername, password)
	products := getProductsBySellerHelper()
	fmt.Println(products)
	body := fmt.Sprintf(`{"amount": 1, "productID": %v}`, products[0].ID)
	jsonBody := []byte(body)
	buyBodyReader := bytes.NewReader(jsonBody)
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	req, _ := http.NewRequest("POST", "http://localhost:8080/api/secured/product/buy", buyBodyReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token.Token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(bodyBytes))

	type Result struct {
		Change  []int `json:"change"`
		Deposit int   `json:"deposit"`
	}
	var change Result

	err = json.Unmarshal(bodyBytes, &change)
	if err != nil {
		fmt.Println("Unable to convert the JSON string to a struct")
	}

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, []int{50, 20, 20, 5}, change.Change)

}

func TestDeleteUsers(t *testing.T) {
	jsonBody := []byte("{}")
	deleteBodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/secured/user/delete", deleteBodyReader)
	token := setupToken(buyerUsername, password)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token.Token)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	assert.Equal(t, 201, res.StatusCode)

	jsonBody = []byte("{}")
	deleteBodyReader = bytes.NewReader(jsonBody)
	token = setupToken(sellerUsername, password)

	req, _ = http.NewRequest("POST", "http://localhost:8080/api/secured/user/delete", deleteBodyReader)
	req.Header.Add("Authorization", token.Token)

	res, err = client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	assert.Equal(t, 201, res.StatusCode)

}
