export default function Input({index,valuePoint,onChangeValue,...props}){
    return(
        <p className="InputXY">
            <h2 id="PointX">Point {index+ 1}</h2>
            <input {...props} 
                placeholder="X" 
                id="InputXY" 
                required value={valuePoint[0]}
                onChange={(event) => onChangeValue(index,[event.target.value,valuePoint[1]])}
                ></input>
            <input {...props} 
                placeholder="Y" 
                id="InputXY" 
                required value={valuePoint[1]}
                onChange={(event) => onChangeValue(index,[valuePoint[0],event.target.value])}
                ></input>
        </p>
    )
}
