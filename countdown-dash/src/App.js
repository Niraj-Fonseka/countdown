import React, { Component } from 'react';
import Card from './Card'
import './App.css';
import { futimesSync } from 'fs';
import { format } from 'path';

class App extends Component {


  constructor(props){
    super(props)
    this.state = {
      tasks : []    
    };

  }



  componentDidMount(){

    setInterval(() =>

    fetch('http://localhost:8080/tasks',{
      method : 'GET',
    })

    .then(results => {
      console.log("Fetching data")
      return results.json()
    })
    .then(data => {
      this.setState({
        tasks :data,
        time: Date.now()
      })
    }), 1000)
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
      {items.sort(function(a,b){
        return a.props.deadline-b.props.deadline;
      })}
    </div>
    );
  }
}

export default App;
