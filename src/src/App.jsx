import React, { useState } from "react";
import Header from "./components/Header";
import Input from "./components/Input";
import Chart from "./components/Chart";
import DnCurves from "./components/dncCurve";
import BFCurves from "./components/Curve";
import InputFields from "./components/InputFields";
import Test from "./components/test"
function App() {
  const [enteredPoint_Iterate, setEnteredPoint_Iterate] = useState({
    Points: 0,
    Iteration: 0
  })
  const [arrayPoint, setArrayPoint] = useState([])
  // const [showChart,setShowChart] = useState(false)
  const [showChart,setShowChart] = useState({
    status: undefined //undefined belum masukin points, preview = udah masukin point, DnC = mau liat algorithm DnC, BF = mau liat algoritma brute force  
  });
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
    setShowChart(() => {
      return {
        status : "PreView"
      }
    });
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
    var temp = typeSearch;
    setTypeSearch((prevState) => !prevState);
    console.log('After: ', typeSearch);
  };

  const handleClick = async () => {
    console.log("Chart clicked");

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
        setShowChart(() => {
          return {
            status : "DnC"
          }
        });
    } catch (error) {
        console.error("Error uploading data:", error.message);
    }
};
  return (
    <section id="player">
      <Header />
      <br/>
      <div className='type-search'>
          <p>Devide n Conquer</p>
          <label className='switch'>
            <input type='checkbox' checked={typeSearch} onChange={handleCheckboxChange}></input>
            <span className='slider'></span>
          </label>
          <p>Brute Force</p>
        </div>
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
      {enteredPoint_Iterate.Points >= 2 && <InputFields nPoints={enteredPoint_Iterate.Points} data={arrayPoint} onChangeValue={handleInputChange}/>}
      <button onClick={handleClick}>CHART!!!</button>
      {/* {showChart && enteredPoint_Iterate.Points > 0 && <Chart data={arrayPoint} />} */}

      {/* {showChart.status === "PreView" && <Chart data={arrayPoint} />}
      {showChart.status === "BF" && <BFCurves data={arrayPoint}/>}
      {showChart.status === "DnC" && <DnCurves data={curvePoint} control={arrayPoint} iterate={enteredPoint_Iterate.Iteration}/>}
      <Test /> */}
      {/* {showChart && <Chart data={arrayPoint} />} */}
      <br/>
      <br/>
      <br/>
      {showChart && !typeSearch && <DnCurves 
                          data={curvePoint} 
                          control={arrayPoint} 
                          iterate={enteredPoint_Iterate.Iteration}
                          time={timeElapse}
                          type={0}
        />}
      {showChart && typeSearch && <DnCurves 
                          data={curvePoint} 
                          control={arrayPoint} 
                          iterate={enteredPoint_Iterate.Iteration}
                          time={timeElapse}
                          type={1}
        />}
        <br/>
      {showChart && typeSearch && <BezierCurves data={arrayPoint}/> }
    </section>
  );
}

export default App;