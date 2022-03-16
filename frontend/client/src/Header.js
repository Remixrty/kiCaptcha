import './style/App.css';
import { Link, useNavigate } from 'react-router-dom';
import ScrollButton from './style/ScrollTop';
import ScrollBottomButton from './style/ScrollBottomButton';
import Login from './pages/LoginForm';
import axios from 'axios';
import { useState, useEffect } from 'react';
import Burger from './components/Burger/Burger';
// import { Burger } from './components/Burger';




function Header(props) {
    const name = props.name
    const [modalActive, setModalActive] = useState(false)
    const [hamburgerOpen, setHamburgerOpen] = useState(false)
    // const navigate = useNavigate()
    
    const toggleHamburger = () => {
        setHamburgerOpen(!hamburgerOpen)
    }
    
    if (name) {
        return (
        <>
            <div className="headcontainer">
               <Link to='/' >
                   <svg className='whitelogo'/>
                </Link>
                <div className='noWrapButtons'>
                <div className="top-button-getstarted" onClick={()=>setModalActive(true)}>
                    {/* <Link to='/' style={{textDecoration:'inherit', color: 'inherit'}}>Log In</Link> */}
                   <Link to={'/'+name} style={{textDecoration:'inherit', color: 'inherit'}}> {'Hi, ' + name} </Link>
                </div></div>
            </div>
        </>
    );}
    else return (
        <>
            <div className="headcontainer">
               <Link to='/' ><svg className='whitelogo'/></Link>
                <div className='noWrapButtons'><div className="top-button-login" onClick={()=>setModalActive(true)}>
                    {/* <Link to='/' style={{textDecoration:'inherit', color: 'inherit'}}>Log In</Link> */}
                   Log In
                </div>
                <Login active={modalActive} setActive={setModalActive} />
                
                <ScrollBottomButton />
                </div>

                <div className='hamburger' onClick={toggleHamburger}>
                   <Burger/>
                </div>
            </div>
        </>
    );
}

export default Header;