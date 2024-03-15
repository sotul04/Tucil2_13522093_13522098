import React, { useState, useEffect } from "react";
import { Coordinates, Line, Mafs, Point, Theme, useMovablePoint } from "mafs";
import { easeInOutCubic } from "js-easing-functions";

function inPairs(arr) {
  const pairs = [];
  for (let i = 0; i < arr.length - 1; i++) {
    pairs.push([arr[i], arr[i + 1]]);
  }
  return pairs;
}

export default function DnCurves({ data, control, iterate }) {
  const [t, setT] = useState(0.5); // State for controlling opacity
  const [iter, setIter] = useState(1);
  const [linePoints, setLinePoints] = useState([]);

  const opacity = 1 - (2 * t - 1) ** 6;

  const corner = control.map(([x, y]) => useMovablePoint([x, y]));
  const cornerPoints = corner.map(point => point.point);

  useEffect(() => {
    refreshLines();
  }, [iter, data]);

  let lines = data.map(([x, y]) => [x, y]);

  function refreshLines() {
    const newLines = [];
    let step = Math.pow(2, iterate - iter);
    for (let i = 0; i < data.length; i += step) {
      newLines.push(data[i]);
    }
    setLinePoints(newLines);
  }

  function drawLineSegments(pointPath, color) {
    return inPairs(pointPath).map(([p1, p2], index) => (
      <Line.Segment
        key={index}
        point1={p1}
        point2={p2}
        opacity={opacity}
        color={color}
      />
    ));
  }

  function drawPoints(points, rad, color) {
    return points.map((point, index) => (
      <Point
        key={index}
        x={point[0]}
        y={point[1]}
        svgCircleProps={{ r: rad }}
        color={color}
      />
    ));
  }

  const handleInputChange = (event) => {
    const newIter = +event.target.value;
    setIter(newIter);
  };

  return (
    <div id="susy">
      <br />
      <Mafs viewBox={{ x: [-5, 5], y: [-4, 4] }} zoom={{ min: 0.001, max: 5 }}>
        <Coordinates.Cartesian
          xAxis={{ labels: false, axis: false }}
          yAxis={{ labels: false, axis: false }}
        />
        {drawLineSegments(linePoints, Theme.indigo)}
        {drawPoints(linePoints, 3, Theme.green)}
        {drawLineSegments(cornerPoints, Theme.blue)}
        {corner.map(point => point.element)}
      </Mafs>
      <div className="p-4 border-gray-700 border-t bg-black text-white">
        <input
          type="range"
          min={1}
          max={iterate}
          step={1}
          value={iter}
          onChange={handleInputChange}
        />
      </div>
    </div>
  );
}
