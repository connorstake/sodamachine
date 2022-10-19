import {Grid, Button} from "@mui/material"
import { useState, useEffect } from "react"
import ProductService from "../../services/product/product.service"
import UserService from "../../services/user/user.service"
import MachineProduct from "./MachineProduct"
import DepositButton from "./DepositButton"
import AuthService from "../../services/auth"
import {useNavigate} from "react-router-dom"

import Table from "@mui/material/Table"
import TableBody from "@mui/material/TableBody"
import TableCell from "@mui/material/TableCell"
import TableContainer from "@mui/material/TableContainer"
import TableHead from "@mui/material/TableHead"
import TableRow from "@mui/material/TableRow"
import Paper from "@mui/material/Paper"
import DeleteIcon from "@mui/icons-material/Delete"
import AddProductModel from "./AddProductModal"


export default function Profile() {
  const [products, setProducts] = useState([])
  const [user, setUser] = useState()
  const [change, setChange] = useState([])
  const [sellerProducts, setSellerProducts] = useState([])
    

  const TOKEN_DENOMS = [5,10,20,50,100]

    
  const navigate= useNavigate()

  useEffect(() => {  
    console.log(localStorage.getItem("user"))  
    const fetchData  = async () => {
      let productResponse =  await ProductService.getAllProducts()
      let userResponse = await UserService.getUserInfo()

      let sellerProductResponse = await ProductService.getAllProductsBySeller()
      setProducts(productResponse["data"]["products"])
      setUser(userResponse)
      console.log(userResponse)
      setSellerProducts(sellerProductResponse["products"])
    }
    fetchData()
      .catch((error)=>{
        console.log(error)
        if (error["response"]["data"]["error"] && error["response"]["data"]["error"].includes("token is expired")){
          logout()
        }}
      )

  }, [])

  const updateDeposit = (amount) => {
    const handleDeposit = async (amount) => {
      const depositResponse = await UserService.depositFunds(amount)
      setUser({...user, deposit: depositResponse["deposit"]})  
    }
    handleDeposit(amount)
  }

  const handlePurchase = (productID) => {
    const purchaseItem = async (productID) => {
      const purchaseResponse = await ProductService.purchaseItem(productID)
      setUser({...user, deposit: purchaseResponse["deposit"]}) 
      setChange(purchaseResponse["change"]) 
    }
    purchaseItem(productID)
  }

  const handleReset= () => {
    const resetDeposit = async () => {
      const resetResponse = await UserService.resetFunds()
      setUser({...user, deposit: resetResponse["deposit"]}) 
    }
    resetDeposit()
  }

  const handleAddProduct = (productName, price, amountAvailable) => {
    const  addProduct = async () => {
      await ProductService.addProduct(productName, amountAvailable, price)

      let sellerProductResponse = await ProductService.getAllProductsBySeller()
      setSellerProducts(sellerProductResponse["products"])
    }
    addProduct()
  }

  const handleDeleteProduct = (productID) => {
    const deleteProduct = async () => {
      await ProductService.deleteProduct(productID)
   
      let sellerProductResponse = await ProductService.getAllProductsBySeller()
      setSellerProducts(sellerProductResponse["products"])
    }
    deleteProduct()
  }

  const logout = () => {
    AuthService.logout()
    navigate("/login")
  }


  if (user && user["role"] =="buyer") {
    return (
      <Grid container>
        <Grid xs={12}>
          <Grid xs={3} style={{margin:20}}>
            <Button style={{backgroundColor:"#DB504A", color: "white"}} onClick={()=>logout()}>Sign Out</Button>
          </Grid>
        </Grid>
        <Grid xs={12} container style={{ justifyContent:"space-between", display: "flex"}}>
          <Grid container style={{border:"solid rgb(112, 144, 176, .20) 1px", backgroundColor: "#8C9A9E"}}>
            <Grid container xs={10} style={{height: "100%"}}>
              {products.map((p, i)=> {
                return <MachineProduct idx={i} handlePurchase={handlePurchase} product={products[i]}/>
              })}
            </Grid>
            <Grid xs={2} style={{minHeight: 400, border:"solid rgb(112, 144, 176, .20) 1px", backgroundColor:"#747578"}}>
              <Grid xs={12} style={{textAlign:"center", padding: 20, backgroundColor: "white"}}>
                                Deposited: {user["deposit"]}
              </Grid>
              <Grid style={{textAlign:"center", backgroundColor: "#DB504A"}}>
                <Button style={{color:"white"}} onClick={()=>handleReset()}>Reset</Button>
              </Grid>


              {TOKEN_DENOMS.map((v)=> {
                return ( 
                  <Grid xs={12}>
                    <DepositButton deposit={updateDeposit} amount={v}/>
                  </Grid>
                )
              })}
              <Grid>
                {
                  change && change.length > 0 ? (
                    <Grid style={{backgroundColor:"black", color: "white", textAlign: "center"}} xs={12}>
                      <Grid >Your Change: </Grid>
                      {change.map((v)=> {
                        return (
                          <Grid>{v}</Grid>
                        )
                      })}
                    </Grid>
                  ) : <></>
                }
              </Grid>
            </Grid>
                        
                        
          </Grid>
        </Grid>
      </Grid>
    )
  } else if (user && user["role"] =="seller"){
    return (
      <Grid container style={{padding: 50}}>
        <Grid xs={12}>
          <Grid xs={3} style={{paddingBottom:20}}>
            <Button style={{backgroundColor:"#DB504A", color: "white", padding:10, borderRadius:15}} onClick={()=>logout()}>Sign Out</Button>
          </Grid>
        </Grid>

        <AddProductModel handleAddProduct={handleAddProduct}/>

                
        {sellerProducts.length > 0 ? (
          <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
              <TableHead>
                <TableRow>
                  <TableCell align="center">Product ID</TableCell>
                  <TableCell align="right">Product Name</TableCell>
                  <TableCell align="right">Price</TableCell>
                  <TableCell align="right">Amount</TableCell>
                  <TableCell align="right"></TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {sellerProducts.map((row) => (
                  <TableRow
                    key={row.productName}
                    sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
                  >
                    <TableCell align="center">{row.ID}</TableCell>
                    <TableCell align="right">{row.productName}</TableCell>
                    <TableCell align="right">{row.cost}</TableCell>
                    <TableCell align="right">{row.amountAvailable}</TableCell>
                    <TableCell align="right"><Button onClick={()=>handleDeleteProduct(row.ID)}><DeleteIcon/></Button></TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        ): null}
      </Grid>
    )
  } else {
    return (
      <div>
            Loading...
      </div>
    )
  }
    
}