import React from 'react'
import { useHistory } from 'react-router-dom/cjs/react-router-dom.min'

export default function Create() {
  const [title, setTitle] = React.useState('')
  const [body, setBody] = React.useState('')
  const [author, setAuthor] = React.useState('mario')
  const [isPending, setIsPending] = React.useState(false)
  const history = useHistory()

  const handleSubmit = (e) => {
    e.preventDefault()
    const blog = { title:title, content:body, create_by:author }
    console.log("thePay_load", JSON.stringify(blog))
    setIsPending(true)
    fetch('http://127.0.0.1:8080/create', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(blog),
    }).then(() => {
      console.log('new blog created')
      setIsPending(false)
      history.push('/') // go back to the homepage
    })
  }
  return (
    <div className='create'>
      <h2> Add a New Blog</h2>
      <form onSubmit={handleSubmit}>
        <label>Blog title:</label>
        <input
          type='text'
          required
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <label>Blog Body:</label>
        <textarea
          required
          value={body}
          onChange={(e) => setBody(e.target.value)}
        >
        </textarea>
        <label> Blog Author:</label>
        <select
          value={author}
          onChange={(e) => setAuthor(e.target.value)}
        >
          <option value='mario'>mario</option>
          <option value='yoshi'>yoshi</option>
        </select>
        {!isPending &&<button  >Add Blog</button> }
        {isPending &&<button  disabled  >Adding Blog....</button> }
      </form>
    </div>
  )
}
