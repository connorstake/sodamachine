import axios from 'axios';
import { authHeader } from '../auth-header';

const API_URL = 'http://localhost:8080/api';

class ProductService {


  getUserBoard() {
    return axios.get(API_URL + 'user', { headers: authHeader() });
  }

  getAllProducts() {
    return axios
      .get(API_URL + "/products")
  }

  getAllProductsBySeller() {
    return axios
      .get(API_URL + "/secured/products",{ headers: authHeader() })
      .then(response => {
        return response.data;
      });
  }

  purchaseItem(productID) {
    return axios
      .post(API_URL + "/secured/product/buy", { 
        productID:productID,
        amount: 1,
        
      },{ headers: authHeader() })
      .then(response => {
        return response.data;
      });
  }

  addProduct(productName, amountAvailable, price) {
    amountAvailable  = Number(amountAvailable)
    price = Number(price)
    return axios
      .post(API_URL + "/secured/product", { 
        amountAvailable,
        cost: price,
        productName,
        
      },{ headers: authHeader() })
      .then(response => {
        return response.data;
      });
  }

  deleteProduct(productID) {
    productID = Number(productID)
    return axios
      .post(API_URL + "/secured/product/delete", { 
        productID,
      },{ headers: authHeader() })
      .then(response => {
        return response.data;
      });
  }
}

export default new ProductService();
