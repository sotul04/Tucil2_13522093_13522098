import { useState } from "react";
import InputXY from "./components/InputXY";
import Header from "./components/Header";
import Input from "./components/Input";
import Chart from "./components/Chart";
function App() {
  const [enteredPoint_Iterate, setEnteredPoint_Iterate] = useState({
    Points: 0,
    Iteration: 0
  })
  const [arrayPoint, setArrayPoint] = useState([0,0])
  const [showChart,setShowChart] = useState(false);
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
  }

  function handleChangePoints(event) {
    setEnteredPoint_Iterate(prevState => {
      return {
        ...prevState,
        Points : event.target.value
      }

    })
    setArrayPoint(Array(parseFloat(event.target.value)).fill([0,0]));
    setShowChart(false)
  }

  function handleClick(){
    setShowChart((prevState) => !prevState)
  }
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
      {console.log(arrayPoint)}
      <button onClick={handleClick}>CHART!!!</button>
      {showChart && <Chart data={arrayPoint}/>}
    </section>
  );
}

export default App;
