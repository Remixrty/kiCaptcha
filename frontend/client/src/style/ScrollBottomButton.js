import React, { useEffect, useState } from 'react';



const ScrollBottomButton = () => {
    const scrollToBottom = () => {
        const scrolled = document.documentElement.scrollHeight;
        window.scrollTo({
            top: document.documentElement.scrollHeight,
            behavior: 'auto'
            /* you can also use 'auto' behaviour
               in place of 'smooth' */
        });

    
    };
    // const scrollToBottom = () => {
    //     const bottomElement = document.getElementById("thirdpage");
    //     bottomElement.scrollIntoView({behavior: 'auto'})
    // };
    


    return (
        <div onClick={scrollToBottom} className="top-button-getstarted">
            Get started
        </div>
    )
}


export default ScrollBottomButton;