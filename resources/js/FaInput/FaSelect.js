const FaSelect = (props) => {
    const inputClasses = props.hasError ? 'form-control is-invalid' : 'form-control';
    return (
        <div className="mb-3 row">
            <label htmlFor={props.id} className="col-sm-3 col-form-label">{props.label}{props.required && <span className="text-danger">*</span>}</label>
            <div className="col-sm-9">
                <select
                    className={inputClasses}
                    id={props.id}
                    onChange={props.valueChangeHandler}
                    onBlur={props.inputBlurHandler}
                    value={props.value}
                    required={props.required}
                >
                    <option value="">Select...</option>
                    {Object.keys(props.options).map((key) => (
                        <option value={key} key={key}>{props.options[key]}</option>
                    ))}
                </select>
                {props.hasError && (
                    <div className="invalid-feedback">{props.errorMessage}</div>
                )}
            </div>
        </div>
    );
}

export default FaSelect
