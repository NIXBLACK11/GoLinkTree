import { User } from "../interfaces/userInterface";

export function validateSigninData(userData: User) :boolean{
    if(!userData.userName || !userData.userPassword) {
        return false;
    }
    return true;
}