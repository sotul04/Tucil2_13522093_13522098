import React, { useState, useEffect } from "react";
import { Coordinates, Plot, Line, Mafs, Point, Theme, useMovablePoint, useStopwatch, vec, Text } from "mafs";
import { easeInOutCubic } from "js-easing-functions";

function getStepLine(collection, length, layer, t) {
  if (length === 2) {
    const newPoint = [vec.scale(vec.add(vec.scale(collection[layer][0], 1 - t), vec.scale(collection[layer][1], t)), 1)];
    collection.push(newPoint);
  } else {
    const newPoint = [];
    for (let i = 1; i < collection[layer].length; i++) {
      newPoint.push(vec.scale(vec.add(vec.scale(collection[layer][i - 1], 1 - t), vec.scale(collection[layer][i], t)), 1));
    }
    collection.push(newPoint);
    getStepLine(collection, length - 1, layer + 1, t);
  }
}

function findCurve(points, length, t) {
  if (length === 2) {
    return vec.scale(vec.add(vec.scale(points[0], 1 - t), vec.scale(points[1], t)), 1);
  }
  const tmpoint = [];
  for (let i = 1; i < points.length; i++) {
    tmpoint.push(vec.scale(vec.add(vec.scale(points[i - 1], 1 - t), vec.scale(points[i], t)), 1));
  }
  return findCurve(tmpoint, length - 1, t);
}

function inNestedPairs(collection) {
  const nestedPairs = [];
  let limit = collection.length - 1;
  for (let i = 0; i < limit; i++) {
    for (let j = 1; j < collection[i].length; j++) {
      nestedPairs.push([collection[i][j - 1], collection[i][j]]);
    }
  }
  return nestedPairs;
}

function getAllPoints(collection) {
  const points = [];
  for (let i = 0; i < collection.length-1; i++) {
    for (let j = 0; j < collection[i].length; j++) {
      points.push(collection[i][j]);
    }
  }
  return points;
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
  return [min_x-0.1*dx, max_x+0.1*dx, min_y-0.1*dy, max_y+0.1*dy];
}

export default function BezierCurves({ data }) {
  const [t, setT] = useState(0.5);
  const [viewContent, setViewContent] = useState([]);
  const opacity = 1 - (2 * t - 1) ** 6;

  const movablePoints = data.map(([x, y]) => useMovablePoint([x, y]));

  const corner = movablePoints.map(point => point.point);
  const collection = [corner];

  getStepLine(collection, movablePoints.length, 0, t);

  const duration = 2;
  const { time, start } = useStopwatch({
    endTime: duration,
  });
  useEffect(() => {
    setTimeout(() => start(), 1500);
  }, [start]);

  useEffect(() => {
    setT(easeInOutCubic(time, 0, 1, duration));
  }, [time]);

  useEffect(() => {
    setViewContent(getMainView(data)); // Move viewContent logic into useEffect
  }, [data]);

  function drawAllLine(collection, color, customOpacity = opacity * 0.5) {
    return inNestedPairs(collection).map(([p1, p2], index) => (
      <Line.Segment
        key={index}
        point1={p1}
        point2={p2}
        opacity={customOpacity}
        color={color}
      />
    ));
  }

  function pointPosition(color, size) {
    return movablePoints.map(point => (
      <Text 
        x={point.x}
        y={point.y}
        color={color}
        size={size}
        attach="w"
        attachDistance={15}
      >
        ({point.x.toFixed(1)}, {point.y.toFixed(1)})
      </Text>
    ));
  }

  function drawPoints(rad, color) {
    return movablePoints.map((point, index) => (
      <Point
        key={index}
        x={point.x}
        y={point.y}
        svgCircleProps={{ r: rad }}
        color={color}
      />
    ));
  }

  function drawMainPoints(collection) {
    return <Point 
        key = {-1}
        x={collection[collection.length-1][0][0]}
        y={collection[collection.length-1][0][1]}
        color={Theme.red}
        svgCircleProps={{r:6}}
    />
  }

  function drawAllPoints(collection) {
    return getAllPoints(collection).map((point, index) => (
      <Point
        key={index}
        x={point[0]}
        y={point[1]}
        color={Theme.green}
        svgCircleProps={{r:3}}
      />
    ))
  }

  return (
    <div id="bjir">
      <p>Plot</p>
      <Mafs viewBox={{ x: [viewContent[0], viewContent[1]], y: [viewContent[2], viewContent[3]] }} zoom={{ min: 0.001, max: 5 }}>
        <Coordinates.Cartesian
          xAxis={{ lines: (viewContent[1]-viewContent[0])/20, labels: false, axis: false }}
          yAxis={{ lines: (viewContent[1]-viewContent[0])/20,labels: false, axis: false }}
        />

        {drawAllLine(
          collection,
          Theme.violet,
          0.5
        )}

        {drawAllPoints(collection)}

        <Plot.Parametric
          t={[0, t]}
          weight={3}
          color={Theme.red}
          xy={(t) =>
            findCurve(corner, movablePoints.length, t)
          }
        />

        <Plot.Parametric
          t={[1, t]}
          weight={3}
          opacity={0.5}
          style="dashed"
          xy={(t) =>
            findCurve(corner, movablePoints.length, t)
          }
        />

        {drawMainPoints(collection)}

        {pointPosition(Theme.foreground, 12)}

        {/* {movablePoints.map(point => point.element)} */}
        {drawPoints(6,Theme.indigo)}
      </Mafs>
      <br/>

      <div className="p-4 border-gray-700 border-t bg-black text-white">
        <input
          type="range"
          min={0}
          max={1}
          step={0.005}
          value={t}
          onChange={(event) => setT(+event.target.value)}
        />
      </div>
    </div>
  );
}
