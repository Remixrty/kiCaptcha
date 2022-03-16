import React, { SyntheticEvent, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import '../style/LoginForm.css'
import '../style/App.css'
import Setvis from '../services/setvis';
import { useCallback, useEffect } from 'react';
import Modal from '../services/Modal';

const Login = ({active, setActive, children}) => {    // const [isVis, setIsVis] = useState(false)
    const [Email, setEmail] = useState('');
    const [Password, setPassword] = useState('');
    const [forgot, setForgot] = useState(false)
    const navigate = useNavigate();
    async function submit(e) {
        e.preventDefault();

        const json = JSON.stringify({ Email, Password })
        await axios({
            url: 'http://localhost:8080/auth/sign-in',
            method: "POST",
            // mode: "no-cors",
            withCredentials: "true",
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
            data: json
        }).then(resp => {
            console.log(resp.data)
        }).catch(e => {
            console.log(e);
        })

        navigate('/')
    }

    function closeForm(e) {
        setActive(false)
        document.getElementById('mainForm').reset()
    }



    if (!forgot)
    return (
                
                // <div className='loginBody' id='loginbody'>
                <div className={active ? "modal active" : 'modal'} onClick={e=> closeForm(e)}>
                <div className='modal__content' onClick={e=>e.stopPropagation()}>
                    <div className="mainform-Login">
                        <main>
                            <form onSubmit={submit} id='mainForm'>
                                <img src='../../pics/vector.svg' className="vector" width={"72px"} height={"72px"} />
                                <p className='formText'>I'm already signed up</p>
                                <div className='superField'>
                                    <p className='labelText'>Email</p>
                                    <input type="email" className="formField"  required
                                        onChange={e => setEmail(e.target.value)}
                                    /></div>
                                <div className='superField'>
                                    <p className='labelText'>Password</p>
                                    <input type="password" className="formField"  required
                                        onChange={e => setPassword(e.target.value)}
                                    />
                                </div>
                                <p className='labelMicroText' onClick={()=>setForgot(true)}>Forgot password?</p>

                                <div className='superFieldBut'>
                                    <button className='submitButton' type="submit">Log In</button></div>

                                <div className='superFieldGetStarted'>
                                    <p className='loginText'>No account?&nbsp;</p>
                                    <p className='loginText'>Get started</p>

                                </div>
                                <p className='bottomText'>This site is protected by kurCaptcha and the Boodle Privacy Policy
                                    and Terms of Service apply.</p>
                            </form>
                        </main>
                    </div>
                    </div></div>
        );
        else {
            return (
                <div className={active ? "modal active" : 'modal'} onClick={e=> closeForm(e)}>
                <div className='modal__content' onClick={e=>e.stopPropagation()}>
                    <div className="mainform-Login">
                        <main>
                            <form onSubmit={submit} id='mainForm'>
                                <img src='../../pics/vector.svg' className="vector" width={"72px"} height={"72px"} />
                                <p className='formText'>Enter yout email to reset password</p>
                                <div className='superField'>
                                    <p className='labelText'>Email</p>
                                    <input type="email" className="formField"  required
                                        onChange={e => setEmail(e.target.value)}
                                    /></div>
                                <div className='superFieldBut'>
                                    <button className='submitButton' type="submit">Reset password</button></div>
                                    <div className='superFieldCancel'>
                                        <p className='cancelBut' onClick={()=>setForgot(false)}>Cancel</p>
                                    </div>
                            </form>
                        </main>
                    </div>
                    </div></div>
        
            )
        }
    
}

export default Login;