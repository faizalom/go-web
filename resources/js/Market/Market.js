import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import useHttp from "../hooks/use-http.js";
import classes from './Coins.module.css';
import { useMatch } from 'react-router';

function Market(props) {
    const [pairs, setPairs] = useState([]);
    const [ sortColumn, setSortColumn ] = useState(null);
    const [ sortDir, setSortDir ] = useState(1);
    const { isLoading, sendRequest } = useHttp();
    const [refreshCoin, setRefreshCoin] = useState(0);

    const match = useMatch("u/great-trade")

    useEffect(() => {
        const transformCoins = (data) => {
            if (data === null) {
                setPairs({})
                return
            }
            setPairs(data)
        };

        let sort = "?sort_dir=" + sortDir
        if (sortColumn !== null) {
            sort += "&sort_column=" + sortColumn
        }

        let url = "/market"
        if (match) {
            url = "/great-trade"
        }
        sendRequest({ "url": url + sort }, transformCoins)
    }, [sendRequest, refreshCoin, sortColumn, sortDir])

    const refreshHandler = () => {
        setSortColumn(null)
        setSortDir(1)
        const d = new Date();
        setRefreshCoin(d.getTime());
    }

    const sortHandler = (column) => {
        setSortColumn((prev) => {
            if (prev == column) {
                setSortColumn((prev) => {
                    if (prev == column) {
                        setSortDir((pre) => pre == 1 ? 0 : 1)
                    } else {
                        setSortDir(1)
                    }
                    return column
                })
            }
            return column
        })
    }

    const cardheight = isLoading ? { "height": "80vh" } : {};

    return (
        <div className="card mb-4" style={cardheight}>
            {isLoading &&
                <div className={classes.loaderRoot + " d-flex justify-content-center align-self-center"} >
                    <div className={classes.loader + " align-self-center"}></div>
                </div>
            }
            <div className="card-header">Coins</div>
            <div className="card-header">
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
                            <th style={{"width":"50px"}}>#</th>
                            <th style={{"width":"50px"}}></th>
                            <th onClick={() => sortHandler("market")} className="pointer">
                                Coin {sortColumn != "market" && <span className="fas fa-sort"/>}
                                {sortColumn == "market" && sortDir == 1 && <span className="fas fa-sort-up"/>}
                                {sortColumn == "market" && sortDir == 0 && <span className="fas fa-sort-down"/>}
                            </th>
                            <th onClick={() => sortHandler("change_24_hour")} className="pointer">
                                24 {sortColumn != "change_24_hour" && <span className="fas fa-sort"/>}
                                {sortColumn == "change_24_hour" && sortDir == 1 && <span className="fas fa-sort-up"/>}
                                {sortColumn == "change_24_hour" && sortDir == 0 && <span className="fas fa-sort-down"/>}
                            </th>
                            <th>E-Mail</th>
                            <th>Username</th>
                            <th>City</th>
                            <th>#</th>
                        </tr>
                    </thead>
                    <tbody>
                        {Object.keys(pairs).map((key) => (
                            <tr key={pairs[key].market}>
                                <td>{parseInt(key) + 1}</td>
                                <td>
                                    <div className="avatar avatar-md">
                                        <img className="avatar-img" src={"https://cdn.coindcx.com/static/coins/" + pairs[key].Coin.toLowerCase() + ".svg"} alt={pairs[key].Coin} />
                                        <span className="avatar-status bg-success" />
                                    </div>
                                </td>
                                <td>
                                    <div>{pairs[key].Coin}</div>
                                    <div className="small text-medium-emphasis"><span>{pairs[key].ask} ASK</span> | {pairs[key].bid} BID</div>
                                </td>
                                <td>
                                    {pairs[key].change_24_hour < 0 ?
                                    (<small className="text-danger mr-1"><i className="fas fa-arrow-down"></i> {pairs[key].change_24_hour}%</small>) :
                                    (<small className="text-success mr-1"><i className="fas fa-arrow-up"></i> {pairs[key].change_24_hour}%</small>)
                                    }
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
                                <td className="text-right">
                                    <a target="_blank" href={"https://coindcx.com/trade/" + pairs[key].market} className="mx-1"><i class="fas fa-chart-line fa-lg"></i></a>
                                    <Link to={"/u/staff/edit/" + pairs[key].id} className="mx-1"><i className="fas fa-edit fa-lg"></i></Link>
                                    <button className="btn btn-link p-0" onClick={(e) => deleteCoins(pairs[key].id)}><i className="fas fa-trash fa-lg"></i></button>
                                </td>
                            </tr>
                        ))}
                        {(Object.keys(pairs).length == 0 && !isLoading) ? <tr><td colSpan={8} className="text-center">No data to show</td></tr> : null}
                    </tbody>
                </table>
            </div>
        </div>
    );
}

export default Market;