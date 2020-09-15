import ReactDOM from 'react-dom';
import React, { useState, useRef } from 'react';

const App = props => {
  const [color, setColor] = useState('black')
  const inputRef = useRef(null)

  const handleButtonClick = () => {
    setColor(inputRef.current.value || 'black')
  }

  return (
    <div id="container">
      <input type="text" ref={inputRef} placeholder="Enter Color Name"/>
      <input type="button" onClick={handleButtonClick} value="Change Circle Color" />
      <div id="circle" style={{ backgroundColor: color }} />
    </div>
  )
}

ReactDOM.render(
  <App />,
  document.getElementById('root'),
);

export default App;