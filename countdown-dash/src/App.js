import React, { Component } from 'react';
import Card from './Card'
import './App.css';

class App extends Component {
  render() {
    const items = []
    for (var i = 0 ; i < 5 ; i++){
        items.push(<Card text={"Task  : " + i } deadline={i * 100}/>)
    }
    return (
     <div className="body-css">
      {items}
    </div>
    );
  }
}

export default App;
