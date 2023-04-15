import React from 'react'
import BlogList from './BlogList'
import useFetch from './useFetch'
export default function Home() {
  const { data, isPending, error } = useFetch('http://127.0.0.1:8080/blogs')
  
  return (
    <div>
      <div className='home'>
        {error && <div>{error}</div>}
        {isPending && <div>Loading...</div>}
        {
          data && <BlogList blogs={data} title="All Blogs" />

        }
      </div>
    </div>
  )
}
