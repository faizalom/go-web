import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import useHttp from "../hooks/use-http.js";
import classes from './Staff.module.css';

function Market(props) {
    const [pairs, setPairs] = useState([]);
    const { isLoading, sendRequest } = useHttp();
    const [refresh, setRefresh] = useState(0);

    useEffect(() => {
        const transformStaff = (data) => {
            if (data === null) {
                setPairs({})
                return
            }
            setPairs(data)
        };

        sendRequest({ "url": "/market" }, transformStaff)
    }, [sendRequest, refresh])

    const refreshHandler = () => {
        const d = new Date();
        setRefresh(d.getTime());
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
                            <th className="text-center">#</th>
                            <th>Coin</th>
                            <th>Coin</th>
                            <th>LH</th>
                            <th>Now</th>
                            <th>24H</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {Object.keys(pairs).map((key) => (
                            <tr key={pairs[key].id}>
                                <td className="text-center">{parseInt(key) + 1}</td>
                                <td>
                                    <div className="avatar avatar-md">
                                        <img className="avatar-img" src={"https://cdn.coindcx.com/static/coins/" + pairs[key].market.replace("USDT", "").toLowerCase() + ".svg"} alt={pairs[key].market.replace("USDT", "")} />
                                        <span className="avatar-status bg-success" />
                                    </div>
                                </td>
                                <td>
                                    <div>{pairs[key].market.replace("USDT", "")}</div>
                                    <div className="small text-medium-emphasis"><span>{pairs[key].ask} ASK</span> | {pairs[key].bid} BID</div>
                                </td>
                                <td>
                                    <div>
                                        <div className="clearfix">
                                            <div className="float-start">
                                                <div className="fw-semibold">{pairs[key].LowHighPer.toFixed(2)}%</div>
                                            </div>
                                            <div className="float-end"><small className="text-medium-emphasis">{pairs[key].low} - {pairs[key].high}</small></div>
                                        </div>
                                        <div className="progress progress-thin">
                                            <div className="progress-bar bg-success" role="progressbar" style={{ width: pairs[key].LowHighPer.toFixed(2) + '%' }} aria-valuenow={pairs[key].LowHighPer.toFixed(2)} aria-valuemin={0} aria-valuemax={100} />
                                        </div>
                                    </div>
                                </td>
                                <td>
                                    <div className="small text-medium-emphasis">{pairs[key].last_price}</div>
                                    <div className="fw-semibold">{pairs[key].LowNowMargin.toFixed(4)}%</div>
                                </td>
                                <td>{pairs[key].change_24_hour}</td>
                                <td className="text-right">
                                    <Link to={"/u/staff/edit/" + pairs[key].id} className="mx-3"><i className="fas fa-edit fa-lg"></i></Link>
                                </td>
                            </tr>
                        ))}
                        {(Object.keys(pairs).length == 0 && !isLoading) ? <tr><td colSpan={8} className="text-center">No data to show</td></tr> : ''}
                    </tbody>
                </table>
            </div>
        </div>
    );
}

export default Market;