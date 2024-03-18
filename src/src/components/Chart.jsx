import { Mafs, Theme, Polyline, Coordinates, } from "mafs";
import { useEffect } from "react";

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

function isPreferred(range, num) {
    return num%range == 0;
}

export default function Chart({data}){
    let mainView = getMainView(data);

    useEffect(() => {
        mainView = getMainView(data);
    }, [data]);

    function handleMin(data,status){
        if (status == 'x'){
            let min = data[0][0]
            for (let i = 0; i < data.length; i++){
                if (min > data[i][0]){
                    min = data[i][0]
                }
            }
            return min*1.1
        } else if (status == 'y'){
            let min = data[0][1]
            for (let i = 0; i < data.length; i++){
                if (min > data[i][1]){
                    min = data[i][1]
                }
            }
            return min*1.1
        }
    }
    function handleMax(data,status){
        if (status == 'x'){
            let max = data[0][0]
            for (let i = 0; i < data.length; i++){
                if (max < data[i][0]){
                    max = data[i][0]
                }
            }
            return max*1.1
        } else if (status == 'y'){
            let max = data[0][1]
            for (let i = 0; i < data.length; i++){
                if (max < data[i][1]){
                    max = data[i][1]
                }
            }
            return max*1.1
        }
    }
    return(
        <>
            <Mafs 
                zoom={{min: 0.7, max: 5}}
                viewBox={{
                    x: [mainView[0], mainView[1]],
                    y: [mainView[2], mainView[3]],
                }} 
                preserveAspectRatio={false}
            >
                <Coordinates.Cartesian
                    xAxis={{ lines: Math.ceil((mainView[1]-mainView[0])/10), labels: (n) => (isPreferred(Math.ceil((mainView[1]-mainView[0])/10), n) ? n.toFixed(0) : "")}}
                    yAxis={{ lines: Math.ceil((mainView[3]-mainView[2])/10), labels: (n) => (isPreferred(Math.ceil((mainView[3]-mainView[2])/10), n) ? n.toFixed(0) : "")}}
                />
                <Polyline 
                    points={data}
                    color={Theme.blue}
                />
            </Mafs>
        </>
    );
}
