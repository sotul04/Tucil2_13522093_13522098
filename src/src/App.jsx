import React, { useState, useEffect } from "react";
import Header from "./components/Header";
import Input from "./components/Input";
import Chart from "./components/Chart";
import DnCurves from "./components/dncCurve";
import BezierCurves from "./components/Curve";
import InputFields from "./components/InputFields";
import Button from "./components/Button";
function App() {
  const [enteredPoint_Iterate, setEnteredPoint_Iterate] = useState({
    Points: 0,
    Iteration: 0
  })
  const [arrayPoint, setArrayPoint] = useState([])
  const [showChart,setShowChart] = useState(false);
  const [curvePoint, setCurvePoint] = useState([]);
  const [typeSearch, setTypeSearch] = useState(false);
  const [timeElapse, setTimeElapsed] = useState(0);


  const handleInputChange = (index, value) => {
    let newValueX = parseFloat(value[0])
    let newValueY = parseFloat(value[1])
    let newValue = [newValueX,newValueY]
    const newArrayPoint = [...arrayPoint];
    newArrayPoint[index] = newValue;
    setArrayPoint(newArrayPoint);
  }

  function handleChangePoints(event) {
    setEnteredPoint_Iterate(prevState => {
      return {
        ...prevState,
        Points : event.target.value
      }

    })

    setArrayPoint(prevState => {
      let newArrayPoint;
      if (event.target.value > arrayPoint.length){
        newArrayPoint = Array(parseFloat(event.target.value)).fill([0,0])
        for (let i = 0; i < prevState.length; i++){
          newArrayPoint[i] = prevState[i]
        }
      } else if (event.target.value < arrayPoint.length){
        newArrayPoint = prevState.splice(arrayPoint.length - 1, arrayPoint.length - event.target.value)
      } else {
        return prevState
      }
      console.log(newArrayPoint);
      return newArrayPoint;
      
    
    });
  }

  const handleCheckboxChange = () => {
    console.log('Before: ', typeSearch);
    setTypeSearch((prevState) => !prevState);
    console.log('After: ', typeSearch);
  };

  const handleClick = async () => {
    console.log("Chart clicked");
    setShowChart(false)
    if (enteredPoint_Iterate.Points < 2) {
        console.log("n_point is unallowed");
        return;
    }

    // Prepare the JSON data
    const jsonData = {
      points: arrayPoint, // arrayPoint sudah berbentuk JSON
      iteration: parseInt(enteredPoint_Iterate.Iteration)
    };

    let typeEndPoint = typeSearch ? "http://localhost:8080/brute-force" : "http://localhost:8080/devidenconquer" ;

    console.log(typeEndPoint);
    try {
        const response = await fetch(typeEndPoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(jsonData),
        });

        if (!response.ok) {
          throw new Error("Failed to upload data");
        }
        
        const data = await response.json();
        console.log(data.points);
        console.log("Elapsed time: ",data.time, "ms");
        setTimeElapsed(parseFloat(data.time));
        const tempcurve = [];
        for (let i = 0; i < data.points.length; i++) {
          tempcurve.push([parseFloat(data.points[i].x), parseFloat(data.points[i].Y)]);
        }
        setCurvePoint(tempcurve);
        console.log("Curve Point:", curvePoint);
        setShowChart(true)
    } catch (error) {
        console.error("Error uploading data:", error.message);
    }
};
  return (
    <section id="player">
      <Header />
      <Input 
        title="N Points" 
        value={enteredPoint_Iterate.Points} 
        onChange={handleChangePoints}
        />
      <Input 
        title="Iteration" 
        value={enteredPoint_Iterate.Iteration} 
        onChange={(event) => setEnteredPoint_Iterate((prevState) => {
          return {
            ...prevState,
            Iteration: event.target.value
          }}
          )}
          />
      {(enteredPoint_Iterate.Points < 2 || enteredPoint_Iterate.Iteration < 1) ? <h2>Please enter a valid number</h2>: 
      <Button onClick={handleClick} type={typeSearch} onChangeToggle={handleCheckboxChange}/>}
      {enteredPoint_Iterate.Points >= 2 && <InputFields nPoints={enteredPoint_Iterate.Points} data={arrayPoint} onChangeValue={handleInputChange}/>}
      {enteredPoint_Iterate.Points >= 2 && !showChart && <Chart data={arrayPoint} />}
      {showChart && 
       !typeSearch && 
       <DnCurves 
        data={curvePoint} 
        control={arrayPoint} 
        iterate={enteredPoint_Iterate.Iteration}
        time={timeElapse/1000}
        type={0}
      />}
      {showChart && 
       typeSearch && 
       <DnCurves 
        data={curvePoint} 
        control={arrayPoint} 
        iterate={enteredPoint_Iterate.Iteration}
        time={timeElapse/1000}
        type={1}
      />}
      {showChart && typeSearch && <BezierCurves data={arrayPoint}/> }
    </section>
  );
}

export default App;