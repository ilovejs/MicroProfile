import React, { useState, useEffect } from 'react'

const EditUserForm = props => {
  const [ profile, setProfile ] = useState(props.currentProfile)

  useEffect(
    () => {
      setProfile(props.currentProfile)
    },
    [ props ]
  )
  // You can tell React to skip applying an effect if certain values haven’t changed between re-renders. [ props ]

  const handleInputChange = event => {
    const { name, value } = event.target
    setProfile({ ...profile, [name]: value })
  }

  return (
    <form
      onSubmit={event => {
        event.preventDefault()
        props.updateProfile(profile.id, profile)
      }}
    >
      <label>Name</label>
      <input type="text" name="name" value={profile.name} onChange={handleInputChange} />

      <label>Gender</label>
      <input type="text" name="gender" value={profile.gender} onChange={handleInputChange} />

      <label>DoB</label>
			<input type="text" name="dob" value={profile.dob} onChange={handleInputChange} />

      <label>Post Code</label>
			<input type="text" name="postCode" value={profile.postCode} onChange={handleInputChange} />

      <label>Phone Number</label>
			<input type="text" name="phoneNo" value={profile.phoneNo} onChange={handleInputChange} />

      <button>Update profile</button>
      <button onClick={() => props.setEditing(false)} className="button muted-button">
        Cancel
      </button>
    </form>
  )
}

export default EditUserForm
