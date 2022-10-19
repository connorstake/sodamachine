import axios from "axios"
import { authHeader } from "../auth-header"

const API_URL = "http://localhost:8080/api/"

class UserService {

  getUserInfo() {
    return axios.get(API_URL + "secured/user", { headers: authHeader() })
      .then(response => {
        return response.data
      }).catch(
        function (error) {
          return error.response.data
        }
      )
  }

  depositFunds(depositAmount) {
    return axios
      .post(API_URL + "secured/deposit", { 
        depositAmount
      },{ headers: authHeader() })
      .then(response => {
        return response.data
      })
  }

  resetFunds() {
    return axios
      .post(API_URL + "secured/user/reset", { 
      },{ headers: authHeader() })
      .then(response => {
        return response.data
      })
  }


}

export default new UserService()
