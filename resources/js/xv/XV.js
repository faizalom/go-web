import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import useHttp from "../hooks/use-http.js";
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import classes from './XV.module.css';

function XV(props) {
    const [sxv, setSxv] = useState([]);
    const { isLoading, sendRequest } = useHttp();
    const [staffDeleteId, setstaffDeleteId] = useState(0);


    useEffect(() => {
        window.addEventListener('scroll', onScroll, { passive: true });
    }, [])

    useEffect(() => {
        const transformStaff = (data) => {
            // if (data === null) {
            //     setSxv({})
            //     return
            // }
            setSxv([...sxv, ...data])
        };

        let page = parseInt(Math.random() * 100);
        sendRequest({ "url": "/xv?sort=-1&page=" + page }, transformStaff)
    }, [sendRequest, staffDeleteId])

    const deleteStaff = (key) => {
        sendRequest({
            "url": `https://mysapp.firebaseio.com/sxv/${key}.json`,
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
    const onScroll = (e) => {

        // const currentScrollY = e.target.scrollTop;
        // if (prevScrollY.current < currentScrollY && goingUp) {
        //     //setGoingUp(false);
        // }
        // if (prevScrollY.current > currentScrollY && !goingUp) {
        //     //setGoingUp(true);
        // }
        // prevScrollY.current = currentScrollY;
        //console.log(e.path[1].visualViewport.height);
        // console.log(e.path[1].scrollY);
        // let body = document.body;
        // console.log(body.offsetHeight);
        //console.log(e.path[1]);

        console.log(e.path[1].outerHeight, e.path[1].pageYOffset)

        if (e.path[1].outerHeight < e.path[1].pageYOffset) {
            console.log(e.path[1].outerHeight, e.path[1].pageYOffset)

            const d = new Date();
            setstaffDeleteId(d.getTime());
        }

        // var body = document.body,
        // html = document.documentElement;

        // var height = Math.max( body.scrollHeight, body.offsetHeight,
        //                html.clientHeight, html.scrollHeight, html.offsetHeight );

        // console.log(height);
        // console.log(e.path[1].offsetHeight);
        //console.log(body.offsetHeight - e.path[1].scrollY)
    };

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
                <div className="card-group">
                    {Object.keys(sxv).map((key) => (
                        <div className="card" style={{ "flex": "unset", "width": "200px" }} key={sxv[key].video_id}>
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
                                <li className="list-group-item">A third item</li>
                            </ul>
                            <div className="card-footer bg-info bg-gradient p-1">
                                <a href={"https://www.xvideos5.com/video" + sxv[key].video_id + "/" + sxv[key].urlName} target="_blank" className="btn btn-primary btn-sm">
                                    <i className="fas fa-globe"></i>
                                </a>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
}

export default XV;