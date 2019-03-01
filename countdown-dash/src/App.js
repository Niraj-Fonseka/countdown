import React, { Component } from 'react';
import Card from './Card'
import './App.css';
import { futimesSync } from 'fs';
import { format } from 'path';

class App extends Component {


  constructor(props){
    super(props)
    this.state = {
      tasks : [],
    };
  }


  componentDidMount(){
   
    fetch('http://localhost:8080/tasks',{
      method : 'GET',
    })
    .then(results => {
      console.log("Fetching data")
      return results.json()
    })
    .then(data => {
      this.setState({
        tasks :data
      })
    })
  }

  render() {
    console.log("Render called")
    const items = []
    console.log(this.state.tasks.length)
    for (var i = 0 ; i < this.state.tasks.length ; i++){
        items.push(<Card text={this.state.tasks[i].task} deadline={this.state.tasks[i].deadline}/>)
    }
    return (
     <div className="body-css">
      {items}
    </div>
    );
  }
}

export default App;
