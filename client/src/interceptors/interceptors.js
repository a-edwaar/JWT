import axios from "axios";
import jwt_decode from "jwt-decode";
import { getAccessToken, setAccessToken } from "../auth/accessToken";

export default function setupInterceptors(){
  axios.interceptors.request.use(async config => {
    console.log("interceptor triggered")
    if (config.url.includes("/login")){
      console.log("logging in..leaving request as is")
      return config // don't alter anything we are just logging in
    }
    let token = getAccessToken()
    if (token === ""){
      // call refresh
      await fetch('/refresh')
      .then(response => response.json())
      .then((data) => {
        console.log("access token empty -> refreshing...")
        token = data.Token;
        setAccessToken(token);
        config.headers.Authorization = 'Bearer '+ token;
      }).catch(error => {
        console.log("refresh token invalid")
        return Promise.reject(error);
      });
    }else{
      // check if access token has expired
      const {exp} = jwt_decode(token)
      if (Date.now() >= exp*1000) {
        // call refresh
        await fetch('/refresh')
        .then(response => response.json())
        .then((data) => {
          console.log("access token expired -> refreshing...")
          token = data.Token;
          setAccessToken(token);
          config.headers.Authorization = 'Bearer '+ token;
        }).catch(error => {
          console.log("refresh token invalid")
          return Promise.reject(error);
        });
      }else{
        console.log("access token valid")
        config.headers.Authorization = 'Bearer '+ token; // access token still good
      }
    }
    return config;
  });
}