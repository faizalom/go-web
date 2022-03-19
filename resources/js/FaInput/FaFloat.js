const FaFloat = (props) => {

    const inputClasses = props.hasError ? 'form-control is-invalid' : 'form-control';

    return (
        <div className="form-floating mb-3">
            <input
                className={inputClasses}
                type={props.type}
                id={props.id}
                onChange={props.valueChangeHandler}
                onBlur={props.inputBlurHandler}
                value={props.value}
                minLength={props.minLength}
                required={props.required}
                max={props.max}
                placeholder={props.label}
            />
            <label htmlFor={props.id}>{props.label}{props.required && <span className="text-danger">*</span>}</label>
            {props.hasError && (
                <div className="invalid-feedback">{props.errorMessage}</div>
            )}
        </div>
    );
}

export default FaFloat;
