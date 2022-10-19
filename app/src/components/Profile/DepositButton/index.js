
import {Grid, Button} from "@mui/material"

import React from "react"

export default function DepositButton(props) {


  return (
    <Grid container xs={12} style={{justifyContent:"center",float:"bottom", display: "flex", alignItems:"flex-end", backgroundColor: "#747578"}}>
      <Button style={{textDecoration:"none", color: "white"}} onClick={()=>props.deposit(props.amount)}>Deposit {props.amount}</Button>
    </Grid>
  )
}