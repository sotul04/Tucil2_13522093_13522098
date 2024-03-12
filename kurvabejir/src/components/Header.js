export default function Player() {
  const [enteredPlayerName, setEnteredPlayerName] = useState(null)

  function handleClick(){
    setEnteredPlayerName(playerName.current.value)
  }

  return (
    <section id="player">
      <h2>Welcome {enteredPlayerName ?? 'unknown entity'}</h2>
      <p>
        <input type="text" />
        <button onClick={handleClick}>Set Name</button>
      </p>
    </section>
  );
}
