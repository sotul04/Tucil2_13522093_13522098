import React, { useState, useEffect } from "react";
import { Coordinates, Line, Mafs, Point, Theme, Text, useStopwatch } from "mafs";

function inPairs(arr) {
  const pairs = [];
  for (let i = 0; i < arr.length - 1; i++) {
    pairs.push([arr[i], arr[i + 1]]);
  }
  return pairs;
}

function getMainView(corner) {
  let min_x = corner[0][0];
  let min_y = corner[0][1];
  let max_x = corner[0][0];
  let max_y = corner[0][1];
  for (let i = 1; i < corner.length; i++) {
    if (min_x > corner[i][0]) {
      min_x = corner[i][0];
    } else if (max_x < corner[i][0]) {
      max_x = corner[i][0];
    }
    if (min_y > corner[i][1]) {
      min_y = corner[i][1];
    } else if (max_y < corner[i][1]) {
      max_y = corner[i][1];
    }
  }
  let dx = max_x-min_x;
  let dy = max_y-min_y;
  return [min_x-0.09*dx, max_x+0.09*dx, min_y-0.09*dy, max_y+0.09*dy];
}

export default function DnCurves({ data, control, iterate }) {
  const [t, setT] = useState(0.5); // State for controlling opacity
  const [iter, setIter] = useState(iterate); //show what iteration to show
  const [linePoints, setLinePoints] = useState([]); //point to draw
  const [viewContent, setViewContent] = useState([]); // kerjaan suta blok

  const opacity = 1 - (2 * t - 1) ** 6;

  const cornerPoints = control.map(p => p);

  useEffect(() => {
    refreshLines();
  }, [iter, data]);

  useEffect(() => {
    setViewContent(getMainView(control));
  }, [control]); 

  function refreshLines() {
    const newLines = [];
    let step = Math.pow(2, iterate - iter);
    console.log("step: ",iterate, iter)
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

  function pointPositoin(points, color, size) {
    return points.map(([x,y]) => (
      <Text 
        key={`${x}_${y}`} //bisa nabrak kalo ada yang masukin sama
        x={x}
        y={y}
        color={color}
        size={size}
        attach="w"
        attachDistance={15}
      >
        ({x.toFixed(1)}, {y.toFixed(1)})
      </Text>
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
      <div className="rounded">
      <Mafs 
        viewBox={{ x: [viewContent[0], viewContent[1]], y: [viewContent[2], viewContent[3]] }} 
        zoom={{ min: 0.001, max: 5 }}
      >
        <Coordinates.Cartesian
          xAxis={{ lines: (viewContent[1]-viewContent[0])/20, labels: false, axis: false }}
          yAxis={{ lines: (viewContent[1]-viewContent[0])/20,labels: false, axis: false }}
        />
        {drawLineSegments(linePoints, Theme.indigo)}
        {drawPoints(linePoints, 3, Theme.green)}
        {drawLineSegments(cornerPoints, Theme.blue)}
        {drawPoints(cornerPoints,5,Theme.pink)}
        {pointPositoin(cornerPoints,Theme.white,12)}
      </Mafs>
      </div>
      <br/>
      <p>Iteration: {iter}</p>
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
