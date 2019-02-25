import React, { Component } from 'react'

export default class NewsDetails extends Component {
  
  constructor(props){
    super(props)
    this.state = {}
  }

  componentWillMount(){
    let URL = "http://0.0.0.0:5000/v1/sources/"

    fetch({
        url: URL,
        method: "GET",
    })
    // .then((response) => response.json())
    .then((data) => {
        console.log(">>> : ")
        console.log(data)
    })
    .catch((err) => {
        console.log(err)
    })
  }

  render() {
    return (
      <div className="border">
        <h1> News Details </h1>
      </div>
    )
  }
}
