export default function Input({title,value,onChange,...props}){
    return (
        <p>
            <h2>{title}</h2>
            <input 
                type="number"
                required 
                value={value}
                onChange={onChange}
                {...props}
            />
        </p>
    )
}
// const renderInputFields = () => {
//     return (
//       <ul>
//         {Array.from({length: enteredPoint_Iterate.Points},(_,index) =>(
          
//           <InputXY 
//             key={index} 
//             type="number" 
//             index={index}
//             valuePoint={arrayPoint[index]} 
//             onChangeValue ={handleInputChange}
//             />
//         ))}
//       </ul>
//     )
//   }