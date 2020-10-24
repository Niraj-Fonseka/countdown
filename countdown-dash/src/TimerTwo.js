import React, { useState, useEffect } from 'react';


const TimerTwo = ({ seconds, text }) => {
    // initialize timeLeft with the seconds prop
    const [timeLeft, setTimeLeft] = useState(seconds);
  
    useEffect(() => {
      // exit early when we reach 0
      if (!timeLeft) return;
  
      // save intervalId to clear the interval when the
      // component re-renders
      const intervalId = setInterval(() => {
        setTimeLeft(timeLeft - 1);
      }, 1000);
  
      // clear interval on re-render to avoid memory leaks
      return () => clearInterval(intervalId);
      // add timeLeft as a dependency to re-rerun the effect
      // when we update it
    }, [timeLeft]);
  
    return (
      <div>
        <h3>{text}</h3>
        <h1>{timeLeft / 3600}</h1>
      </div>
    );
  };

  export default TimerTwo;