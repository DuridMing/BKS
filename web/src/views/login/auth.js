import CryptoJS from "crypto-js";
var authcrypto = {
    encrypt(Password){
        const encodedWord = CryptoJS.enc.Utf8.parse(Password); 
        const encoded = CryptoJS.enc.Base64.stringify(encodedWord);
        return encoded.toString();
    },

}
export default authcrypto