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
