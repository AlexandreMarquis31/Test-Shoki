import React, {useState} from "react";
import TextField from "@material-ui/core/TextField";
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";

const AVGAPI = "http://localhost:8083/avg?url=";
const WORDSAPI = "http://localhost:8083/words?url=";

function Home() {
    const fetchData = () => {
        fetch(AVGAPI + url).then(response => response.json())
            .then(data => setAVG(data));
        fetch(WORDSAPI + url).then(response => response.json())
            .then(data => setWords(data));
    };
    var [url, setUrl] = useState("");
    var [avg, setAVG] = useState(0);
    var [words, setWords] = useState(0);
    return (<Grid
        container
        spacing={0}
        direction="column"
        alignItems="center"
        justify="center"
        style={{minHeight: '80vh'}}
    >
        <Grid item xs={18}>
            <TextField style={{marginLeft:"-50%",width:"200%"}} onChange={event => url = setUrl(event.target.value)}/>
        </Grid>
        <Grid item xs={12}>
            <Button variant="contained" color="primary" style={{margin: "20px"}} onClick={fetchData}>Valider</Button>
        </Grid>
        <Grid item xs={12}>
            <Typography align="left" variant="h5" component="h5">
                WORDS : {words}
            </Typography>
        </Grid>
        <Grid item xs={12}>
            <Typography align="left" variant="h5" component="h5">
                AVG : {avg}
            </Typography>
        </Grid>
    </Grid>);
}

export default Home;
