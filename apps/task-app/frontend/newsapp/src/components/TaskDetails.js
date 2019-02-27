import React, { Component } from 'react'

export default class TaskDetails extends Component {
  
  constructor(props){
    super(props)
    this.state = {}
  }

  componentWillMount(){
    let URL = "http://localhost:8080/v1/task"

    fetch({
        url: URL,
        method: "GET",
    })
    .then((response) => response.json())
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
        <h1> Task Details </h1>
      </div>
    )
  }
}
