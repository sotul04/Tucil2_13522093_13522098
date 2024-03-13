import { Mafs, Line, Coordinates, useMovablePoint } from "mafs";
import { Fragment } from "react";
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
                            />``
                        </Fragment>
                    )
                })}
            </>
        )
    }
    
    return(
        <Mafs viewBox={{x: [-10,10], y: [-10,10]}} preserveAspectRatio={false}>
            <Coordinates.Cartesian/>
                {beziercurve()}
        </Mafs>
    );
}
