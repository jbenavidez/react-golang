import React from 'react'
import { Link } from 'react-router-dom/cjs/react-router-dom.min'

export default function BlogList({ blogs, title, handleDelete }) {

  return (
    <div children="blog-list">
      <h2>{title}</h2>
      {blogs.map((blog) => (
        <div className='blog-preview' key={blog.id}>
          <Link to={`/blogs/${blog.id}`}>
            <h2> {blog.title}</h2>
            <p> Witter by {blog.author}</p>
          </Link>
        </div>
      ))}

    </div>
  )
}
