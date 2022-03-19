import './App.css';
import CoreUI from './CoreUI/CoreUI';
import { Routes, Route } from "react-router-dom";
import Staff from './Staff/Staff';
import StaffAdd from './Staff/StaffAdd';
import Dashboard from './Dashboard/Dashboard';
import { ToastContainer } from 'react-toastify';

function App() {
  return (
    <CoreUI>
      <ToastContainer />
      <Routes>
        <Route path="/u" exact element={<Dashboard />} />
        <Route path="/u/staff" exact element={<Staff />} />
        <Route path="/u/staff/add" exact element={<StaffAdd />} />
        <Route path="/u/staff/edit/:id" exact element={<StaffAdd />} />
      </Routes>
    </CoreUI>
  );
}

export default App;
