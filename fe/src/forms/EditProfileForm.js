import React, { useState, useEffect } from 'react'
import axios from 'axios'
import { PUT_PROFILE_URL_PRE } from '../urls.js'

const EditUserForm = props => {
  const [ profile, setProfile ] = useState(props.currentProfile)

  useEffect(
    () => { setProfile(props.currentProfile) },
    [ props ]
  )
  // You can tell React to skip applying an effect if certain values
  // haven’t changed between re-renders. [ props ]
  const handleInputChange = event => {
    const { name, value } = event.target
    // update only specified kv
    setProfile({ ...profile, [name]: value })
  }

  return (
    <form
      onSubmit={event => {
        event.preventDefault()
        props.updateProfile(profile.id, profile)
        // remote
        console.dir(PUT_PROFILE_URL_PRE)
        axios.post(PUT_PROFILE_URL_PRE + profile.id, {
          "name": profile.name,
          "gender": (profile.gender === "true"),
          "dob": profile.dob,
          "phoneNo": profile.phoneNo,
          "postCode": parseInt(profile.postCode)
        })
        .then(res => {
          console.log(res.data)
        }).catch(e => {
          alert(e)
        })
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
