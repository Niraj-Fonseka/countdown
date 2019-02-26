import React, { Component } from 'react';
import Timer from './Timer'

class Card extends Component {

  constructor(props){
      super(props)
  }
   
  render() {
    return (
     <div className="card-css">
        <p>{this.props.text}</p>
        <Timer timeSeconds={this.props.deadline} />
     </div>
    );
  }
}

export default Card;
