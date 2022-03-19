const FaInput = (props) => {

    const inputClasses = props.hasError ? 'form-control is-invalid' : 'form-control';

    return (
        <div className="mb-3 row">
            <label htmlFor={props.id} className="col-sm-3 col-form-label">{props.label}{props.required && <span className="text-danger">*</span>}</label>
            <div className="col-sm-9">
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
                />
                {props.hasError && (
                    <div className="invalid-feedback">{props.errorMessage}</div>
                )}
            </div>
        </div>
    );
}

const FaInputAjax = (props) => {
    let inputClasses = !props.hasError && props.isValid ? 'form-control is-valid' : 'form-control';
    inputClasses = props.hasError ? 'form-control is-invalid' : inputClasses;

    return (
        <div className="mb-3 row">
            <label htmlFor={props.id} className="col-sm-3 col-form-label">{props.label}{props.required && <span className="text-danger">*</span>}</label>
            <div className="col-sm-9" style={{"position": "relative"}}>
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
                />
                {props.isLoading && (
                    <div style={{"position": "absolute", "top": "7px", "right": "23px"}}><i className="fa fa-spinner fa-spin" aria-hidden="true" /></div>
                )}
                {props.hasError && (
                    <div className="invalid-feedback">{props.errorMessage}</div>
                )}
                {(!props.hasError && props.isValid) && (
                    <div className="valid-feedback">{props.errorMessage}</div>
                )}
            </div>
        </div>
    );
}

export default FaInput ;
export {FaInputAjax};
