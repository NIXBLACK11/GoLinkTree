import Cookies from 'js-cookie';

// Function to set a cookie
export const setCookie = (key: string, value: string, options?: Cookies.CookieAttributes): void => {
    Cookies.set(key, value, options);
};

// Function to get a cookie
export const getCookie = (key: string): string | undefined => {
    return Cookies.get(key);
};

// Function to remove a cookie
// const removeCookie = (key: string): void => {
//     Cookies.remove(key);
// };

// Example usage:
// setCookie('user', 'John Doe');
// const user = getCookie('user');
// console.log('User:', user); // Output: User: John Doe
// removeCookie('user');
