import React, { useState, useEffect } from "react";
import { Coordinates, Plot, Line, Mafs, Point, Theme, useMovablePoint, useStopwatch, vec } from "mafs";
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

function inPairs(arr) {
  const pairs = [];
  for (let i = 0; i < arr.length - 1; i++) {
    pairs.push([arr[i], arr[i + 1]]);
  }
  return pairs;
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

export default function BezierCurves({ data }) {
  const [t, setT] = useState(0.5);
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
    setTimeout(() => start(), 100);
  }, [start]);
  useEffect(() => {
    setT(easeInOutCubic(time, 0, 1, duration));
  }, [time]);

  function drawLineSegments(pointPath, color, customOpacity = opacity * 0.5) {
    return inPairs(pointPath).map(([p1, p2], index) => (
      <Line.Segment
        key={index}
        point1={p1}
        point2={p2}
        opacity={customOpacity}
        color={color}
      />
    ));
  }

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

  function drawPoints(points, color) {
    return points.map((point, index) => (
      <Point
        key={index}
        x={point[0]}
        y={point[1]}
        color={color}
        opacity={opacity}
      />
    ));
  }

  function drawMainPoints(collection) {
    return <Point 
        key = {-1}
        x={collection[collection.length-1][0][0]}
        y={collection[collection.length-1][0][1]}
        color={Theme.red}
        opacity={opacity}
    />
  }

  function drawAllPoints(collection) {
    return getAllPoints(collection).map((point, index) => (
      <Point
        key={index}
        x={point[0]}
        y={point[1]}
        color={Theme.green}
        opacity={opacity}
      />
    ))
  }

  return (
    <div id="kontol">
      <Mafs viewBox={{ x: [-5, 5], y: [-4, 4] }} zoom={{ min: 0.001, max: 5 }}>
        <Coordinates.Cartesian
          xAxis={{ labels: false, axis: false }}
          yAxis={{ labels: false, axis: false }}
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

        {movablePoints.map(point => point.element)}
      </Mafs>

      <div className="p-4 border-gray-700 border-t bg-black text-white">
        <span className="font-display">t =</span>{" "}
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
