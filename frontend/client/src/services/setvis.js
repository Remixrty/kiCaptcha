import React, { useState } from "react";
import { render } from "react-dom";
import Login from "../pages/LoginForm";

export default function Setvis(vis) {
    if (vis==true) {
        document.getElementById('loginbody').style.visibility = 'visible';
        //   
    }
    else {
        document.getElementById('loginbody').style.visibility = 'hidden';
    }
}





