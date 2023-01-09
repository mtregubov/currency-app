import React, {useState} from 'react';
import ReactDOM from 'react-dom/client';
import Button from 'react-bootstrap/Button';
import Table from 'react-bootstrap/Table';
import Form from 'react-bootstrap/Form';
import Container from 'react-bootstrap/Container';
import 'bootstrap/dist/css/bootstrap.min.css';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import './index.css';


class Title extends React.Component {
  render() {
    return (
        <Container className="p-3">
          <Container className="p-5 mb-4 bg-light rounded-3">
          <h1 className="header text-secondary">Rates from www.cnb.cz</h1>
          </Container>
        </Container>
    );
  }
}

function RateSearchForm ({handleInputChange, formInputData, handleSubmit}) {
    return (
      <Form onSubmit={handleSubmit}>
        <Row className="justify-content-md-center">
        <Col className="col-md-auto">
          <Form.Control id="curr" name="curr" type="text" value={formInputData.curr} onChange={handleInputChange} placeholder="Enter currency code"/>
        </Col>
        <Col className="col-md-3">
          <Form.Control name="date" type="text" value={formInputData.date} onChange={handleInputChange} placeholder="Enter date in 2020 in format dd.mm.yyyy"/>
        </Col>
        <Col className="col-md-auto">
          <Button variant="primary" type="submit">
            Find
          </Button>
        </Col>
        </Row>
        <Row className="justify-content-md-center">
          <Col className="col-md-auto">
          <div className="message">{formInputData.message ? <p>{formInputData.message}</p> : null}</div>
        </Col>
        </Row>
      </Form>
    );
}


function ResultTable ({tableData}) {
  return (
    <Container>
    <Table striped>
      <thead>
        <tr>
         <th>Currency</th>
          <th>Amount</th>
          <th>Rate</th>
        </tr>
      </thead>
      <tbody>
        {tableData.map((row, index) => (
          <tr key={index}>
            <td>{row.curr}</td>
            <td>{row.amount}</td>
            <td>{row.rate}</td>
            </tr>
            ))}      
      </tbody>
    </Table>
    </Container>
  );
}


function App() {
  const [tableData, setTableData] = useState([])
  const [formInputData, setformInputData] = useState(
    {
    curr: '',
    date: '',
    message: ''
    }
  );  

    const handleInputChange=(evnt)=>{  
      const newInput = (data)=>({...data, [evnt.target.name]:evnt.target.value})
    setformInputData(newInput)
    }

    const handleSubmit= (evnt) => {
      evnt.preventDefault();
      fetch("http://localhost:8080/api/rates?curr=" + formInputData.curr + "&date=" + formInputData.date)
      .then((response) => {
        if (response.status === 200) {
          return response.json();
        } else {
          const errOccured = {curr:'', date:'', message:'Some error occured'}
          setformInputData(errOccured);
          throw new Error(
            `This is an HTTP error: The status is ${response.status}`
          );          
        } 
      })
      .then((actualData) =>  {
        setTableData(actualData);
        },
      )
      .catch((err) => {
        console.log(err.message);
      });

      // const checkEmptyInput = !Object.values(formInputData).every(res=>res==="")
      // if(checkEmptyInput)
      // {
      // const emptyInput= {curr:'', date:'', message:''}
      // setformInputData(emptyInput)
      // }
    } 

    return (
      <div>
        <Title />
        <RateSearchForm handleInputChange={handleInputChange} formInputData={formInputData} handleSubmit={handleSubmit}/>
        <ResultTable tableData={tableData}/>
      </div>
    );
  } 

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(<App />);