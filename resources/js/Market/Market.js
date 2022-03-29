import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import useHttp from "../hooks/use-http.js";
import classes from './Coins.module.css';
import { useMatch } from 'react-router';

function Market(props) {
    const [coins, setCoins] = useState([]);
    const [ sortColumn, setSortColumn ] = useState(null);
    const [ sortDir, setSortDir ] = useState(1);
    const { isLoading, sendRequest } = useHttp();
    const [refreshCoin, setRefreshCoin] = useState(0);

    const match = useMatch("u/great-trade")

    useEffect(() => {
        const transformCoins = (data) => {
            if (data === null) {
                setCoins({})
                return
            }
            setCoins(data)
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
                        {Object.keys(coins).map((key) => (
                            <tr key={coins[key].market}>
                                <td>{parseInt(key) + 1}</td>
                                <td>
                                    <div className="avatar avatar-md">
                                        <img className="avatar-img" src={"https://cdn.coindcx.com/static/coins/" + coins[key].Coin.toLowerCase() + ".svg"} alt={coins[key].Coin} /><span className="avatar-status bg-success" />
                                    </div>
                                </td>
                                <td>
                                    <div>{coins[key].Coin}</div>
                                    <div className="small text-medium-emphasis"><span>New</span> | Registered: Jan 1, 2020</div>
                                </td>
                                <td>
                                    {coins[key].change_24_hour < 0 ?
                                    (<small className="text-danger mr-1"><i className="fas fa-arrow-down"></i> {coins[key].change_24_hour}%</small>) :
                                    (<small className="text-success mr-1"><i className="fas fa-arrow-up"></i> {coins[key].change_24_hour}%</small>)
                                    }
                                </td>
                                <td>{coins[key].email}</td>
                                <td>{coins[key].username}</td>
                                <td>{coins[key].city}</td>
                                <td className="text-right">
                                    <a target="_blank" href={"https://coindcx.com/trade/" + coins[key].market} className="mx-1"><i class="fas fa-chart-line fa-lg"></i></a>
                                    <Link to={"/u/staff/edit/" + coins[key].id} className="mx-1"><i className="fas fa-edit fa-lg"></i></Link>
                                    <button className="btn btn-link p-0" onClick={(e) => deleteCoins(coins[key].id)}><i className="fas fa-trash fa-lg"></i></button>
                                </td>
                            </tr>
                        ))}
                        {(Object.keys(coins).length == 0 && !isLoading) ? <tr><td colSpan={8} className="text-center">No data to show</td></tr> : null}
                    </tbody>
                </table>
            </div>
        </div>
    );
}

export default Market;