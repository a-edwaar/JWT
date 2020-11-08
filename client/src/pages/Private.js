import axios from "axios";
import React, { useEffect } from "react";
import { useHistory } from "react-router-dom";

function Private(){

  let history = useHistory();

  useEffect(() => {
    axios.post("/private").then(response => {
      console.log("success")
    }).catch(error => {
      console.log(error);
      history.push("/login");
    })
  }, [history]);

  return(
    <div>
      <h1>Private page</h1>
    </div>
  )
}

export default Private;