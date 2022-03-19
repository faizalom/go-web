const FaText = (props) => {

    const inputClasses = props.hasError ? 'form-control is-invalid' : 'form-control';

    return (
        <div className="mb-3 row">
            <label htmlFor={props.id} className="col-sm-3 col-form-label">{props.label}{props.required && <span className="text-danger">*</span>}</label>
            <div className="col-sm-9">
                <textarea
                    className={inputClasses}
                    type={props.type}
                    id={props.id}
                    onChange={props.valueChangeHandler}
                    onBlur={props.inputBlurHandler}
                    minLength={props.minLength}
                    required={props.required}
                    max={props.max}
                    value={props.value}
                >{props.value}</textarea>
                {props.hasError && (
                    <div className="invalid-feedback">{props.errorMessage}</div>
                )}
            </div>
        </div>
    );
}

export default FaText
