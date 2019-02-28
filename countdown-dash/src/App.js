import React, { Component } from 'react';
import Card from './Card'
import './App.css';

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
      crossDomain :true,
      headers : { 'Access-Control-Allow-Origin':'*'}
    })
    .then(results => {
      console.log("Fetching data")
      console.log(results.json())
      return results.json();
    })
  }
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
