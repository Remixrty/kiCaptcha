import React, { cloneElement, SyntheticEvent, useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import Login from './LoginForm';
import Modal from '../services/Modal';

function Register() {
    const [Name, setName] = useState('');
    const [LastName, setLastname] = useState('');
    const [Email, setEmail] = useState('');
    const [Password, setPassword] = useState('')
    const [Phone, setPhone] = useState('');
    const navigate = useNavigate();   
    const [modalActive, setModalActive] = useState(false)
    const [ isFail, setIsFail ] = useState(true)

    async function submit(e) {
        e.preventDefault();
        if (modalActive){

        const json = JSON.stringify({ Name, LastName, Password, Phone, Email })

        await axios({
            url: 'http://localhost:8080/auth/sign-up',
            method: "POST",
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
          
            data: json
        }).then(resp => {
            console.log(resp)
            alert("Please, go to your email-address")
            scrollToTop()
        }).catch(e => {
            console.log(e);
        })}
        else 
            console.log("error");
    }

    const scrollToTop = () => {
        window.scrollTo({
            top: 0,
            behavior: 'smooth'
            /* you can also use 'auto' behaviour
               in place of 'smooth' */
        });
    };



    function isEq(rPassword) {
        if (Password != rPassword) {
            document.getElementById('testy').style.visibility = "visible"
            document.getElementById('testy').textContent = "Passwords does'nt match!"

            const elms = document.querySelectorAll("[id='floatingPassword']")
            for (var i = 0; i < elms.length; i++) {
                elms[i].style.borderColor = "red";

            }
        }
        else {
            document.getElementById('testy').style.visibility = "hidden"
            const elms = document.querySelectorAll("[id='floatingPassword']")
            for (i = 0; i < elms.length; i++) {
                elms[i].style.borderColor = "#FAFAFA";
            }
        }
    }

    function setValues() {
   
        //password length
        if (Password.length < 5) {
            console.log(Password)
            document.getElementById('testy').style.visibility = "visible"
            document.getElementById('testy').textContent = "Password must contain 5 or more symbols!"
            const elms = document.querySelectorAll("[id='floatingPassword']")
            for (var i = 0; i < elms.length; i++) {
                elms[i].style.borderColor = "red";
            }
            setIsFail(false)
        }
        else {
            const elms = document.querySelectorAll("[id='floatingPassword']")
            for (i = 0; i < elms.length; i++) {
                elms[i].style.borderColor = "#FAFAFA";
            }
            document.getElementById('testy').style.visibility = "hidden"
            setIsFail(true)
        }
        const passw = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{6,20}$/;
        if (!Password.match(passw))
        {
            console.log(Password)
            document.getElementById('testy').style.visibility = "visible"
            document.getElementById('testy').textContent = "Password must contain at least 1 numeric digit and 1 uppercase letter"
            const elms = document.querySelectorAll("[id='floatingPassword']")
            for (var i = 0; i < elms.length; i++) {
                elms[i].style.borderColor = "red";
            }
            setIsFail(false)
        }
        else {
            const elms = document.querySelectorAll("[id='floatingPassword']")
            for (i = 0; i < elms.length; i++) {
                elms[i].style.borderColor = "#FAFAFA";
            }
            document.getElementById('testy').style.visibility = "hidden"
            setIsFail(true)
        }

        //password contains
        console.log(isFail);

        setModalActive(isFail)
    }



    return (
        <>
            <div className="mainform">
                <form onSubmit={submit}>
                    <p className="formText">Get started</p>
                    <main className="form-signin">
                        <div className='columnOne'>
                            <div className="superField">
                                <p className='labelText'>First name*</p>
                                <input type="text" className='formField' id="floatingNameInput" required
                                    onChange={e => setName(e.target.value)}
                                />
                            </div>
                            <div className="superField">
                                <p className='labelText'>Last name*</p>
                                <input type="text" className='formField' id="floatingLastNameInput" required
                                    onChange={e => setLastname(e.target.value)}
                                />
                            </div>
                            <div className="superField">
                                <p className='labelText'>Email*</p>
                                <input type="email" className='formField' id="floatingEmailInput" required
                                    onChange={e => setEmail(e.target.value)}
                                />
                            </div>

                        </div>
                        <div className='columnTwo'>
                            <div className="superField">
                                <p className='labelText'>Password*</p>
                                <input type="password" className='formField' id="floatingPassword" required
                                    onChange={e => setPassword(e.target.value)}
                                />
                            </div>
                            <div className="superField">
                                <p className='labelText'>Confirm password*</p>
                                <input type="password" className='formField' id="floatingPassword" required
                                    onChange={e => isEq(e.target.value)}
                                />
                                <p className='errorPasswordText' id='testy'>Error</p>
                            </div>
                            <div className="superField">
                                <p className='labelText'>Phone number</p>
                                <input type="number" className='formField' id="floatingInput"
                                    onChange={e => setPhone(e.target.value)}
                                />
                            </div>
                        </div>
                        <div className="superField">
                            <input type="checkbox" className="custom-checkbox" id="happy" name="happy" value="no" required />
                            <label htmlFor="happy" className='policyText'>I have read and accept the privacy policy</label>
                        </div>

                    </main>
                    <button className='submitButton' type="submit" onClick={setValues}>Create account</button>
                </form>
            </div>
            {/* <Modal active={modalActive} setActive={setModalActive}>
                
            </Modal> */}
        </>
    );
}

export default Register;