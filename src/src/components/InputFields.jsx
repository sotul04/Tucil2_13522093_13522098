import InputXY from "./InputXY"
export default function InputFields({nPoints, data, onChangeValue}){
    let nContent = 0
    let nRest = 0

    if (nPoints % 5 == 0){
      nContent = nPoints / 5;
    } else {
      nContent = Math.floor(nPoints /5)
      nRest = nPoints % 5
    }
    return(
      <div id="challenges">
        {Array.from({length: nContent},(_,arridx) => (
          <section className="challenge" key={arridx}>
            {Array.from({length: 5},(_,index) =>(
              <InputXY
                index={arridx*5 + index}
                key={index}
                valuePoint={data[arridx*5 + index]}
                onChangeValue={onChangeValue}
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
              valuePoint={data[(nContent)*5 + index]}
              onChangeValue={onChangeValue}
              type="number"
            />
          ))}
        </section>}
      </div>
    )

}