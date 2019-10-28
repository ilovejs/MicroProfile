import React, { useState } from 'react'
import axios from 'axios'
import POST_PROFILE_URL from '../urls'

const AddProfileForm = props => {
  const initialFormState = {
    id: 0,
    name: "",
    gender: true,
    dob: new Date().toISOString().substring(0, 10),
    postCode: 2000,
    phoneNo: '',
  }

	const [ user, setProfile ] = useState(initialFormState)

	const handleInputChange = event => {
		const { name, value } = event.target
		setProfile({ ...user, [name]: value })
  }

  // handleChange(event) {
  //   setState({value: event.target.value});
  // }

	return (
		<form
			onSubmit={event => {
				event.preventDefault()
        if (!user.name || !user.dob) return
        // Add user button
				props.addProfile(user)
        setProfile(initialFormState)
        // remote
        // const {id, ...userNoId} = user
        axios.post(POST_PROFILE_URL,{
          "name": user.name,
          "gender": user.gender,
          "dob": user.dob,
          "phoneNo": user.phoneNo,
          "postCode": parseInt(user.postCode)
        })
        .then(res => {
          // console.log(`userNoId: `, userNoId);
          console.log(res.data);
        })
			}}
		>
			<label>Name</label>
			<input type="text" name="name" value={user.name} onChange={handleInputChange} />

      <label>Gender</label>
      <select name="gender" value={user.gender} onChange={handleInputChange}>
        <option value="true">Male</option>
        <option value="false">Female</option>
      </select>

      <label>DoB</label>
			<input type="text" name="dob" value={user.dob} onChange={handleInputChange} />

      <label>Post Code</label>
			<input type="number" name="postCode" value={user.postCode} onChange={handleInputChange} />

      <label>Phone Number</label>
			<input type="text" name="phoneNo" value={user.phoneNo} onChange={handleInputChange} />


      <button>Add Profile</button>
		</form>
	)
}

export default AddProfileForm
