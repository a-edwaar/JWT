import React from "react";
import Axios from "axios";
import { useHistory } from "react-router-dom";

function Login(){

  const history = useHistory();

  function login() {
    Axios.post("/login", {
      username: 'user1',
      password: 'pass1',
    }).then((_) => {
      history.push("/private");
    }, (error) => {
      console.log(error);
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