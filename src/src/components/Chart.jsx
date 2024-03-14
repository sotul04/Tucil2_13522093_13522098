import { Mafs, Line, Coordinates, useMovablePoint } from "mafs";
import { Fragment,useState } from "react";
export default function Chart({data}){
    const beziercurve = () => {
        return(
            <>
                {data.slice(0,data.length - 1).map((_,index) => {
                    let point1 = useMovablePoint(data[index])
                    let point2 = useMovablePoint(data[index+1])
                    return(
                        <Fragment key={index}>
                            <Line.Segment
                                point1={point1.point}
                                point2={point2.point}
                            />
                        </Fragment>
                    )
                })}
            </>
        )
    }
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
                zoom={{min: 0.1, max: 5}}
                viewBox={{
                    x: [handleMin(data,'x'),handleMax(data,'x')],
                    y: [handleMin(data,'y'),handleMax(data,'y')],
                }} 
                    preserveAspectRatio={false}>
                <Coordinates.Cartesian
                />
                    {beziercurve()}
            </Mafs>
        </>
    );
}
