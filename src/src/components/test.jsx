import React, { useState } from 'react';

function ToggleButton() {
  const [currentType, setCurrentType] = useState(1);

  const toggleType = () => {
    setCurrentType((prevType) => {
      // Cycle through types 1, 2, 3
      return prevType === 3 ? 1 : prevType + 1;
    });
  };

  return (
    <div className="toggle-button">
      <button onClick={toggleType}>
        {currentType === 1 && 'Type 1'}
        {currentType === 2 && 'Type 2'}
        {currentType === 3 && 'Type 3'}
      </button>
    </div>
  );
}

export default ToggleButton;