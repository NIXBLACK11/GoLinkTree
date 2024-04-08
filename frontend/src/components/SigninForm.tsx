import { Link, useNavigate } from "react-router-dom";
import React, { useState } from 'react';

import { User } from "../interfaces/userInterface";
import { validateSigninData } from "../utils/validateSignin";
import { Alert, AlertProps } from "./Alert";
import { signinUser } from "../backendcalls/signin";

export const SigninForm = () => {
    const [details, setDetails] = useState<User>({userName: "null", userPassword: "null"});
    const navigate = useNavigate();
    const [alert, setAlert] = useState<AlertProps | null>(null); // State to manage alert

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        console.log(details);
        const { name, value } = e.target;
        setDetails(prevState => ({
            ...prevState,
            [name]: value
        }));
    };

    const signin = async () => {
        const valid = validateSigninData(details);
        if (!valid) {
            setAlert({
                type: 'error',
                heading: 'Signin Error',
                data: 'Invalid signin data provided. Please check your inputs.',
            });
            return;
        }
        
        const successSignIn = signinUser(details);
        console.log(successSignIn);
        if (!successSignIn) {
            setAlert({
                type: "warning" ,
                heading: "Invalid credentials", 
                data: "Check username or password"
            });
            return;
        }
        

        if(await successSignIn && valid) {
            setAlert({
                type: "success" ,
                heading: "Successfully signed in" ,
                data: "success" 
            });
            {navigate(`/${details.userName}`)}
        }
    }

    return <>
        <div className="bg-white dark:bg-gray-900 ">
            <div className="flex justify-center h-screen">
                <div className="hidden bg-cover lg:block lg:w-2/3" style={{ backgroundImage: 'url("https://images.unsplash.com/photo-1616763355603-9755a640a287?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80")' }}>
                    <div className="flex items-center h-full px-20 bg-gray-900 bg-opacity-40">
                        <div>
                            <h2 className="text-2xl font-bold text-white sm:text-3xl">Meraki UI</h2>

                            <p className="max-w-xl mt-3 text-gray-300">
                                Lorem ipsum dolor sit, amet consectetur adipisicing elit. In
                                autem ipsa, nulla laboriosam dolores, repellendus perferendis libero suscipit nam temporibus
                                molestiae
                            </p>
                        </div>
                    </div>
                </div>

                <div className="flex items-center w-full max-w-md px-6 mx-auto lg:w-2/6">
                    <div className="flex-1">
                        <div className="text-center">
                            <div className="flex justify-center mx-auto">
                                <img className="w-auto h-7 sm:h-8" src="https://merakiui.com/images/logo.svg" alt=""/>
                            </div>

                            <p className="mt-3 text-gray-500 dark:text-gray-300">Sign in to access your account</p>
                        </div>

                        <div className="mt-8">
                            <div className="mt-6">
                                <label className="block mb-2 text-sm text-gray-600 dark:text-gray-200">Username</label>
                                <input 
                                    type="string" 
                                    name="userName" 
                                    id="userName" 
                                    placeholder="john1123" 
                                    onChange={handleChange}
                                    className="block w-full px-4 py-2 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-lg dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40" 
                                />
                            </div>
                            
                            <div className="mt-6">
                                <div className="flex justify-between mb-2">
                                    <label className="text-sm text-gray-600 dark:text-gray-200">Password</label>
                                    <a href="#" className="text-sm text-gray-400 focus:text-blue-500 hover:text-blue-500 hover:underline">Forgot password?</a>
                                </div>

                                <input 
                                    type="password" 
                                    name="userPassword" 
                                    id="userPassword" 
                                    placeholder="Your Password" 
                                    onChange={handleChange}
                                    className="block w-full px-4 py-2 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-lg dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40" 
                                />
                            </div>

                            <div className="mt-6">
                                <button 
                                    onClick={signin}
                                    className="w-full px-4 py-2 tracking-wide text-white transition-colors duration-300 transform bg-blue-500 rounded-lg hover:bg-blue-400 focus:outline-none focus:bg-blue-400 focus:ring focus:ring-blue-300 focus:ring-opacity-50"
                                >
                                    Sign in
                                </button>
                            </div>
                            <Link to="/signup">
                                <p className="mt-6 text-sm text-center text-gray-400">
                                    Don't have an account yet? <span className="text-blue-500 hover:underline">Sign up</span>.
                                </p>
                            </Link>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </>
}