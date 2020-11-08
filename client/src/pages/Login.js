import React from "react";
import axios from "axios";
import { useHistory } from "react-router-dom";
import { setAccessToken } from "../auth/accessToken";

function Login(){

  const history = useHistory();

  function login() {
    axios.post("/login", {
      username: 'user1',
      password: 'pass1',
    }).then(response => {
      if (response.data.Token){
        setAccessToken(response.data.Token)
        history.push("/private");
      }
    }).catch(error => {
      console.log(error)
    })
  }
  
  return(
    <div>
      <h1>Login Page</h1>
      <button onClick={login}>Login</button>
    </div>
  )
}

export default Login;