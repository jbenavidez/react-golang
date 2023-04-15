import './App.css';
import Navbar from './components/Navbar';
import Home from './components/Home';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom/cjs/react-router-dom.min';
import Create from './components/Create';
import BlogDetails from './components/BlogDetails';

function App() {
  return (
    <Router> 
      <div className="App">
        <Navbar />
        <div className='content'>
        <Switch>
          <Route exact path='/' component={Home} />
          <Route exact path='/create' component={Create} />
          <Route exact path='/blogs/:id' component={BlogDetails} />
        </Switch>
        </div>
      </div>
    </Router>
  );
}

export default App;
