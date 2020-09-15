import ReactDOM from 'react-dom';
import React  from 'react';
 
class App extends React.Component {
  state = {bgColor: "black"}
  changeCircleColor = (colorinfo) => {
    this.setState({bgColor:colorinfo})
    console.log(colorinfo);
  };

  render() {
    return (
      <div>
      <Form onSubmit={this.changeCircleColor} bgColor={this.state.bgColor}/>
      </div>
    );
  }
}

class Form extends React.Component {
  state = {colorName: ""}

  handleSubmit = async(event) => {
    event.preventDefault();
    const resp = this.state.colorName
    this.props.onSubmit(resp)
    this.setState({colorName:""})
  }

	render() {
  	return (
    	<form onSubmit={this.handleSubmit}>
      <span className="formtext">&#x3C;App&#x3E;</span>
        <div className="circle" style={{backgroundColor:this.props.bgColor}}></div>
    	  <input 
          type="text" 
          value={this.state.colorName}
          onChange={event => this.setState({colorName: event.target.value})}
          placeholder="Enter Color Name" 
          required 
        />
        <button>Change Circle Color</button>
    	</form>
    );
  }
}

ReactDOM.render(
  <App />, 
  document.getElementById('root'),
);
 
export default App;