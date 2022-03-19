const FaRadio = (props) => {
    const inputClasses = props.hasError ? 'form-check-input is-invalid' : 'form-check-input';
    return (
        <div className="mb-3 row">
            <label htmlFor={props.id} className="col-sm-3 col-form-label">{props.label}{props.required && <span className="text-danger">*</span>}</label>
            <div className="col-sm-9 pt-2">
                {Object.keys(props.options).map((key) => (
                    <div className="form-check form-check-inline" key={key}>
                        <input
                            name={props.id}
                            type="radio"
                            id={props.id + '-' + props.options[key]}
                            value={props.options[key]}
                            className={inputClasses}
                            onChange={props.valueChangeHandler}
                            onBlur={props.inputBlurHandler}
                            checked={props.value === props.options[key]}
                        />
                        <label className="form-check-label" htmlFor={props.id + '-' + props.options[key]}>{props.options[key]}</label>
                    </div>
                ))}
                {props.hasError && (
                    <div className="invalid-feedback" style={{"display": "block"}}>{props.errorMessage}</div>
                )}
            </div>
        </div>
    );
}

export default FaRadio
