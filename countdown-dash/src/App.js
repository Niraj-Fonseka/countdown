import React from 'react';
import { ThemeProvider, createTheme, Words, Frame } from 'arwes';
import { useState, useEffect } from 'react';
import axios from 'axios';
import Display from './Display';


function App(){
    const [tasks , setItems ] = useState([])

    useEffect(() => {

      const fetchUsers = async () => {
        try {
          const response = await axios.get('http://localhost:8080/tasks');
          const items = []
          response.data.map((v,i)=>{
            if (v.deadline <= 0) {
              items.push(
                <div className="frame-wrapper" style={{animation: "blink 1s infinite"}}>
                    <Frame animate={true} level={3} corners={4} layer='primary'>
                            <div className="task"><Words animate>{v.task}</Words></div>
                            <div className="countdown">{generateCountdown(v.deadline)}</div>  
                    </Frame>
                </div>
                )
            }else{
              items.push(
                <div className="frame-wrapper">

                    <Frame animate={true} level={3} corners={4} layer='primary'>
                            <div className="task"><Words animate>{v.task}</Words></div>
                            <div className="countdown">{generateCountdown(v.deadline)}</div>  
                    </Frame>
                    
                </div>
              )}         
          })
          setItems(items);

        } catch (e) {
          setItems(...tasks);
        }
      };

      const t = setInterval(fetchUsers, 1000)

      return() => clearInterval(t);

    }, []); 

    return(
    <ThemeProvider theme={createTheme()}>
            <Display tasks={tasks}/>
    </ThemeProvider>
    )
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