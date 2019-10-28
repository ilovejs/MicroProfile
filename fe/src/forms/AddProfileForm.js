import React, { useState } from 'react'

const AddProfileForm = props => {
  const initialFormState = {
    id: 4,
    name: "",
    gender: "",
    dob: new Date().toISOString().substring(0, 10),
    phoneNo: '',
    postCode: 2000
  }

	const [ user, setProfile ] = useState(initialFormState)

	const handleInputChange = event => {
		const { name, value } = event.target
		setProfile({ ...user, [name]: value })
	}

	return (
		<form
			onSubmit={event => {
				event.preventDefault()
				if (!user.name || !user.dob) return

				props.addUser(user)
				setProfile(initialFormState)
			}}
		>
			<label>Name</label>
			<input type="text" name="name" value={user.name} onChange={handleInputChange} />

      <label>Gender</label>
			<input type="number" name="gender" value={user.gender} onChange={handleInputChange} />

      <label>DoB</label>
			<input type="text" name="dob" value={user.dob} onChange={handleInputChange} />

      <label>Post Code</label>
			<input type="text" name="postCode" value={user.postCode} onChange={handleInputChange} />

      <label>Phone Number</label>
			<input type="text" name="phoneNo" value={user.phoneNo} onChange={handleInputChange} />


      <button>Add new user</button>
		</form>
	)
}

export default AddProfileForm
