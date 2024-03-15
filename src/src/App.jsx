import { useState } from "react";
import InputXY from "./components/InputXY";
import Header from "./components/Header";
import Input from "./components/Input";
import Chart from "./components/Chart";
import { Point } from "mafs";
import DnCurves from "./components/dncCurve";
import BezierCurves from "./components/Curve";
function App() {
  const [enteredPoint_Iterate, setEnteredPoint_Iterate] = useState({
    Points: 0,
    Iteration: 0
  })
  const [arrayPoint, setArrayPoint] = useState([])
  const [showChart,setShowChart] = useState(false);
  const [curvePoint, setCurvePoint] = useState([]);
  const renderInputFields = () => {
    let nContent = 0
    let nRest = 0
    if (enteredPoint_Iterate.Points % 5 == 0){
      nContent = enteredPoint_Iterate.Points / 5;
    } else {
      nContent = Math.floor(enteredPoint_Iterate.Points /5)
      nRest = enteredPoint_Iterate.Points % 5
    }
    return(
      <div id="challenges">
        {Array.from({length: nContent},(_,arridx) => (
          <section className="challenge" key={arridx}>
            {Array.from({length: 5},(_,index) =>(
              <InputXY
                index={arridx*5 + index}
                key={index}
                valuePoint={arrayPoint[arridx*5 + index]}
                onChangeValue={handleInputChange}
                type="number"
              />
            ))}
          </section>
        ))}
        {nRest > 0 && 
        <section className="challenge" >
          {Array.from({length: nRest},(_,index) =>(
            <InputXY 
              index={(nContent)*5 + index}
              key={(nContent)*5 + index}
              valuePoint={arrayPoint[(nContent)*5 + index]}
              onChangeValue={handleInputChange}
              type="number"
            />
          ))}
        </section>}
      </div>
    )
  }

  const handleInputChange = (index, value) => {
    let newValueX = parseFloat(value[0])
    let newValueY = parseFloat(value[1])
    let newValue = [newValueX,newValueY]
    const newArrayPoint = [...arrayPoint];
    newArrayPoint[index] = newValue;
    setArrayPoint(newArrayPoint);
    setShowChart(true);
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

  const handleClick = async () => {
    console.log("Chart clicked");
    setShowChart(false);

    if (enteredPoint_Iterate.Points < 2) {
        console.log("n_point is unallowed");
        return;
    }

    // Prepare the JSON data
    const jsonData = {
      points: arrayPoint, // arrayPoint sudah berbentuk JSON
      iteration: parseInt(enteredPoint_Iterate.Iteration)
    };

    try {
        const response = await fetch("http://localhost:8080/find-curve", {
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
        const tempcurve = [];
        for (let i = 0; i < data.points.length; i++) {
          console.log(data.points[i].x, data.points[i].Y);
          tempcurve.push([parseFloat(data.points[i].x), parseFloat(data.points[i].Y)]);
        }
        setCurvePoint(tempcurve);
        console.log("Curve Point:", curvePoint);
        setShowChart(true);
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
      {enteredPoint_Iterate.Points >= 2 && renderInputFields()}
      {/* {console.log(arrayPoint)} */}
      <button onClick={handleClick}>CHART!!!</button>
      {showChart && enteredPoint_Iterate.Points > 0 && <Chart data={arrayPoint} />
}    </section>
      {/* {showChart && <Chart data={arrayPoint} />} */}
      <br/>
      <br/>
      <br/>
      {showChart && <BezierCurves data={arrayPoint}/> }
      // <br/>
      // {showChart && <DnCurves data={curvePoint} control={arrayPoint} iterate={enteredPoint_Iterate.Iteration}/>}
    </section>
  );
}

export default App;

// N Points minus muncul pop up