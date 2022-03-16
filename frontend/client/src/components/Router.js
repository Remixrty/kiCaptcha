 import { BrowserRouter, Route, Routes, Switch } from 'react-router-dom';
import CreateProject from '../pages/CreateProject';
import EditProject from '../pages/EditProject';
import GuideLine from '../pages/GuideLine';
import PersonalKab from '../pages/PersonalKab';
import Header from '../Header';
import Mainpage from '../pages/Mainpage';
import axios from 'axios';
import { useState, useEffect } from 'react'

function Router() {
  const [email, setEmail] = useState('') //use it after, when i need to get username (user kab, create project,...)
  const [name, setName] = useState('')
  const [id, setId] = useState('')
  let check = false
  useEffect(() => {
    (
      async () => {
        const respy = await axios({
          url: 'http://localhost:8080/auth/sign-check',
          method: "GET",
          withCredentials: "true",
          headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        }).then(respy => {
          // _resp = respy.data
          setName(respy.data.name) 
          setEmail(respy.data.email)
          setId(respy.data.id)
          check = true
          console.log('bye');
          // console.log(respy.data);
          // console.log(name, email, id);
        }).catch(e => {
          console.log(e)
        })
      }
      
    )
    ()
  })
      // const superName = '/'+name;
  
    return (
    <>
      <BrowserRouter>
        <Header name={name} />
        {console.log('hello')}
        <Routes>
          
          <Route path="/" element={<Mainpage check={check}/>} />
          <Route path="/username/create-project" exact element={<CreateProject />} />
          <Route path={'/'+name} element={<PersonalKab name={name}/>} />
          <Route path='/guide' exact element={<GuideLine />} />
          <Route path='/username/project-name/edit-project' exact element={<EditProject />} />
        </Routes>
      </BrowserRouter>
    </>
  );

}

export default Router;