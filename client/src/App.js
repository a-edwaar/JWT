import './App.css';
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import Login from './pages/login';
import Private from './pages/private';
import Home from './pages/home';
import setupInterceptors from './interceptors/interceptors';

setupInterceptors();

function App() {
  return (
    <Router>
      <div className="App">
      <div className="Nav">
        <Link to="/">Home</Link>
        <Link to="/login">Login</Link>
        <Link to="/private">Private</Link>
      </div>
      <Switch>
          <Route path="/" exact component={Home} />
          <Route path="/login" component={Login} />
          <Route path="/private" component={Private} />
        </Switch>
    </div>
    </Router>
  );
}

export default App;
