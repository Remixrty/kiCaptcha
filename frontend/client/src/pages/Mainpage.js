import React from 'react';
import Login from './LoginForm';
import Firstpage from './mainpage/Firstpage';
import Secondpage from './mainpage/Secondpage';
import Thirdpage from './mainpage/Thirdpage';
import Setvis from '../services/setvis';

export default function Mainpage(props) {
  // const checkJWT = Headers.

  let mainP

  if (!props.check) {
    mainP = (
      <div>
        <Firstpage />
        <Secondpage />
        <Thirdpage />
      </div>
    )
  }
  else mainP = (
    <div>
        <Firstpage />
        <Secondpage />
    </div>
  )

  // console.log(isLogged);
    return (
      <>{mainP}</>
    );

}
