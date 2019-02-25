import React, { Component } from 'react';
import './App.css';
import NewsListing from './components/NewsListing';
import NewsDetails  from './components/NewsDetails';
import { Container , Row, Col } from "react-bootstrap";

class App extends Component {
  render() {
    return (
      <Container fluid={true} className="App">
        <Row>
            <Col> 
               <h1 className="header border"> News App </h1>
            </Col>
        </Row>
        <Row>
            <Col xs={3}> 
              <NewsListing> </NewsListing>
            </Col>
            <Col xs={9}> 
               <NewsDetails> </NewsDetails>
            </Col>
        </Row>
      </Container>
    );
  }
}

export default App;
