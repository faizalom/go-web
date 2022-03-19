const FaCard = (props) => {
    return (
        <div className={"card card-info card-outline"}>
            {props.title &&
                <div className="card-header"><b>{props.title}</b></div>
            }
            {props.children}
        </div>
    )
}

export default FaCard