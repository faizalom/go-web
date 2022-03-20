import { useEffect } from 'react';
import { useNavigate, useMatch } from 'react-router';
import FaCard from '../Fa/FaCard';
import FaInput, { FaInputAjax } from '../FaInput/FaInput';
import useHttp from '../hooks/use-http';
import useInput from '../hooks/use-input';
import useInputAjax from '../hooks/use-inputAjax';
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import FaText from '@/FaInput/FaText';
import FaUserImg from '@/FaInput/FaUserImg';
import FaRadio from '@/FaInput/FaRadio';
import FaSelect from '@/FaInput/FaSelect';
import FaSwitch from '@/FaInput/FaSwitch';
import classes from './Staff.module.css';

function getBloodGroup() {
    var bloodGroup = []
    bloodGroup['A+'] = 'A+';
    bloodGroup['A-'] = 'A-';
    bloodGroup['B+'] = 'B+';
    bloodGroup['B-'] = 'B-';
    bloodGroup['0+'] = '0+';
    bloodGroup['0-'] = '0-';
    bloodGroup['AB+'] = 'AB+';
    bloodGroup['AB-'] = 'AB-';
    return bloodGroup;
}

const StaffAdd = (props) => {

    const { isLoading, sendRequest } = useHttp();
    const { isLoading: ajaxLoading, sendRequest: ajaxSendRequest } = useHttp();

    const match = useMatch("u/staff/edit/:id")
    const history = useNavigate()

    let inputs = []
    inputs['memberCode'] = useInput((value) => value.trim() !== '');
    inputs['firstName'] = useInput((value) => value.trim() !== '');
    inputs['lastName'] = useInput(() => true);
    inputs['email'] = useInput((value) => value.includes('@'));
    inputs['allow_login'] = useInput(() => true);
    inputs['username'] = useInputAjax();
    inputs['password'] = useInput((value) => {
        return value.trim() !== '' || inputs['allow_login'].value !== 'yes'
    });
    inputs['mobile'] = useInput((value) => value.trim() !== '');
    inputs['address_1'] = useInput((value) => value.trim() !== '');
    inputs['address_2'] = useInput(() => true);
    inputs['city'] = useInput((value) => value.trim() !== '');
    inputs['postal_code'] = useInput((value) => value.trim() !== '');
    inputs['dob'] = useInput((value) => value.trim() !== '');
    inputs['comments'] = useInput(() => true);
    inputs['blood_group'] = useInput((value) => value.trim() !== '');
    inputs['gender'] = useInput((value) => value.trim() !== '');
    inputs['profile_photo'] = useInput(() => true);

    let minDob = new Date();
    minDob.setFullYear(minDob.getFullYear() - 18);
    minDob = minDob.getFullYear() + "-" + ('0' + (minDob.getMonth() + 1)).slice(-2) + "-" + ('0' + minDob.getDate()).slice(-2);

    let formIsValid = true;
    for (let i in inputs) {
        if (!inputs[i].isValid) {
            formIsValid = false;
        }
    }

    const allowLogiChangeHandler = (event) => {
        if (event.target.checked) {
            inputs['allow_login'].setValue(event.target.value)
            inputs['username'].setIsValid(false)
            inputs['username'].setMessage('')
        } else {
            inputs['username'].setValue('')
            inputs['username'].setIsValid(true)
            inputs['password'].setValue('')
            inputs['allow_login'].setValue('')
        }
    };

    useEffect(() => {
        inputs['username'].setValue(inputs['username'].value.toLowerCase());
        const identifier = setTimeout(() => {
            inputs['username'].setIsValid(false)
            if (inputs['allow_login'].value != 'yes') {
                inputs['username'].setIsValid(true)
                return;
            }
            if (inputs['username'].value.length < 5) {
                inputs['username'].setLoading(false);
                inputs['username'].setIsValid(false);
                inputs['username'].setMessage("Login ID atleast 5 character");
            } else {
                inputs['username'].setLoading(true);
                ajaxSendRequest({
                    "url": "/api/staff-available",
                    method: "POST",
                    body: {'username': inputs['username'].value, 'id': match.params.id},
                    headers: {
                        "Content-Type": "application/json"
                    }
                }, (res) => {
                    inputs['username'].setLoading(false)
                    if (res.is_available) {
                        inputs['username'].setMessage(inputs['username'].value + " is available");
                        inputs['username'].setIsValid(true);
                    } else {
                        inputs['username'].setMessage(inputs['username'].value + " is already taken");
                        inputs['username'].setIsValid(false);
                    }
                })
            }
        }, 500);
        return () => {
          clearTimeout(identifier);
        };
    }, [inputs['username'].value]);

    useEffect(() => {
        const transformStaff = (staffData) => {
            for (let i in staffData) {
                if (i == 'id') continue
                inputs[i].setValue(staffData[i])
            }
        };

        if (match) {
            sendRequest({ "url": `/api/staff/${match.params.id}` }, transformStaff)
        }
    }, [match, sendRequest])

    const goToSTaff = (message, data) => {
        //toast.update(toastId, { isLoading: false });
        toast.success(message, { theme: "colored" });
        history("/u/staff")
    }

    const formSubmissionHandler = (event) => {
        event.preventDefault();
        //toastId = toast.loading("Please wait...")
        let staff = {};
        for (let i in inputs) {
            staff[i] = inputs[i].value;
            inputs[i].reset();
        }
        if (match) {
            sendRequest({
                "url": `/api/staff/${match.params.id}`,
                method: "PUT",
                body: staff,
                headers: {
                    "Content-Type": "application/json"
                }
            }, goToSTaff.bind(null, "Staff updated successfully"))
        } else {
            sendRequest({
                "url": "https://mysapp.firebaseio.com/users.json",
                method: "POST",
                body: staff,
                headers: {
                    "Content-Type": "application/json"
                }
            }, goToSTaff.bind(null, "Staff added successfully"))
        }
    };

    let title = "Add Staff"
    if (match) {
        title = "Edit Staff"
    }

    return (
        <FaCard color="info" title={title}>
            {isLoading &&
                <div className={classes.loaderRoot + " d-flex justify-content-center align-self-center"} >
                    <div className={classes.loader + " align-self-center"}></div>
                </div>
            }
            <form className="form-horizontal" onSubmit={formSubmissionHandler}>
                <div className="card-body">
                    <div className="row">
                        <div className="col-md-6 pr-4">
                            <FaInput type='text'
                                label="Staff Code"
                                id='memberCode'
                                {...inputs['memberCode']}
                                errorMessage={"Member Code must not be empty."}
                                required
                            />
                            <FaInput type='text'
                                label="First Name"
                                id='firstName'
                                {...inputs['firstName']}
                                errorMessage={"First name must not be empty."}
                                required
                            />
                            <FaInput type='text'
                                label="Last Name"
                                id='lastName'
                                {...inputs['lastName']}
                                errorMessage={"Last name must not be empty."}
                                hasError={inputs['lastName'].hasError}
                            />
                            <FaInput type='text'
                                label="E-mail"
                                id='email'
                                {...inputs['email']}
                                errorMessage={"Please enter a valid email."}
                                required
                            />
                            <FaInput type='text'
                                label="Mobile"
                                id='mobile'
                                {...inputs['mobile']}
                                errorMessage={"Mobile must not be empty."}
                                required
                            />
                            <FaInput type='text'
                                label="Address 1"
                                id='address_1'
                                {...inputs['address_1']}
                                errorMessage={"Address 1 must not be empty."}
                                required
                            />
                            <FaInput type='text'
                                label="Address 2"
                                id='address_2'
                                {...inputs['address_2']}
                                errorMessage={"Address 2 must not be empty."}
                            />
                            <FaInput type='text'
                                label="City"
                                id='city'
                                {...inputs['city']}
                                errorMessage={"City must not be empty."}
                                hasError={inputs['city'].hasError}
                                required
                            />
                            <FaInput type='text'
                                label="Postal Code"
                                id='postal_code'
                                {...inputs['postal_code']}
                                errorMessage={"Postal Code must not be empty."}
                                required
                            />
                        </div>
                        <div className="col-md-6 pl-4">
                            <FaUserImg {...inputs['profile_photo']} />
                            <FaRadio options={{ "Male": "Male", "Female": "Female" }}
                                label="Gender"
                                id='gender'
                                {...inputs['gender']}
                                errorMessage={"Please select gender."}
                                required
                            />
                            <FaSelect options={getBloodGroup()}
                                label="Blood Group"
                                id='blood_group'
                                {...inputs['blood_group']}
                                errorMessage={"Please choose blood group."}
                                required
                            />
                            <FaInput type='date'
                                label="Date of birth"
                                id='dob'
                                {...inputs['dob']}
                                errorMessage={"Date of birth must not be empty."}
                                required
                                max={minDob}
                            />
                            <FaSwitch
                                label="Allow Login"
                                id='allow_login'
                                {...inputs['allow_login']}
                                value="yes"
                                valueChangeHandler={allowLogiChangeHandler}
                                checked={inputs['allow_login'].value === 'yes'}
                            />
                            {inputs['allow_login'].value &&
                                (<FaInputAjax type='text'
                                    label="Login ID"
                                    id='username'
                                    {...inputs['username']}
                                    errorMessage={inputs['username'].message}
                                    required
                                />)}
                            {inputs['allow_login'].value &&
                                (<FaInput type='password'
                                    label="Password"
                                    id='password'
                                    {...inputs['password']}
                                    errorMessage={"Please enter a password."}
                                    required
                                />)}
                            <FaText type='text'
                                label="Comments"
                                id='comments'
                                {...inputs['comments']}
                            />
                        </div>
                    </div>
                </div>
                <div className="card-footer">
                    <button disabled={!formIsValid || isLoading} className="btn btn-primary float-end">{isLoading ? 'Processing...' : 'Save'}</button>
                    <button type="button" onClick={() => history(-1)} className="btn btn-primary"><i className="fas fa-long-arrow-alt-left"></i> Back</button>
                </div>
            </form>
        </FaCard>
    );
}

export default StaffAdd;
