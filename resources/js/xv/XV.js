import { useState, useRef, useCallback, useEffect } from "react";
import { Link } from "react-router-dom";
import useHttp from "../hooks/use-http.js";
import 'react-toastify/dist/ReactToastify.css';

function XV(props) {
    const [pageNum, setPageNum] = useState(1);
    const [hasMore, sethasMore] = useState(1);
    const [sxv, setSxv] = useState([]);
    const { isLoading, sendRequest } = useHttp();

    const observer = useRef();
    const lastXVElementRef = useCallback((node) => {
        if (isLoading) return;
        if (observer.current) observer.current.disconnect();
        observer.current = new IntersectionObserver((entries) => {
            if (entries[0].isIntersecting && hasMore) {
                setPageNum((prev) => prev + 1);
            }
        });
        if (node) observer.current.observe(node);
    },
        [isLoading, hasMore]
    );

    useEffect(() => {
        const transformStaff = (data) => {
            setSxv((prev) => {
                return [...new Set([...prev, ...data])];
            });
            sethasMore(parseInt(Math.random() * 100))
        };

        let page = parseInt(Math.random() * 100);
        sendRequest({ "url": "/xv?limit=50&sort=-1&page=" + page }, transformStaff)
    }, [sendRequest, pageNum])

    const refreshHandler = () => {
        const d = new Date();
        setstaffDeleteId(d.getTime());
    }

    // const handleChange = (e) => {
    //     setQuery(e.target.value);
    //     setPageNum(1);
    // };

    return (
        <div className="card mb-4">
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
                <div className="card-group">

                    {Object.keys(sxv).map((key) => {
                        if (sxv.length === parseInt(key) + 1) {
                            return (<div className="card" style={{ "flex": "unset", "width": "200px" }} key={sxv[key].video_id} ref={lastXVElementRef}>
                                <img src={sxv[key].img.thumbs} className="card-img-top" alt={sxv[key].video_id} title={sxv[key].video_id} />
                                <div className="card-img-overlay p-1" style={{ "height": "30px" }}>
                                    <small className="badge rounded-pill bg-primary float-end opacity-75 fw-light">{sxv[key].uploader.name}</small>
                                </div>
                                <div className="card-body p-2">
                                    <p className="card-text text-truncate" title={sxv[key].title}><small><b>{sxv[key].title}</b></small></p>
                                    {/* <p className="card-text">This is a wider card with support</p> */}
                                </div>
                                <ul className="list-group list-group-flush">
                                    <li className="list-group-item p-1">
                                        <span className="badge bg-primary rounded-0">{sxv[key].category !== null ? sxv[key].category.join(', ') : "null"}</span>
                                        <span className="badge bg-info rounded-0">{sxv[key].subcategory !== null ? sxv[key].subcategory.join(', ') : "null"}</span>
                                    </li>
                                    <li className="list-group-item p-1 text-truncate" title={sxv[key].tags !== null ? sxv[key].tags.join(', ') : "null"}>
                                        <small>{sxv[key].tags !== null ? sxv[key].tags.join(', ') : "&nbsp;"}</small>
                                    </li>
                                </ul>
                                <div className="card-footer bg-info bg-gradient p-1">
                                    <a href={"https://www.xvideos5.com/video" + sxv[key].video_id + "/" + sxv[key].urlName} target="_blank" className="btn btn-primary btn-sm">
                                        <i className="fas fa-globe"></i>
                                    </a>
                                </div>
                            </div>)
                        } else {
                            return (<div className="card" style={{ "flex": "unset", "width": "200px" }} key={sxv[key].video_id}>
                                <img src={sxv[key].img.thumbs} className="card-img-top" alt={sxv[key].video_id} title={sxv[key].video_id} />
                                <div className="card-img-overlay p-1" style={{ "height": "30px" }}>
                                    <small className="badge rounded-pill bg-primary float-end opacity-75 fw-light">{sxv[key].uploader.name}</small>
                                </div>
                                <div className="card-body p-2">
                                    <p className="card-text text-truncate" title={sxv[key].title}><small><b>{sxv[key].title}</b></small></p>
                                    {/* <p className="card-text">This is a wider card with support</p> */}
                                </div>
                                <ul className="list-group list-group-flush">
                                    <li className="list-group-item p-1">
                                        <span className="badge bg-primary rounded-0">{sxv[key].category !== null ? sxv[key].category.join(', ') : "null"}</span>
                                        <span className="badge bg-info rounded-0">{sxv[key].subcategory !== null ? sxv[key].subcategory.join(', ') : "null"}</span>
                                    </li>
                                    <li className="list-group-item p-1 text-truncate" title={sxv[key].tags !== null ? sxv[key].tags.join(', ') : "null"}>
                                        <small>{sxv[key].tags !== null ? sxv[key].tags.join(', ') : "&nbsp;"}</small>
                                    </li>
                                </ul>
                                <div className="card-footer bg-info bg-gradient p-1">
                                    <a href={"https://www.xvideos5.com/video" + sxv[key].video_id + "/" + sxv[key].urlName} target="_blank" className="btn btn-primary btn-sm">
                                        <i className="fas fa-globe"></i>
                                    </a>
                                </div>
                            </div>)
                        }
                    })}
                </div>
            </div>
        </div>
    );
}

export default XV;