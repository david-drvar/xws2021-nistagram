import Loader from "react-loader-spinner";
import "react-loader-spinner/dist/loader/css/react-spinner-loader.css";

const defaultConfig = {
    height: 100,
    width: 100,
    color: "#e03c67",
    secondaryColor: "#da3394",
    type: "ThreeDots"
}

const Spinner = (props) => {
    const { height, width, color, secondaryColor, type } = props;

    return (
        <Loader
            type={type ? type : defaultConfig.type}
            color={color ? color : defaultConfig.color}
            secondaryColor={secondaryColor ? secondaryColor : defaultConfig.secondaryColor}
            height={height ? height : defaultConfig.height}
            width={width ? width : defaultConfig.width}
        />
    )
}

export default Spinner;