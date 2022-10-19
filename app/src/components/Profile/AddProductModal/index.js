import * as React from "react"
import Box from "@mui/material/Box"
import Grid from "@mui/material/Grid"
import Button from "@mui/material/Button"
import Modal from "@mui/material/Modal"
import { TextField } from "@mui/material"

const style = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 400,
  bgcolor: "background.paper",
  border: "2px solid #000",
  boxShadow: 24,
  p: 4,
}

export default function AddProductModel(props) {
  const [open, setOpen] = React.useState(false)
  const handleOpen = () => setOpen(true)
  const handleClose = () => setOpen(false)

  const handleSubmit = (event) => {
    event.preventDefault()
    const data = new FormData(event.currentTarget)
    console.log({
      productName: data.get("productName"),
      amountAvailable: data.get("amountAvailable"),
      price: data.get("price")
    })
    props.handleAddProduct(data.get("productName"),  data.get("price"), data.get("amountAvailable"))
    handleClose()
  }

  return (
    <div>
      <Grid style={{backgroundColor: "#36F1CD", borderRadius: 15, padding: 5, marginBottom: 20}}>
        <Button style={{color: "white"}} onClick={handleOpen}>ADD PRODUCT</Button>
      </Grid>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box xs={12} component="form" onSubmit={handleSubmit} sx={style}>
          <Grid style={{marginBottom: 20}}>
            <TextField
              id="outlined-number"
              label="Product Name"
              name="productName"
              type="text"
              InputLabelProps={{shrink: true}}
              required
              fullWidth
            />
          </Grid>

          <Grid style={{marginBottom: 20}}>
            <TextField
              inputProps={{ inputMode: "numeric", pattern: "[0-9]*" }}
              id="outlined-number"
              label="Amount Available"
              name="amountAvailable"
              type="number"
              required
              fullWidth
            />
          </Grid>

          <Grid style={{marginBottom: 20}}>
            <TextField
              inputProps={{ inputMode: "numeric", pattern: "[0-9]*" }}
              id="outlined-number"
              label="Price"
              name="price"
              type="number"
              required
              fullWidth
            />
          </Grid>
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
              Add Product
          </Button>
        </Box>
      </Modal>
    </div>
  )
}
