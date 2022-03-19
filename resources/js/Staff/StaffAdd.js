import { useEffect } from 'react';
import { useNavigate, useMatch } from 'react-router';
import FaCard from '../Fa/FaCard';
import FaInput from '../FaInput/FaInput';
import useHttp from '../hooks/use-http';
import useInput from '../hooks/use-input';
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const StaffAdd = (props) => {

    const { isLoading, sendRequest } = useHttp();
    const match = useMatch("u/staff/edit/:id")

    const history = useNavigate()
    //let toastId = 0;

    let lastNameErrorMessage = "Name must not be empty."
    let inputs = []
    inputs['memberCode'] = useInput((value) => value.trim() !== '');
    inputs['firstName'] = useInput((value) => value.trim() !== '');
    inputs['lastName'] = useInput(() => true);
    inputs['email'] = useInput((value) => value.includes('@'));
    inputs['username'] = useInput((value) => value.trim() !== '');
    inputs['mobile'] = useInput((value) => value.trim() !== '');
    inputs['address_1'] = useInput((value) => value.trim() !== '');
    inputs['address_2'] = useInput(() => true);
    inputs['city'] = useInput((value) => value.trim() !== '');
    inputs['postal_code'] = useInput((value) => value.trim() !== '');
    inputs['dob'] = useInput((value) => value.trim() !== '');

    let minDob = new Date();
    minDob.setFullYear(minDob.getFullYear() - 18);
    minDob = minDob.getFullYear() + "-" + ('0' + (minDob.getMonth() + 1)).slice(-2) + "-" + ('0' + minDob.getDate()).slice(-2);

    let formIsValid = true;
    for (let i in inputs) {
        if (!inputs[i].isValid) {
            formIsValid = false;
        }
    }

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
                "url": `https://mysapp.firebaseio.com/users/${match.params.id}.json`,
                method: "PATCH",
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

    let title="Add Staff"
    if (match) {
        title = "Edit Staff"
    }

    return (
        <FaCard color="info" title={title}>
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
                                errorMessage={"Name must not be empty."}
                                required
                            />
                            <FaInput type='text'
                                label="Last Name"
                                id='lastName'
                                {...inputs['lastName']}
                                errorMessage={lastNameErrorMessage}
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
                                label="Login ID"
                                id='username'
                                {...inputs['username']}
                                errorMessage={"Please enter a valid Login ID."}
                                required
                            />
                        </div>
                        <div className="col-md-6 pl-4">
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
                            <FaInput type='date'
                                label="Date of birth"
                                id='dob'
                                {...inputs['dob']}
                                errorMessage={"Date of birth must not be empty."}
                                required
                                max={minDob}
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
