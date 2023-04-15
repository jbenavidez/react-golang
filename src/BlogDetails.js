import { useParams } from "react-router-dom/cjs/react-router-dom.min"
import useFetch from './useFetch'
import { useHistory } from 'react-router-dom/cjs/react-router-dom.min'

export default function BlogDetails() {
  const history = useHistory()
  const { id } = useParams()
  const { data: blog, error, isPending } = useFetch("http://127.0.0.1:8080/blogs/" + id)
  
  const handleDelete = () => {
    fetch("http://localhost:8000/blogs/" + id, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    }).then(() => {
      history.push("/") // redirect to home 
    })
  }
  return (
    <div className="blog-details">
      {isPending && <div>Loading...</div>}
      {error && <div>Error: {error}</div>}
      {blog && (
        <article>
          <h2>{blog.title}</h2>
          <p> Written by {blog.create_by}</p>
          <div>  {blog.content}</div>
          <button onClick={handleDelete}> Delete</button>
        </article>
      )}

    </div>
  )
}
