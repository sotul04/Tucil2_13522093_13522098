import { Mafs, Line, Theme, Polyline, Coordinates, useMovablePoint } from "mafs";
export default function Chart({data}){

    function getDataPoint(){
        let newDataPoints = Array.from({length: data.length},(_,index) =>
        (
            useMovablePoint(data[index]).point
        ))

        console.log('data : ',data);
        console.log('newDataPoints : ',newDataPoints);
        return newDataPoints
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
                preserveAspectRatio={false}
            >
                <Coordinates.Cartesian/>
                <Polyline 
                    points={data}
                    // points={getDataPoint()}
                    color={Theme.blue}
                />
                {getDataPoint()}    
            </Mafs>
        </>
    );
}
