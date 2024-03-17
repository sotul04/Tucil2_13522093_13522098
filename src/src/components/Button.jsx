export default function Button({onClick, type, onChangeToggle}){
    return(
        <>
        <p>
            <h2>Devide n Conquer</h2>
            <label className='switch'>
              <input type='checkbox' checked={type} onChange={onChangeToggle}></input>
              <span className='slider'></span>
            </label>
            <h2>Brute Force</h2>
        </p>
            <button onClick={onClick}>CHART!!!</button>
        </>
    )
}