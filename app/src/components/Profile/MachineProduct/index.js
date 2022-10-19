
import {Grid, Button} from "@mui/material"
import bottle from "../../../img/cokebottle.png"

import React from "react"

export default function MachineProduct(props) {


  return (
    <Grid container xs={3} style={{justifyContent:"center",float:"bottom", display: "flex", alignItems:"flex-end"}}>
      <Grid xs={8} style={{textAlign:"center", backgroundColor: "#36F1CD", border:"solid rgb(112, 144, 176, .20) 1px", borderRadius: 10, margin: 10}}>
        <Button onClick={()=>props.handlePurchase(props.product["ID"])} style={{textDecoration:"none", color: "white", fontWeight: 700}}>Buy</Button>
      </Grid>
      <Grid xs={8} style={{textAlign:"center"}}>
        <img style={{width:"100%" , border:"solid rgb(112, 144, 176, .20) 1px" , backgroundColor: genRandomColor(props.idx)}} src={bottle} />
        <Grid>
          {props.product["productName"]}
        </Grid>
        <Grid >
                    Price: {props.product["cost"]}
        </Grid>       
      </Grid>
    </Grid>
  )
}


const genRandomColor = (idx) => {
  let colors =["#39A0ED", "#36F1CD", "#C6D2ED", "#DB504A", "#084C61", "#13C4A3", "#E3B505"]

  if (idx > colors.length -1) {
    idx = Math.floor(idx/colors.length)
  }
  return colors[idx]
}