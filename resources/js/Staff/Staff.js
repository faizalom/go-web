import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import useHttp from "../hooks/use-http.js";
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import classes from './Staff.module.css';

function Staff(props) {
    const [users, setUsers] = useState([]);
    const { isLoading, sendRequest } = useHttp();
    const [staffDeleteId, setstaffDeleteId] = useState(0);

    useEffect(() => {
        const transformStaff = (data) => {
            if (data === null) {
                setUsers({})
                return
            }
            setUsers(data)
        };

        sendRequest({ "url": "/api/staff" }, transformStaff)
    }, [sendRequest, staffDeleteId])

    const deleteStaff = (key) => {
        sendRequest({
            "url": `https://mysapp.firebaseio.com/users/${key}.json`,
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            }
        }, () => {
            toast.success('Staff deleted successfully', { theme: "colored" });
            setstaffDeleteId(key)
        });
    };

    const refreshHandler = () => {
        const d = new Date();
        setstaffDeleteId(d.getTime());
    }

    const cardheight = isLoading ? { "minHeight": "40vh" } : {};

    return (
        <div className="card mb-4" style={cardheight}>
            {isLoading &&
                <div className={classes.loaderRoot + " d-flex justify-content-center align-self-center"} >
                    <div className={classes.loader + " align-self-center"}></div>
                </div>
            }
            <div className="card-header">Staff</div>
            <div className="card-header">
                <Link to="/u/staff/add" className="btn btn-primary float-end"><i className="fas fa-plus" /> Add Staff</Link>
                <div className="input-group" style={{ width: "350px" }}>
                    <button className="btn btn-primary" tabIndex={0} type="button" onClick={refreshHandler}>
                        <i className="fas fa-sync" /> Refresh
                    </button>
                </div>
            </div>
            <div className="card-body table-responsive p-0 position-relative">
                <table className="table table-hover text-nowrap table-striped m-0">
                    <thead>
                        <tr>
                            <th className="text-center">
                                <svg className="icon">
                                    <use xlinkHref="vendors/@coreui/icons/svg/free.svg#cil-people" />
                                </svg>
                            </th>
                            <th>Staff Code</th>
                            <th>First Name</th>
                            <th>Last Name</th>
                            <th>E-Mail</th>
                            <th>Username</th>
                            <th>City</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {Object.keys(users).map((key) => (
                            <tr key={users[key].id}>
                                <td className="text-center">
                                    <div className="avatar avatar-md"><img className="avatar-img" src={users[key].profile_photo} /></div>
                                </td>
                                <td>{users[key].memberCode}</td>
                                <td>
                                    <div>{users[key].firstName}</div>
                                    {/* <div className="small text-medium-emphasis"><span>{users[key].memberCode}</span></div> */}
                                </td>
                                <td>{users[key].lastName}</td>
                                <td>{users[key].email}</td>
                                <td>{users[key].username}</td>
                                <td>{users[key].city}</td>
                                <td className="text-right">
                                    <Link to={"/u/staff/edit/" + users[key].id} className="mx-3"><i className="fas fa-edit fa-lg"></i></Link>
                                    <button className="btn btn-link p-0" onClick={(e) => deleteStaff(users[key].id)}><i className="fas fa-trash fa-lg"></i></button>
                                </td>
                            </tr>
                        ))}
                        {(Object.keys(users).length == 0 && !isLoading) ? <tr><td colSpan={8} className="text-center">No data to show</td></tr> : null}
                    </tbody>
                </table>
            </div>
        </div>
    );
}

export default Staff;