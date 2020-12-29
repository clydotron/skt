import {React,useEffect,useState} from 'react'
import axios from 'axios'

const PingComponent = (props) => {

    const [serverResponse,setServerResponse] = useState("pending")
   
    useEffect(() => {
        axios.get("api/ping")
        .then(resp => {
            setServerResponse(resp.data.message);
        })
        .catch(resp => console.log(resp))
    })

    return (
        <h1>ping {serverResponse}</h1>
    )
}
export default PingComponent
