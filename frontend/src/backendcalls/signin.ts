import axios, { AxiosResponse, AxiosError } from 'axios';

import { setCookie } from "../utils/saveCookie";
import { User } from "../interfaces/userInterface";
import { Token } from '../interfaces/tokenInterface';

export function signinUser(userData: User) :boolean{
    axios.post('http://localhost:8080/login', userData)
        .then((response: AxiosResponse) => {
            const token: Token = response.data;
            setCookie("jwtToken", token.token);
        })
        .catch((error: AxiosError) => {
            console.log(error);
            return false;
        });
    return true;
}