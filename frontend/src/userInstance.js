import { tokenLocalStorageKey } from "./consts/tokenKey.js";

class User {
    constructor() {
        this.claims = {
            userUuid: '',
            login: '',
            role: '',
        }

        this.token = localStorage.getItem(tokenLocalStorageKey);
        if (this.token) {
            console.log('Found token - using it')
            this.extractClaims(this.token)
        }
    }

    login(token) {
        localStorage.setItem(tokenLocalStorageKey, token)
        this.token = token
        this.extractClaims(this.token)
    }

    logout() {
        this.token = null
        for (const values in Object.entries(this.claims)) {
            const key = values[0]
            this.claims[key] = ''
        }
        localStorage.removeItem(tokenLocalStorageKey)
        console.log('logged out')
    }

    extractClaims(token) {
        const base64Url = token.split('.')[1];
        const body = JSON.parse(window.atob(base64Url));
        console.log('Extracted JWT token body: ', body)
        if (body && body.UserUuid) {
            this.claims.userUuid =  body.UserUuid
        }
        if (body && body.Login) {
            this.claims.login =  body.Login
        }
        if (body && body.Role) {
            this.claims.role =  body.Role
        }
    }
}

export default User