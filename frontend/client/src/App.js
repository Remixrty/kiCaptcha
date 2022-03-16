import './style/pics.css';
import './style/App.css';
import {BrowserRouter, Route} from 'react-router-dom'
import Firstpage from './pages/mainpage/Firstpage';
import CreateProject from './pages/CreateProject';
import EditProject from './pages/EditProject';
import GuideLine from './pages/GuideLine';
import PersonalKab from './pages/PersonalKab';
import Register from './pages/RegForm';

function App() {
  return (
    <>
        <BrowserRouter>
          <Route path="/" exact component={Firstpage}/>
          <Route path="/username/create-project" component={CreateProject}/>
          <Route path="/login" component={Register}/>
          <Route path='/username' component={PersonalKab}/>
          <Route path='/guide' component={GuideLine}/>
          <Route path='/username/project-name/edit-project' component={EditProject}/>
        </BrowserRouter>  
    </>
  );
}

export default App;
