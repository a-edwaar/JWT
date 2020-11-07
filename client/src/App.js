import './App.css';
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Login from './pages/Login';
import Private from './pages/Private';

function App() {
  return (
    <Router>
      <div className="App">
      <Switch>
          <Route path="/login" component={Login} />
          <Route path="/private" component={Private} />
        </Switch>
    </div>
    </Router>
  );
}

export default App;
