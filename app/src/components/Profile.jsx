import {Grid, Button} from '@mui/material';

export default function Profile() {
    return (
        <Grid container>
            <Grid xs={12} container style={{border:'solid black 1px', justifyContent:'space-between', display: 'flex'}}>
                <Grid container style={{border:'solid black 1px', height: 600}}>
                    <Grid container xs={10} style={{height: '100%', border:'solid green 1px'}}>
                        <Grid container xs={3} style={{border:'solid red 1px', justifyContent:'center',float:'bottom', display: 'flex', alignItems:'flex-end', backgroundColor: 'grey'}}>
                            <Grid xs={8} style={{textAlign:'center', backgroundColor: 'green', borderRadius: 10}}>
                                <Button style={{textDecoration:'none', color: 'white'}}>Buy</Button>
                            </Grid>
                            <Grid xs={8} style={{textAlign:'center', backgroundColor: 'red'}}>
                                <Grid>
                                    Coca-Cola
                                </Grid>
                                <Grid >
                                    Cost: 50
                                </Grid>       
                            </Grid>
                        </Grid>
                        <Grid xs={3}>
                            Product 2
                        </Grid>
                        <Grid xs={3}>
                            Product 3
                        </Grid>
                        <Grid xs={3}>
                            Product 1
                        </Grid>
                        <Grid xs={3}>
                            Product 1
                        </Grid>
                        <Grid xs={3}>
                            Product 2
                        </Grid>
                        <Grid xs={3}>
                            Product 3
                        </Grid>
                        <Grid xs={3}>
                            Product 1
                        </Grid>
                        <Grid xs={3}>
                            Product 1
                        </Grid>
                        <Grid xs={3}>
                            Product 2
                        </Grid>
                        <Grid xs={3}>
                            Product 3
                        </Grid>
                        <Grid xs={3}>
                            Product 1
                        </Grid>
                        <Grid xs={3}>
                            Product 1
                        </Grid>
                        <Grid xs={3}>
                            Product 2
                        </Grid>
                        <Grid xs={3}>
                            Product 3
                        </Grid>
                        <Grid xs={3}>
                            Product 1
                        </Grid>
                        <Grid xs={12} style={{border:'solid blue 1px', display:'flex', justifyContent:'center', alignItems:'center'}}>
                            <Grid xs={8} style={{height:'70%', border:'solid grey 1px'}}></Grid>
                        </Grid>
                    </Grid>
                    <Grid xs={2}>
                        <Grid >
                            Deposited: 50
                        </Grid>
                    </Grid>
                </Grid>
            </Grid>
        </Grid>
    )
}