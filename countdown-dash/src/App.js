
import React, { useState, useEffect } from 'react';
import Display from './Display'
import './App.css';
import axios from 'axios';

function App() {
    const [tasks , setItems ] = useState([])

    useEffect(() => {

      const fetchUsers = async () => {
        try {
          const response = await axios.get("http://192.168.0.16:8080/tasks");
          const items = []
          response.data.map((v,i)=>{

            if (v.deadline <= 0) {
              items.push(<div key={i} className="card" style={{animation: "blink 1s infinite"}}>
                  <div className="task">{v.task}</div>
                  <div className="countdown">{generateCountdown(v.deadline)}</div>
                  </div>)
            }else{
              items.push(<div key={i} className="card" >
                  <div className="task">{v.task}</div>
                  <div className="countdown">{generateCountdown(v.deadline)}</div>
                  </div>)              
            }
          
          })
          setItems(items);

        } catch (e) {
          setItems(...tasks);
        }
      };
      const t = setInterval(fetchUsers, 1000)
      return() => clearInterval(t);

    }, []); 

    return ( 

        <div>
            <Display tasks={tasks}/>
        </div>
    ); 
} 


function generateCountdown(deadline){
  let hours = Math.floor(deadline / (60 * 60));
    
  let divisor_for_minutes = deadline % (60 * 60);
  let minutes = Math.floor(divisor_for_minutes / 60);

  let divisor_for_seconds = divisor_for_minutes % 60;
  let seconds = Math.ceil(divisor_for_seconds);

  return <div>{hours + ' H :' + minutes + ' M :' + seconds + ' S' }</div>
}

export default App;