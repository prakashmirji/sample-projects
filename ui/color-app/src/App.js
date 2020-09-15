import ReactDOM from 'react-dom';
import React, {useState, useRef}  from 'react';

const App = props => {
	const [color, setColor] = useState('black')
  const inputRef = useRef(null)

  const handleButtonClick = () => {
      setColor(inputRef.current.value || 'black')
  }

  return (
  	<div id="container">
      <div id="circle" style={{backgroundColor: color}} />
  	  <input type="text" ref={inputRef} />
      <input type="button" onClick={handleButtonClick} value="Change Circle Color" />
  	</div>
  )
}

ReactDOM.render(
  <App />, 
  document.getElementById('root'),
);

export default App;