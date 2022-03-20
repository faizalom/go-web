import { useState, useCallback } from 'react';
import { Button, ButtonGroup, Modal } from "react-bootstrap";
import classes from './FaUserImg.module.css';
import Cropper from 'react-easy-crop';
import getCroppedImg from './cropImage';

const FaUserImg = (props) => {
    const [show, setShow] = useState(false);
    if (props.value === '') {
        props.setValue("/images/profilePhoto.png")
    }
    const [selectedImage, setSelectedImage] = useState(props.value);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);

    const [crop, setCrop] = useState({ x: 0, y: 0 })
    const [rotation, setRotation] = useState(0)
    const [zoom, setZoom] = useState(1)
    const [croppedAreaPixels, setCroppedAreaPixels] = useState(null)

    const onCropComplete = useCallback((croppedArea, croppedAreaPixels) => {
      setCroppedAreaPixels(croppedAreaPixels)
    }, [])

    const showCroppedImage = useCallback(async () => {
        try {
          const croppedImage = await getCroppedImg(
            selectedImage,
            croppedAreaPixels,
            rotation
          )
          props.setValue(croppedImage)
          handleClose()
        } catch (e) {
          console.error(e)
        }
    }, [croppedAreaPixels, rotation])

    const assignHandler = (event) => {
        var attachment = event.target.files[0];
        // fileType = attachment.type;
        // fileSize = attachment.size;
        // fileName = attachment.name;
        var imgReader = new FileReader();
        imgReader.readAsDataURL(attachment);
        imgReader.onload = function (e) {
            setSelectedImage(e.target.result);
        };
    }

    const inputClasses = props.hasError ? 'form-control is-invalid' : 'form-control';
    return (
        <>
            <div className="mb-3 row">
                <label htmlFor={props.id} className="col-sm-3 col-form-label">{props.label}{props.required && <span className="text-danger">*</span>}</label>
                <div className="col-sm-9 d-flex justify-content-center">
                    <div className={classes.userImage} onClick={handleShow}>
                        <img className="img-thumbnail rounded" id="user-photo-preview" src={props.value} alt="User profile picture" />
                        {props.hasError && (
                            <div className="invalid-feedback">{props.errorMessage}</div>
                        )}
                        <div className={classes.imgTxt}>Click to change image</div>
                    </div>
                </div>
            </div>
            <Modal show={show} onHide={handleClose}>
                <Modal.Body className={classes.cropContainer}>
                    <Cropper
                        image={selectedImage}
                        crop={crop}
                        zoom={zoom}
                        objectFit="contain"
                        aspect={1 / 1}
                        onCropChange={setCrop}
                        onCropComplete={onCropComplete}
                        onZoomChange={setZoom}
                    />
                </Modal.Body>
                <Modal.Footer className="p-1">
                    <Button className="me-auto" variant="secondary" onClick={handleClose} size="sm" >Close</Button>
                    <ButtonGroup className="m-auto" aria-label="Basic example">
                        <label className="btn btn-secondary">
                            <i className="fas fa-plus"></i> Add <input type="file" hidden onChange={assignHandler}/>
                        </label>
                        <Button variant="secondary">
                            <svg xmlns="http://www.w3.org/2000/svg" width={16} height={16} fill="currentColor" className="bi bi-arrow-clockwise" viewBox="0 0 16 16">
                                <path fillRule="evenodd" d="M8 3a5 5 0 1 0 4.546 2.914.5.5 0 0 1 .908-.417A6 6 0 1 1 8 2v1z" />
                                <path d="M8 4.466V.534a.25.25 0 0 1 .41-.192l2.36 1.966c.12.1.12.284 0 .384L8.41 4.658A.25.25 0 0 1 8 4.466z" />
                            </svg>
                        </Button>
                        <Button variant="secondary">
                            <svg xmlns="http://www.w3.org/2000/svg" width={16} height={16} fill="currentColor" className="bi bi-arrow-counterclockwise" viewBox="0 0 16 16">
                                <path fillRule="evenodd" d="M8 3a5 5 0 1 1-4.546 2.914.5.5 0 0 0-.908-.417A6 6 0 1 0 8 2v1z" />
                                <path d="M8 4.466V.534a.25.25 0 0 0-.41-.192L5.23 2.308a.25.25 0 0 0 0 .384l2.36 1.966A.25.25 0 0 0 8 4.466z" />
                            </svg>
                        </Button>
                    </ButtonGroup>
                    <Button variant="primary" onClick={showCroppedImage} size="sm">Crop</Button>
                </Modal.Footer>
            </Modal>
        </>
    );
}

export default FaUserImg;
