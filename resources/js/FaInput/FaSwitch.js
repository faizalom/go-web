const FaSwitch = (props) => {

    const inputClasses = props.hasError ? 'form-check-input is-invalid' : 'form-check-input';

    return (
        <div className="mb-3 row">
            <label htmlFor={props.id} className="col-sm-3 col-form-label">{props.label}{props.required && <span className="text-danger">*</span>}</label>
            <div className="col-sm-9 pt-2">
                <div className="form-check form-switch">
                    <input
                        name={props.id}
                        type="checkbox"
                        id={props.id}
                        value={props.value}
                        className={inputClasses}
                        onChange={props.valueChangeHandler}
                        onBlur={props.inputBlurHandler}
                        checked={props.checked}
                    />
                </div>
                {props.hasError && (
                    <div className="invalid-feedback" style={{ "display": "block" }}>{props.errorMessage}</div>
                )}
            </div>
        </div>
    );
}

export default FaSwitch
