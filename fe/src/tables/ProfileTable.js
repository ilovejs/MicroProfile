import React from 'react'

const ProfileTable = props => (
  <table>
    <thead>
      <tr>
        <th>Name</th>
        <th>Gender</th>
        <th>DoB</th>
        <th>Post Code</th>
        <th>Phone No</th>
        <th>Action</th>
      </tr>
    </thead>
    <tbody>
      {props.profiles.length > 0 ? (
        props.profiles.map(profile => (
          <tr key={profile.id}>
            <td>{profile.name}</td>
            <td>{profile.gender}</td>
            <td>{profile.dob}</td>
            <td>{profile.postCode}</td>
            <td>{profile.phoneNo}</td>
            <td>
              <button
                onClick={() => {
                  props.editRow(profile)
                }}
                className="button muted-button"
              >
                Edit
              </button>
              <button
                onClick={() => props.deleteProfile(profile.id)}
                className="button muted-button"
              >
                Delete
              </button>
            </td>
          </tr>
        ))
      ) : (
        <tr>
          <td colSpan={3}>No profiles</td>
        </tr>
      )}
    </tbody>
  </table>
)

export default ProfileTable
