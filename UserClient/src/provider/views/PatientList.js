import React from 'react';

const PatientList = () => (
  <div className="mainPanel">
  	<h2>List of patients</h2>
  	<table className="table-striped">
			  <thead>
			    <tr>
			      <th>First name</th>
			      <th>Last name</th>
			      <th>DOB</th>
			      <th>Insurance</th>
			    </tr>
			  </thead>
			  <tbody>
			    <tr>
			      <td>Random</td>
			      <td>Person</td>
			      <td>3/8/1998</td>
			      <td>An insurance</td>
			    </tr>
			    <tr>
			      <td>Another</td>
			      <td>Person</td>
			      <td>3/8/1998</td>
			      <td>Another insurance</td>
			    </tr>
			  </tbody>
    </table>
  </div>

);

export default PatientList;
