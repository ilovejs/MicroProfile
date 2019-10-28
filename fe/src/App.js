import React, { useState, useEffect, Fragment } from 'react'
import AddProfileForm from './forms/AddProfileForm'
import EditProfileForm from './forms/EditProfileForm'
import ProfileTable from './tables/ProfileTable'
import axios from 'axios'

const App = () => {

  const [ data, setData  ] = useState({ profiles: []});

  useEffect(() => {
    const fetchData = async () => {
      const result = await axios('http://localhost:8080/profiles')
      console.log(result)
      setData(result.data)
    }
    fetchData()
  }, []);

  // Setting states
	// const profileData = [
	// 	{ id: 1, name: 'Tasha', gender: 0, dob: "2019-03-01", postCode: "2001", phoneNo: "0423312313" },
	// 	{ id: 2, name: 'Mike', gender: 1, dob: "2009-03-01", postCode: "3001", phoneNo: "042222222" },
	// 	{ id: 3, name: 'Joe', gender: 0, dob: "2001-02-11", postCode: "2011", phoneNo: "0433333333" },
	// ]

	const initialFormState = { id: 4, name: '', gender: 0, dob: "2001-02-11", postCode: "2011", phoneNo: ""  }
	const [ currentProfile, setCurrentProfile ] = useState(initialFormState)

  const [ editing, setEditing ] = useState(false)

	// CRUD
	const addProfile = profile => {
		profile.id = data.length + 1
		setData([ ...data, profile ])
	}

	const deleteProfile = id => {
		setEditing(false)
		setData(data.filter(p => p.id !== id))
	}

	const updateProfile = (id, updatedProfile) => {
		setEditing(false)
		setData(data.map(p => (p.id === id ? updatedProfile : p)))
	}

	const editRow = user => {
    setEditing(true)
    // { id: 1, name: 'Tania', gender: 0, dob: "2019-03-01", postcode: "2001", phoneNo: "0423312313" },
    setCurrentProfile({
      id: user.id,
      name: user.name,
      gender: user.gender,
      dob: user.dob,
      postCode: user.postCode,
      phoneNo: user.phoneNo
    })
	}

	return (
		<div className="container">
			<h1>CRUD for Profile</h1>
			<div className="flex-row">
				<div className="flex-large">
					{editing ? (
						<Fragment>
							<h2>Edit</h2>
							<EditProfileForm
								editing={editing}
								setEditing={setEditing}
								currentProfile={currentProfile}
								updateProfile={updateProfile}
							/>
						</Fragment>
					) : (
						<Fragment>
							<h2>Add</h2>
							<AddProfileForm addUser={addProfile} />
						</Fragment>
					)}
				</div>

				<div className="flex-large">
					<h2>View users</h2>
          <ProfileTable
            profiles={data}
            editRow={editRow}
            deleteProfile={deleteProfile} />
				</div>
			</div>
		</div>
	)
}

export default App
