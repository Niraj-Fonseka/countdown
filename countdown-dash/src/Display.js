
import React, { useState, useEffect } from 'react';
import './App.css';

function Display(props) {
   
    return ( 
        <div className="body-css">
        {props.tasks}
        </div>
    ); 
} 


export default Display;