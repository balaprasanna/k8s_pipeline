import React, { Component } from 'react';
import './App.css';
import TaskListing from './components/TaskListing';
import TaskDetails  from './components/TaskDetails';
import { Container , Row, Col } from "react-bootstrap";

class App extends Component {
  render() {
    return (
      <Container fluid={true} className="App">
        <Row>
            <Col> 
               <h1 className="header border"> Task App </h1>
            </Col>
        </Row>
        <Row>
            <Col xs={3}> 
              <TaskListing> </TaskListing>
            </Col>
            <Col xs={9}> 
               <TaskDetails> </TaskDetails>
            </Col>
        </Row>
      </Container>
    );
  }
}

export default App;
